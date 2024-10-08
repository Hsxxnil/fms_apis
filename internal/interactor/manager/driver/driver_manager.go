package driver

import (
	"encoding/json"
	"errors"
	driverService "fms/internal/interactor/service/driver"

	"fms/internal/interactor/pkg/util"

	driverModel "fms/internal/interactor/models/drivers"

	"gorm.io/gorm"

	"fms/internal/interactor/pkg/util/code"
	"fms/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *driverModel.Create) (int, any)
	GetByList(input *driverModel.Fields) (int, any)
	GetBySingle(input *driverModel.Field) (int, any)
	Delete(input *driverModel.Field) (int, any)
	Update(input *driverModel.Update) (int, any)
}

type manager struct {
	driverService driverService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		driverService: driverService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *driverModel.Create) (int, any) {
	defer trx.Rollback()

	driverBase, err := m.driverService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, driverBase.ID)
}

func (m *manager) GetByList(input *driverModel.Fields) (int, any) {
	output := &driverModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, driverBase, err := m.driverService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	driverByte, err := json.Marshal(driverBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(driverByte, &output.Drivers)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, driver := range output.Drivers {
		driver.CreatedBy = *driverBase[i].CreatedByUsers.Name
		driver.UpdatedBy = *driverBase[i].UpdatedByUsers.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *driverModel.Field) (int, any) {
	driverBase, err := m.driverService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &driverModel.Single{}
	driverByte, _ := json.Marshal(driverBase)
	err = json.Unmarshal(driverByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.CreatedBy = *driverBase.CreatedByUsers.Name
	output.UpdatedBy = *driverBase.UpdatedByUsers.Name

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *driverModel.Field) (int, any) {
	_, err := m.driverService.GetBySingle(&driverModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.driverService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *driverModel.Update) (int, any) {
	driverBase, err := m.driverService.GetBySingle(&driverModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.driverService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, driverBase.ID)
}
