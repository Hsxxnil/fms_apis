package gps_device

import (
	"encoding/json"
	"errors"
	gpsDeviceService "fms/internal/interactor/service/gps_device"

	"fms/internal/interactor/pkg/util"

	gpsDeviceModel "fms/internal/interactor/models/gps_devices"

	"gorm.io/gorm"

	"fms/internal/interactor/pkg/util/code"
	"fms/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *gpsDeviceModel.Create) (int, any)
	GetByList(input *gpsDeviceModel.Fields) (int, any)
	GetBySingle(input *gpsDeviceModel.Field) (int, any)
	Delete(input *gpsDeviceModel.Field) (int, any)
	Update(input *gpsDeviceModel.Update) (int, any)
}

type manager struct {
	gpsDeviceService gpsDeviceService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		gpsDeviceService: gpsDeviceService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *gpsDeviceModel.Create) (int, any) {
	defer trx.Rollback()

	gpsDeviceBase, err := m.gpsDeviceService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, gpsDeviceBase.ID)
}

func (m *manager) GetByList(input *gpsDeviceModel.Fields) (int, any) {
	output := &gpsDeviceModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, gpsDeviceBase, err := m.gpsDeviceService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	gpsDeviceByte, err := json.Marshal(gpsDeviceBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(gpsDeviceByte, &output.GpsDevices)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, gpsDevice := range output.GpsDevices {
		gpsDevice.CreatedBy = *gpsDeviceBase[i].CreatedByUsers.Name
		gpsDevice.UpdatedBy = *gpsDeviceBase[i].UpdatedByUsers.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *gpsDeviceModel.Field) (int, any) {
	gpsDeviceBase, err := m.gpsDeviceService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &gpsDeviceModel.Single{}
	gpsDeviceByte, _ := json.Marshal(gpsDeviceBase)
	err = json.Unmarshal(gpsDeviceByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.CreatedBy = *gpsDeviceBase.CreatedByUsers.Name
	output.UpdatedBy = *gpsDeviceBase.UpdatedByUsers.Name

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *gpsDeviceModel.Field) (int, any) {
	_, err := m.gpsDeviceService.GetBySingle(&gpsDeviceModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.gpsDeviceService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *gpsDeviceModel.Update) (int, any) {
	gpsDeviceBase, err := m.gpsDeviceService.GetBySingle(&gpsDeviceModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.gpsDeviceService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, gpsDeviceBase.ID)
}
