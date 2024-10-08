package gps

import (
	"encoding/json"
	"errors"
	"fms/internal/entity/postgresql/db/vehicles"
	gpsModel "fms/internal/interactor/models/gps"
	gpsDeviceModel "fms/internal/interactor/models/gps_devices"
	vehicleModel "fms/internal/interactor/models/vehicles"
	googleMap "fms/internal/interactor/pkg/google_map"
	"fms/internal/interactor/pkg/util"
	avemaService "fms/internal/interactor/service/avema_data"
	gpsDeviceService "fms/internal/interactor/service/gps_device"
	jasslinService "fms/internal/interactor/service/jasslin_data"
	poisonService "fms/internal/interactor/service/poison_data"
	vehicleService "fms/internal/interactor/service/vehicle"
	"sort"
	"sync"
	"time"

	"fms/internal/interactor/pkg/util/code"
	"fms/internal/interactor/pkg/util/log"
	"gorm.io/gorm"
)

type Manager interface {
	AppGetByListNoPagination(input *gpsModel.Field) (int, any)
	AppGetByLicensePlateList(input *gpsModel.Field) (int, any)
	WebGetByListNoPagination(input *gpsModel.Field) (int, any)
	WebGetByLicensePlateList(input *gpsModel.Field) (int, any)
}

type manager struct {
	AvemaService     avemaService.Service
	PoisonService    poisonService.Service
	JasslinService   jasslinService.Service
	VehicleService   vehicleService.Service
	gpsDeviceService gpsDeviceService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		AvemaService:     avemaService.Init(db),
		PoisonService:    poisonService.Init(db),
		JasslinService:   jasslinService.Init(db),
		VehicleService:   vehicleService.Init(db),
		gpsDeviceService: gpsDeviceService.Init(db),
	}
}

// google map api每秒最大請求次數
const maxRequestsPerSecond = 50

// getStatusTime is a helper function used to get status accumulation times.
func getStatusTime(input map[string]time.Duration) *gpsModel.StatusTime {
	return &gpsModel.StatusTime{
		IdleTime:      input["怠停"].String(),
		LongIdleTime:  input["久停"].String(),
		DrivingTime:   input["行駛"].String(),
		EngineOffTime: input["熄火"].String(),
		LostConnTime:  input["失聯"].String(),
	}
}

// determineStatus is a helper function used to determine status and calculate status accumulation times.
func determineStatus(gpsData []*gpsModel.DetermineStatus, engineOffLimit, idleLimit float64) ([]*gpsModel.RealTimeSingle, map[string]time.Duration, error) {
	log.Debug("determineStatus start")
	// 根據DateTime排序
	sort.Slice(gpsData, func(i, j int) bool {
		return gpsData[i].DateTime.Before(gpsData[j].DateTime)
	})

	// 各狀態累積時間
	var idleTime, longIdleTime, drivingTime, engineOffTime, lostConnTime time.Duration
	// 建立狀態时间映射表
	statusTimeMap := map[string]time.Duration{
		"怠停": idleTime,
		"久停": longIdleTime,
		"行駛": drivingTime,
		"熄火": engineOffTime,
		"失聯": lostConnTime,
	}

	for i, current := range gpsData {
		if i > 0 {
			prev := gpsData[i-1]
			currentDiff := current.DateTime.Sub(prev.DateTime)
			if prev.StatusTime == "" {
				prev.StatusTime = "0s"
			}
			prevStatusTime, _ := time.ParseDuration(prev.StatusTime)

			if prev.Address == current.Address {
				currentStatusTime := currentDiff + prevStatusTime
				if currentStatusTime.Minutes() <= idleLimit {
					current.Status = "怠停"
				} else if currentStatusTime.Minutes() > idleLimit && currentStatusTime.Minutes() < engineOffLimit {
					current.Status = "久停"
				} else if currentStatusTime.Minutes() >= engineOffLimit {
					current.Status = "失聯"
				}
			}

			// 預設狀態為"行駛"
			if current.Status == "" {
				current.Status = "行駛"
			}

			// 判斷熄火源
			if current.Io1 == 6 || (current.ReportID == 51 && current.Inputs == 0) {
				current.Status = "熄火"
			}

			// 判斷狀態累積時間
			var StatusTime time.Duration
			if current.Status == prev.Status {
				StatusTime = currentDiff + prevStatusTime
			} else {
				StatusTime = currentDiff
			}
			current.StatusTime = StatusTime.String()

			// 根據各狀態累積相應時間
			statusTimeMap[current.Status] += currentDiff
		}
	}

	var output []*gpsModel.RealTimeSingle
	gpsByte, err := json.Marshal(gpsData)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	err = json.Unmarshal(gpsByte, &output)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	return output, statusTimeMap, nil
}

