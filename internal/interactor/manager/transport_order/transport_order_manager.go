package transport_order

import (
	"encoding/json"
	"errors"
	"fms/internal/interactor/models/transport_order_details"
	transportOrderModel "fms/internal/interactor/models/transport_orders"
	transportTaskModel "fms/internal/interactor/models/transport_tasks"
	"fms/internal/interactor/pkg/util"
	transportOrderService "fms/internal/interactor/service/transport_order"
	transportOrderDetailService "fms/internal/interactor/service/transport_order_detail"
	transportTaskService "fms/internal/interactor/service/transport_task"

	"gorm.io/gorm"

	"fms/internal/interactor/pkg/util/code"
	"fms/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *transportOrderModel.Create) (int, any)
	GetByList(input *transportOrderModel.Fields) (int, any)
	GetBySingle(input *transportOrderModel.Field) (int, any)
	Delete(trx *gorm.DB, input *transportOrderModel.Field) (int, any)
	Update(trx *gorm.DB, input *transportOrderModel.Update) (int, any)
}

type manager struct {
	TransportOrderService       transportOrderService.Service
	TransportOrderDetailService transportOrderDetailService.Service
	TransportTaskService        transportTaskService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		TransportOrderService:       transportOrderService.Init(db),
		TransportOrderDetailService: transportOrderDetailService.Init(db),
		TransportTaskService:        transportTaskService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *transportOrderModel.Create) (int, any) {
	defer trx.Rollback()

	transportOrderBase, err := m.TransportOrderService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步新增transport_order_detail
	var detailList []*transport_order_details.Create
	if len(input.TransporterOrderDetails) > 0 {
		for _, detail := range input.TransporterOrderDetails {
			details := &transport_order_details.Create{
				TransportOrderID: *transportOrderBase.ID,
				ProductName:      detail.ProductName,
				Quantity:         detail.Quantity,
				UnitPrice:        detail.UnitPrice,
				Tonnage:          detail.Tonnage,
				TrailerID:        detail.TrailerID,
				CreatedBy:        input.CreatedBy,
			}
			detailList = append(detailList, details)
		}

		_, err = m.TransportOrderDetailService.WithTrx(trx).Create(detailList)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, transportOrderBase.ID)
}

func (m *manager) GetByList(input *transportOrderModel.Fields) (int, any) {
	output := &transportOrderModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, transportOrderBase, err := m.TransportOrderService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	transportOrderByte, err := json.Marshal(transportOrderBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(transportOrderByte, &output.TransportOrders)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 取得transport_tasks
	tasks := []transportTaskModel.Single{}
	transportTaskBase, err := m.TransportTaskService.GetByListNoPagination(&transportTaskModel.Field{})
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	transportTaskByte, err := json.Marshal(transportTaskBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(transportTaskByte, &tasks)

	// 建立transport_task map
	transportTaskMap := make(map[string]transportTaskModel.Single)
	for _, transportTask := range tasks {
		transportTaskMap[transportTask.ID] = transportTask
	}

	for i, transportOrder := range output.TransportOrders {
		transportOrder.CreatedBy = *transportOrderBase[i].CreatedByUsers.Name
		transportOrder.UpdatedBy = *transportOrderBase[i].UpdatedByUsers.Name
		transportOrder.ClientName = *transportOrderBase[i].Client.Name
		if transportOrderBase[i].TransportTaskID != nil {
			transportOrder.TransportTaskName = transportTaskMap[*transportOrderBase[i].TransportTaskID].Title
		}
		for j, detail := range transportOrder.TransporterOrderDetails {
			detail.TrailerCode = *transportOrderBase[i].TransportOrderDetails[j].Trailers.Code
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *transportOrderModel.Field) (int, any) {
	transportOrderBase, err := m.TransportOrderService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &transportOrderModel.Single{}
	transportOrderByte, _ := json.Marshal(transportOrderBase)
	err = json.Unmarshal(transportOrderByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 取得transport_tasks
	if transportOrderBase.TransportTaskID != nil {
		transportTaskBase, err := m.TransportTaskService.GetBySingle(&transportTaskModel.Field{
			ID: *transportOrderBase.TransportTaskID,
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
			}

			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
		output.TransportTaskName = *transportTaskBase.Title
	}

	output.CreatedBy = *transportOrderBase.CreatedByUsers.Name
	output.UpdatedBy = *transportOrderBase.UpdatedByUsers.Name
	output.ClientName = *transportOrderBase.Client.Name
	for i, detail := range output.TransporterOrderDetails {
		detail.TrailerCode = *transportOrderBase.TransportOrderDetails[i].Trailers.Code
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(trx *gorm.DB, input *transportOrderModel.Field) (int, any) {
	defer trx.Rollback()

	transportOrderBase, err := m.TransportOrderService.GetBySingle(&transportOrderModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.TransportOrderService.WithTrx(trx).Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步刪除transport_order_detail
	if len(transportOrderBase.TransportOrderDetails) > 0 {
		err = m.TransportOrderDetailService.WithTrx(trx).Delete(&transport_order_details.Field{
			TransportOrderID: util.PointerString(input.ID),
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(trx *gorm.DB, input *transportOrderModel.Update) (int, any) {
	defer trx.Rollback()

	transportOrderBase, err := m.TransportOrderService.GetBySingle(&transportOrderModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.TransportOrderService.WithTrx(trx).Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步更新transport_order_detail
	if len(input.TransporterOrderDetails) > 0 {
		// 刪除transport_order_detail
		err = m.TransportOrderDetailService.WithTrx(trx).Delete(&transport_order_details.Field{
			TransportOrderID: util.PointerString(input.ID),
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		// 新增transport_order_detail
		var detailList []*transport_order_details.Create
		for _, detail := range input.TransporterOrderDetails {
			details := &transport_order_details.Create{
				TransportOrderID: input.ID,
				ProductName:      *detail.ProductName,
				Quantity:         detail.Quantity,
				UnitPrice:        detail.UnitPrice,
				Tonnage:          detail.Tonnage,
				TrailerID:        detail.TrailerID,
				CreatedBy:        *input.UpdatedBy,
			}
			detailList = append(detailList, details)
		}

		_, err = m.TransportOrderDetailService.WithTrx(trx).Create(detailList)
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, transportOrderBase.ID)
}
