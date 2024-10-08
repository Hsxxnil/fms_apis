package status

import (
	"encoding/json"
	"errors"
	statusService "fms/internal/interactor/service/status"

	"fms/internal/interactor/pkg/util"

	statusModel "fms/internal/interactor/models/statuses"

	"gorm.io/gorm"

	"fms/internal/interactor/pkg/util/code"
	"fms/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *statusModel.Create) (int, any)
	GetByList(input *statusModel.Fields) (int, any)
	GetBySingle(input *statusModel.Field) (int, any)
	Delete(input *statusModel.Field) (int, any)
	Update(input *statusModel.Update) (int, any)
}

type manager struct {
	statusService statusService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		statusService: statusService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *statusModel.Create) (int, any) {
	defer trx.Rollback()

	_, err := m.statusService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, "Create successful!")
}

func (m *manager) GetByList(input *statusModel.Fields) (int, any) {
	output := &statusModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, statusBase, err := m.statusService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	statusByte, err := json.Marshal(statusBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(statusByte, &output.Statuses)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, status := range output.Statuses {
		status.CreatedBy = *statusBase[i].CreatedByUsers.Name
		status.UpdatedBy = *statusBase[i].UpdatedByUsers.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *statusModel.Field) (int, any) {
	statusBase, err := m.statusService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &statusModel.Single{}
	statusByte, _ := json.Marshal(statusBase)
	err = json.Unmarshal(statusByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.CreatedBy = *statusBase.CreatedByUsers.Name
	output.UpdatedBy = *statusBase.UpdatedByUsers.Name

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *statusModel.Field) (int, any) {
	_, err := m.statusService.GetBySingle(&statusModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.statusService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *statusModel.Update) (int, any) {
	_, err := m.statusService.GetBySingle(&statusModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.statusService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Update successful!")
}