func (m *manager) AppGetByListNoPagination(input *gpsModel.Field) (int, any) {
	var (
		output         = &gpsModel.RealTimeList{}
		allData        []*gpsModel.RealTimeSingle
		wg             sync.WaitGroup
		engineOffLimit = 30.0
		idleLimit      = 10.0
	)

	// 取得車輛sid
	vehicleBase, err := m.VehicleService.GetByListNoPagination(&vehicleModel.Field{
		FleetID: util.PointerString(input.FleetID),
	})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	// 取得車機資料
	gpsDeviceBase, err := m.gpsDeviceService.GetByListNoPagination(&gpsDeviceModel.Field{})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	// 將車輛sid與車機model映射表
	vehicleToModelMap := make(map[string]string)
	for _, vehicle := range vehicleBase {
		vehicleToModelMap[*vehicle.SID] = ""
	}

	for _, gpsDevice := range gpsDeviceBase {
		if _, exists := vehicleToModelMap[*gpsDevice.SID]; exists {
			if gpsDevice.Model != nil {
				vehicleToModelMap[*gpsDevice.SID] = *gpsDevice.Model
			}
		}
	}

	// 根據車機model加入相應的sid陣列
	for vehicleSID, model := range vehicleToModelMap {
		if model == "Z1" || model == "AT35" {
			input.AvemaSID = append(input.AvemaSID, util.PointerString(vehicleSID))
		} else if model == "CH68" || model == "D1 PLUS" {
			input.PoisonSID = append(input.PoisonSID, util.PointerString(vehicleSID))
		} else if model == "208" || model == "306" {
			input.JasslinSID = append(input.JasslinSID, util.PointerString(vehicleSID))
		}
	}

	goroutineErr := make(chan error)
	// 取得avema_data
	for _, avema := range input.AvemaSID {
		wg.Add(1)
		go func(avema *string) {
			defer wg.Done()
			var avemaData []*gpsModel.DetermineStatus
			avemaBase, err := m.AvemaService.GetByLastData(&gpsModel.Field{
				SID: avema,
			})
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			avemaByte, err := json.Marshal(avemaBase)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			err = json.Unmarshal(avemaByte, &avemaData)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			// 取得地址
			for i, data := range avemaData {
				if i > 0 {
					if avemaData[i-1].Lon != data.Lon || avemaData[i-1].Lat != data.Lat {
						address, err := googleMap.LatLngToAddress(data.Lat, data.Lon)
						if err != nil {
							log.Error(err)
							goroutineErr <- err
						}
						data.Address = address
						log.Debug(data.Address)
					} else {
						data.Address = avemaData[i-1].Address
					}
				} else {
					// 取得第一筆地址
					address, err := googleMap.LatLngToAddress(data.Lat, data.Lon)
					if err != nil {
						log.Error(err)
						goroutineErr <- err
					}
					data.Address = address
					log.Debug(data.Address)
				}
			}

			// 判斷狀態
			resultData, _, err := determineStatus(avemaData, engineOffLimit, idleLimit)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			allData = append(allData, resultData...)
		}(avema)
	}

	// 取得poison_data
	for _, poison := range input.PoisonSID {
		var poisonData []*gpsModel.DetermineStatus
		wg.Add(1)
		go func(poison *string) {
			defer wg.Done()
			poisonBase, err := m.PoisonService.GetByLastData(&gpsModel.Field{
				SID: poison,
			})
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			poisonByte, err := json.Marshal(poisonBase)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			err = json.Unmarshal(poisonByte, &poisonData)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			// 取得地址
			for i, data := range poisonData {
				if i > 0 {
					if poisonData[i-1].Lon != data.Lon || poisonData[i-1].Lat != data.Lat {
						address, err := googleMap.LatLngToAddress(data.Lat, data.Lon)
						if err != nil {
							log.Error(err)
							goroutineErr <- err
						}
						data.Address = address
						log.Debug(data.Address)
					} else {
						data.Address = poisonData[i-1].Address
					}
				} else {
					// 取得第一筆地址
					address, err := googleMap.LatLngToAddress(data.Lat, data.Lon)
					if err != nil {
						log.Error(err)
						goroutineErr <- err
					}
					data.Address = address
					log.Debug(data.Address)
				}
			}

			// 判斷狀態
			resultData, _, err := determineStatus(poisonData, engineOffLimit, idleLimit)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			allData = append(allData, resultData...)
		}(poison)
	}

	// 取得jasslin_data
	for _, jasslin := range input.JasslinSID {
		var jasslinData []*gpsModel.DetermineStatus
		wg.Add(1)
		go func(jasslin *string) {
			defer wg.Done()
			jasslinBase, err := m.JasslinService.GetByLastData(&gpsModel.Field{
				SID: jasslin,
			})
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			jasslinByte, err := json.Marshal(jasslinBase)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			err = json.Unmarshal(jasslinByte, &jasslinData)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			// 取得地址
			for i, data := range jasslinData {
				if i > 0 {
					if jasslinData[i-1].Lon != data.Lon || jasslinData[i-1].Lat != data.Lat {
						address, err := googleMap.LatLngToAddress(data.Lat, data.Lon)
						if err != nil {
							log.Error(err)
							goroutineErr <- err
						}
						data.Address = address
						log.Debug(data.Address)
					} else {
						data.Address = jasslinData[i-1].Address
					}
				} else {
					// 取得第一筆地址
					address, err := googleMap.LatLngToAddress(data.Lat, data.Lon)
					if err != nil {
						log.Error(err)
						goroutineErr <- err
					}
					data.Address = address
					log.Debug(data.Address)
				}
			}

			// 判斷狀態
			resultData, _, err := determineStatus(jasslinData, engineOffLimit, idleLimit)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			allData = append(allData, resultData...)
		}(jasslin)
	}

	// 等待所有goroutine完成
	wg.Wait()
	close(goroutineErr)

	// 檢查錯誤
	err = <-goroutineErr
	if err != nil {
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 建立sid與車輛資料映射表
	sidToVehicle := make(map[string]*vehicles.Base)
	for _, vehicle := range vehicleBase {
		sidToVehicle[*vehicle.SID] = vehicle
	}
	// 用來追蹤已處理的sid
	processedSID := make(map[string]bool)

	// 篩選allData每個sid的最後一筆
	for i := len(allData) - 1; i >= 0; i-- {
		data := allData[i]
		sid := data.SID
		vehicle := sidToVehicle[sid]
		if !processedSID[sid] {
			if vehicle != nil {
				data.LicensePlate = *vehicle.LicensePlate
				data.VehicleName = *vehicle.Name
				data.Driver = *vehicle.Driver
			}
			output.Gps = append(output.Gps, data)
			processedSID[sid] = true
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) AppGetByLicensePlateList(input *gpsModel.Field) (int, any) {
	var (
		output         = &gpsModel.RealTimeList{}
		allData        []*gpsModel.DetermineStatus
		wg             sync.WaitGroup
		engineOffLimit = 30.0
		idleLimit      = 10.0
	)

	// 取得車輛sid
	vehicleBase, err := m.VehicleService.GetBySingle(&vehicleModel.Field{
		LicensePlate: input.LicensePlate,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	input.SID = vehicleBase.SID

	// 取得avema_data
	avemaBase, err := m.AvemaService.GetBySIDList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 取得poison_data
	poisonBase, err := m.PoisonService.GetBySIDList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 取得jasslin_data
	jasslinBase, err := m.JasslinService.GetBySIDList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	if len(avemaBase) > 0 {
		avemaByte, err := json.Marshal(avemaBase)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		err = json.Unmarshal(avemaByte, &allData)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

	} else if len(poisonBase) > 0 {
		poisonByte, err := json.Marshal(poisonBase)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		err = json.Unmarshal(poisonByte, &allData)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

	} else if len(jasslinBase) > 0 {
		jasslinByte, err := json.Marshal(jasslinBase)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		err = json.Unmarshal(jasslinByte, &allData)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	// todo 取得地址待優化
	goroutineErr := make(chan error)
	done := make(chan bool)
	for i, data := range allData {
		if i > 0 {
			if allData[i-1].Address == "" {
				<-done
			}
			if allData[i-1].Lon != data.Lon || allData[i-1].Lat != data.Lat {
				wg.Add(1)
				go func(data *gpsModel.DetermineStatus) {
					defer wg.Done()
					// 取得地址
					address, err := googleMap.LatLngToAddress(data.Lat, data.Lon)
					if err != nil {
						log.Error(err)
						goroutineErr <- err
					}
					data.Address = address
					log.Debug(data.Address)

					// 最後一筆不存在等待中的goroutine
					if i < len(allData)-1 {
						done <- true
					}
				}(data)

				// 控制請求速率
				time.Sleep(time.Second / maxRequestsPerSecond)
			} else {
				data.Address = allData[i-1].Address
			}
		} else {
			// 取得第一筆地址
			address, err := googleMap.LatLngToAddress(data.Lat, data.Lon)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}
			data.Address = address
			log.Debug(data.Address)
		}
	}

	// 等待所有goroutine完成
	wg.Wait()
	close(goroutineErr)

	// 檢查錯誤
	err = <-goroutineErr
	if err != nil {
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 判斷狀態
	resultData, statusTime, err := determineStatus(allData, engineOffLimit, idleLimit)
	if err != nil {
		log.Error(err)
		goroutineErr <- err
	}
	output.StatusTime = getStatusTime(statusTime)

	// 回傳查詢時間內的資料
	for _, data := range resultData {
		if data.DateTime.After(*input.FilterStartTime) {
			output.Gps = append(output.Gps, data)
			output.Total.Total++
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) WebGetByListNoPagination(input *gpsModel.Field) (int, any) {
	var (
		output         = &gpsModel.RealTimeList{}
		allData        []*gpsModel.RealTimeSingle
		wg             sync.WaitGroup
		engineOffLimit = 30.0
		idleLimit      = 10.0
	)

	// 取得車輛sid
	vehicleBase, err := m.VehicleService.GetByListNoPagination(&vehicleModel.Field{
		FleetID: util.PointerString(input.FleetID),
	})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	// 取得車機資料
	gpsDeviceBase, err := m.gpsDeviceService.GetByListNoPagination(&gpsDeviceModel.Field{})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	// 將車輛sid與車機model映射表
	vehicleToModelMap := make(map[string]string)
	for _, vehicle := range vehicleBase {
		vehicleToModelMap[*vehicle.SID] = ""
	}

	for _, gpsDevice := range gpsDeviceBase {
		if _, exists := vehicleToModelMap[*gpsDevice.SID]; exists {
			if gpsDevice.Model != nil {
				vehicleToModelMap[*gpsDevice.SID] = *gpsDevice.Model
			}
		}
	}

	// 根據車機model加入相應的sid陣列
	for vehicleSID, model := range vehicleToModelMap {
		if model == "Z1" || model == "AT35" {
			input.AvemaSID = append(input.AvemaSID, util.PointerString(vehicleSID))
		} else if model == "CH68" || model == "D1 PLUS" {
			input.PoisonSID = append(input.PoisonSID, util.PointerString(vehicleSID))
		} else if model == "208" || model == "306" {
			input.JasslinSID = append(input.JasslinSID, util.PointerString(vehicleSID))
		}
	}

	goroutineErr := make(chan error)
	// 取得avema_data
	for _, avema := range input.AvemaSID {
		wg.Add(1)
		go func(avema *string) {
			defer wg.Done()
			var avemaData []*gpsModel.DetermineStatus
			avemaBase, err := m.AvemaService.GetByLastData(&gpsModel.Field{
				SID: avema,
			})
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			avemaByte, err := json.Marshal(avemaBase)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			err = json.Unmarshal(avemaByte, &avemaData)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			// 取得地址
			for i, data := range avemaData {
				if i > 0 {
					if avemaData[i-1].Lon != data.Lon || avemaData[i-1].Lat != data.Lat {
						address, err := googleMap.LatLngToAddress(data.Lat, data.Lon)
						if err != nil {
							log.Error(err)
							goroutineErr <- err
						}
						data.Address = address
						log.Debug(data.Address)
					} else {
						data.Address = avemaData[i-1].Address
					}
				} else {
					// 取得第一筆地址
					address, err := googleMap.LatLngToAddress(data.Lat, data.Lon)
					if err != nil {
						log.Error(err)
						goroutineErr <- err
					}
					data.Address = address
					log.Debug(data.Address)
				}
			}

			// 判斷狀態
			resultData, _, err := determineStatus(avemaData, engineOffLimit, idleLimit)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			allData = append(allData, resultData...)
		}(avema)
	}

	// 取得poison_data
	for _, poison := range input.PoisonSID {
		var poisonData []*gpsModel.DetermineStatus
		wg.Add(1)
		go func(poison *string) {
			defer wg.Done()
			poisonBase, err := m.PoisonService.GetByLastData(&gpsModel.Field{
				SID: poison,
			})
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			poisonByte, err := json.Marshal(poisonBase)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			err = json.Unmarshal(poisonByte, &poisonData)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			// 取得地址
			for i, data := range poisonData {
				if i > 0 {
					if poisonData[i-1].Lon != data.Lon || poisonData[i-1].Lat != data.Lat {
						address, err := googleMap.LatLngToAddress(data.Lat, data.Lon)
						if err != nil {
							log.Error(err)
							goroutineErr <- err
						}
						data.Address = address
						log.Debug(data.Address)
					} else {
						data.Address = poisonData[i-1].Address
					}
				} else {
					// 取得第一筆地址
					address, err := googleMap.LatLngToAddress(data.Lat, data.Lon)
					if err != nil {
						log.Error(err)
						goroutineErr <- err
					}
					data.Address = address
					log.Debug(data.Address)
				}
			}

			// 判斷狀態
			resultData, _, err := determineStatus(poisonData, engineOffLimit, idleLimit)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			allData = append(allData, resultData...)
		}(poison)
	}

	// 取得jasslin_data
	for _, jasslin := range input.JasslinSID {
		var jasslinData []*gpsModel.DetermineStatus
		wg.Add(1)
		go func(jasslin *string) {
			defer wg.Done()
			jasslinBase, err := m.JasslinService.GetByLastData(&gpsModel.Field{
				SID: jasslin,
			})
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			jasslinByte, err := json.Marshal(jasslinBase)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			err = json.Unmarshal(jasslinByte, &jasslinData)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			// 取得地址
			for i, data := range jasslinData {
				if i > 0 {
					if jasslinData[i-1].Lon != data.Lon || jasslinData[i-1].Lat != data.Lat {
						address, err := googleMap.LatLngToAddress(data.Lat, data.Lon)
						if err != nil {
							log.Error(err)
							goroutineErr <- err
						}
						data.Address = address
						log.Debug(data.Address)
					} else {
						data.Address = jasslinData[i-1].Address
					}
				} else {
					// 取得第一筆地址
					address, err := googleMap.LatLngToAddress(data.Lat, data.Lon)
					if err != nil {
						log.Error(err)
						goroutineErr <- err
					}
					data.Address = address
					log.Debug(data.Address)
				}
			}

			// 判斷狀態
			resultData, _, err := determineStatus(jasslinData, engineOffLimit, idleLimit)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}

			allData = append(allData, resultData...)
		}(jasslin)
	}

	// 等待所有goroutine完成
	wg.Wait()
	close(goroutineErr)

	// 檢查錯誤
	err = <-goroutineErr
	if err != nil {
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 建立sid與車輛資料映射表
	sidToVehicle := make(map[string]*vehicles.Base)
	for _, vehicle := range vehicleBase {
		sidToVehicle[*vehicle.SID] = vehicle
	}
	// 用來追蹤已處理的sid
	processedSID := make(map[string]bool)

	// 篩選allData每個sid的最後一筆
	for i := len(allData) - 1; i >= 0; i-- {
		data := allData[i]
		sid := data.SID
		vehicle := sidToVehicle[sid]
		if !processedSID[sid] {
			data.Lng = allData[i].Lon
			if vehicle != nil {
				data.LicensePlate = *vehicle.LicensePlate
				data.VehicleName = *vehicle.Name
				data.Driver = *vehicle.Driver
			}
			output.Gps = append(output.Gps, data)
			processedSID[sid] = true
		}
	}

	// output.Gps根據LicensePlate排序
	sort.Slice(output.Gps, func(i, j int) bool {
		return output.Gps[i].LicensePlate < output.Gps[j].LicensePlate
	})

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) WebGetByLicensePlateList(input *gpsModel.Field) (int, any) {
	var (
		output         = &gpsModel.RealTimeList{}
		allData        []*gpsModel.DetermineStatus
		wg             sync.WaitGroup
		engineOffLimit = 30.0
		idleLimit      = 10.0
	)

	// 取得車輛sid
	vehicleBase, err := m.VehicleService.GetBySingle(&vehicleModel.Field{
		LicensePlate: input.LicensePlate,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	input.SID = vehicleBase.SID

	// 取得avema_data
	avemaBase, err := m.AvemaService.GetBySIDList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 取得poison_data
	poisonBase, err := m.PoisonService.GetBySIDList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 取得jasslin_data
	jasslinBase, err := m.JasslinService.GetBySIDList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	if len(avemaBase) > 0 {
		avemaByte, err := json.Marshal(avemaBase)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		err = json.Unmarshal(avemaByte, &allData)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

	} else if len(poisonBase) > 0 {
		poisonByte, err := json.Marshal(poisonBase)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		err = json.Unmarshal(poisonByte, &allData)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

	} else if len(jasslinBase) > 0 {
		jasslinByte, err := json.Marshal(jasslinBase)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		err = json.Unmarshal(jasslinByte, &allData)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	// todo 取得地址待優化
	goroutineErr := make(chan error)
	done := make(chan bool)
	for i, data := range allData {
		if i > 0 {
			if allData[i-1].Address == "" {
				<-done
			}
			if allData[i-1].Lon != data.Lon || allData[i-1].Lat != data.Lat {
				wg.Add(1)
				go func(data *gpsModel.DetermineStatus) {
					defer wg.Done()
					// 取得地址
					address, err := googleMap.LatLngToAddress(data.Lat, data.Lon)
					if err != nil {
						log.Error(err)
						goroutineErr <- err
					}
					data.Address = address
					log.Debug(data.Address)

					// 最後一筆不存在等待中的goroutine
					if i < len(allData)-1 {
						done <- true
					}
				}(data)

				// 控制請求速率
				time.Sleep(time.Second / maxRequestsPerSecond)
			} else {
				data.Address = allData[i-1].Address
			}
		} else {
			// 取得第一筆地址
			address, err := googleMap.LatLngToAddress(data.Lat, data.Lon)
			if err != nil {
				log.Error(err)
				goroutineErr <- err
			}
			data.Address = address
			log.Debug(data.Address)
		}
	}

	// 等待所有goroutine完成
	wg.Wait()
	close(goroutineErr)

	// 檢查錯誤
	err = <-goroutineErr
	if err != nil {
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 判斷狀態
	resultData, statusTime, err := determineStatus(allData, engineOffLimit, idleLimit)
	if err != nil {
		log.Error(err)
		goroutineErr <- err
	}
	output.StatusTime = getStatusTime(statusTime)

	// 回傳查詢時間內的資料
	for _, data := range resultData {
		if data.DateTime.After(*input.FilterStartTime) {
			output.Gps = append(output.Gps, data)
			output.Total.Total++
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}
