package transport_task

import (
	"encoding/json"
	"errors"
	transportOrderService "fms/internal/interactor/service/transport_order"
	transportTaskService "fms/internal/interactor/service/transport_task"

	transportOrderModel "fms/internal/interactor/models/transport_orders"
	"fms/internal/interactor/pkg/util"

	transportTaskModel "fms/internal/interactor/models/transport_tasks"

	"gorm.io/gorm"

	"fms/internal/interactor/pkg/util/code"
	"fms/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *transportTaskModel.Create) (int, any)
	GetByList(input *transportTaskModel.Fields) (int, any)
	GetBySingle(input *transportTaskModel.Field) (int, any)
	Delete(trx *gorm.DB, input *transportTaskModel.Field) (int, any)
	Update(trx *gorm.DB, input *transportTaskModel.Update) (int, any)
}

type manager struct {
	TransportTaskService  transportTaskService.Service
	TransportOrderService transportOrderService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		TransportTaskService:  transportTaskService.Init(db),
		TransportOrderService: transportOrderService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *transportTaskModel.Create) (int, any) {
	defer trx.Rollback()

	transportTaskBase, err := m.TransportTaskService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步修改transport_orders的順序
	if len(input.TransportOrderIDs) > 0 {
		for i, transportOrderID := range input.TransportOrderIDs {
			err = m.TransportOrderService.WithTrx(trx).Update(&transportOrderModel.Update{
				ID:              transportOrderID,
				TransportTaskID: transportTaskBase.ID,
				Sequence:        util.PointerInt64(int64(i + 1)),
			})
			if err != nil {
				log.Error(err)
				return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
			}
		}
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, transportTaskBase.ID)
}

func (m *manager) GetByList(input *transportTaskModel.Fields) (int, any) {
	output := &transportTaskModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, transportTaskBase, err := m.TransportTaskService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	transportTaskByte, err := json.Marshal(transportTaskBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(transportTaskByte, &output.TransportTasks)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, transportTask := range output.TransportTasks {
		transportOrderByte, err := json.Marshal(transportTaskBase[i].TransportOrders)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		err = json.Unmarshal(transportOrderByte, &transportTask.TransportOrders)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		transportTask.CreatedBy = *transportTaskBase[i].CreatedByUsers.Name
		transportTask.UpdatedBy = *transportTaskBase[i].UpdatedByUsers.Name
		transportTask.VehicleName = *transportTaskBase[i].Vehicles.Name
		transportTask.DriverName = *transportTaskBase[i].Drivers.Name
		for j, transportOrder := range transportTask.TransportOrders {
			for k, transportOrderDetail := range transportOrder.TransporterOrderDetails {
				transportOrderDetail.TrailerCode = *transportTaskBase[i].TransportOrders[j].TransportOrderDetails[k].Trailers.Code
			}
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *transportTaskModel.Field) (int, any) {
	transportTaskBase, err := m.TransportTaskService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &transportTaskModel.Single{}
	transportTaskByte, _ := json.Marshal(transportTaskBase)
	err = json.Unmarshal(transportTaskByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	transportOrderByte, err := json.Marshal(transportTaskBase.TransportOrders)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(transportOrderByte, &output.TransportOrders)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.CreatedBy = *transportTaskBase.CreatedByUsers.Name
	output.UpdatedBy = *transportTaskBase.UpdatedByUsers.Name
	output.VehicleName = *transportTaskBase.Vehicles.Name
	output.DriverName = *transportTaskBase.Drivers.Name
	for i, transportOrder := range output.TransportOrders {
		for j, transportOrderDetail := range transportOrder.TransporterOrderDetails {
			transportOrderDetail.TrailerCode = *transportTaskBase.TransportOrders[i].TransportOrderDetails[j].Trailers.Code
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(trx *gorm.DB, input *transportTaskModel.Field) (int, any) {
	defer trx.Rollback()

	transportTaskBase, err := m.TransportTaskService.GetBySingle(&transportTaskModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.TransportTaskService.WithTrx(trx).Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步刪除transport_orders的順序
	if len(transportTaskBase.TransportOrders) > 0 {
		var ids []*string
		for _, transportOrder := range transportTaskBase.TransportOrders {
			ids = append(ids, transportOrder.ID)
		}
		err = m.TransportOrderService.WithTrx(trx).Update(&transportOrderModel.Update{
			IDs:             ids,
			TransportTaskID: nil,
			Sequence:        nil,
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(trx *gorm.DB, input *transportTaskModel.Update) (int, any) {
	defer trx.Rollback()

	transportTaskBase, err := m.TransportTaskService.GetBySingle(&transportTaskModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.TransportTaskService.WithTrx(trx).Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 將原本的順序清空
	if len(transportTaskBase.TransportOrders) > 0 {
		var ids []*string
		for _, transportOrder := range transportTaskBase.TransportOrders {
			ids = append(ids, transportOrder.ID)
		}
		err = m.TransportOrderService.WithTrx(trx).Update(&transportOrderModel.Update{
			IDs:             ids,
			TransportTaskID: nil,
			Sequence:        nil,
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	// 同步修改transport_orders的順序
	if len(input.TransportOrderIDs) > 0 {
		// 重新定義順序
		for i, transportOrderID := range input.TransportOrderIDs {
			err = m.TransportOrderService.WithTrx(trx).Update(&transportOrderModel.Update{
				ID:              transportOrderID,
				TransportTaskID: transportTaskBase.ID,
				Sequence:        util.PointerInt64(int64(i + 1)),
			})
			if err != nil {
				log.Error(err)
				return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
			}
		}
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, transportTaskBase.ID)
}
