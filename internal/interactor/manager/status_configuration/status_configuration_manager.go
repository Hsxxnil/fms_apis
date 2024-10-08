package status_configuration

import (
	"encoding/json"
	"errors"
	statusConfigurationService "fms/internal/interactor/service/status_configuration"

	"fms/internal/interactor/pkg/util"

	statusConfigurationModel "fms/internal/interactor/models/status_configurations"

	"gorm.io/gorm"

	"fms/internal/interactor/pkg/util/code"
	"fms/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *statusConfigurationModel.Create) (int, any)
	GetByList(input *statusConfigurationModel.Fields) (int, any)
	GetBySingle(input *statusConfigurationModel.Field) (int, any)
	Delete(input *statusConfigurationModel.Field) (int, any)
	Update(input *statusConfigurationModel.Update) (int, any)
}

type manager struct {
	StatusConfigurationService statusConfigurationService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		StatusConfigurationService: statusConfigurationService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *statusConfigurationModel.Create) (int, any) {
	defer trx.Rollback()

	statusConfigurationBase, err := m.StatusConfigurationService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, statusConfigurationBase.ID)
}

func (m *manager) GetByList(input *statusConfigurationModel.Fields) (int, any) {
	output := &statusConfigurationModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, statusConfigurationBase, err := m.StatusConfigurationService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	statusConfigurationByte, err := json.Marshal(statusConfigurationBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(statusConfigurationByte, &output.StatusConfigurations)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, statusConfiguration := range output.StatusConfigurations {
		statusConfiguration.CreatedBy = *statusConfigurationBase[i].CreatedByUsers.Name
		statusConfiguration.UpdatedBy = *statusConfigurationBase[i].UpdatedByUsers.Name
		statusConfiguration.Status = *statusConfigurationBase[i].Statuses.Status
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *statusConfigurationModel.Field) (int, any) {
	statusConfigurationBase, err := m.StatusConfigurationService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &statusConfigurationModel.Single{}
	statusConfigurationByte, _ := json.Marshal(statusConfigurationBase)
	err = json.Unmarshal(statusConfigurationByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.CreatedBy = *statusConfigurationBase.CreatedByUsers.Name
	output.UpdatedBy = *statusConfigurationBase.UpdatedByUsers.Name
	output.Status = *statusConfigurationBase.Statuses.Status

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *statusConfigurationModel.Field) (int, any) {
	_, err := m.StatusConfigurationService.GetBySingle(&statusConfigurationModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.StatusConfigurationService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *statusConfigurationModel.Update) (int, any) {
	statusConfigurationBase, err := m.StatusConfigurationService.GetBySingle(&statusConfigurationModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.StatusConfigurationService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, statusConfigurationBase.ID)
}
