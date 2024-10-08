package fleet

import (
	"encoding/json"
	"errors"
	fleetService "fms/internal/interactor/service/fleet"

	"fms/internal/interactor/pkg/util"

	fleetModel "fms/internal/interactor/models/fleets"

	"gorm.io/gorm"

	"fms/internal/interactor/pkg/util/code"
	"fms/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *fleetModel.Create) (int, any)
	GetByList(input *fleetModel.Fields) (int, any)
	GetByListNoPagination(input *fleetModel.Field) (int, any)
	GetBySingle(input *fleetModel.Field) (int, any)
	Delete(input *fleetModel.Field) (int, any)
	Update(input *fleetModel.Update) (int, any)
}

type manager struct {
	FleetService fleetService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		FleetService: fleetService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *fleetModel.Create) (int, any) {
	defer trx.Rollback()

	// 判斷車隊ID是否重複
	quantity, _ := m.FleetService.GetByQuantity(&fleetModel.Field{
		FleetCode: util.PointerString(input.FleetCode),
	})

	if quantity > 0 {
		log.Info("FleetCode already exists. FleetCode: ", input.FleetCode)
		return code.BadRequest, code.GetCodeMessage(code.BadRequest, "Fleet already exists.")
	}

	fleetBase, err := m.FleetService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, fleetBase.ID)
}

func (m *manager) GetByList(input *fleetModel.Fields) (int, any) {
	output := &fleetModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, fleetBase, err := m.FleetService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	fleetByte, err := json.Marshal(fleetBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(fleetByte, &output.Fleets)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, fleet := range output.Fleets {
		fleet.CreatedBy = *fleetBase[i].CreatedByUsers.Name
		fleet.UpdatedBy = *fleetBase[i].UpdatedByUsers.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetByListNoPagination(input *fleetModel.Field) (int, any) {
	output := &fleetModel.ListNoPagination{}
	fleetBase, err := m.FleetService.GetByListNoPagination(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	fleetByte, err := json.Marshal(fleetBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(fleetByte, &output.Fleets)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *fleetModel.Field) (int, any) {
	fleetBase, err := m.FleetService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &fleetModel.Single{}
	fleetByte, _ := json.Marshal(fleetBase)
	err = json.Unmarshal(fleetByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.CreatedBy = *fleetBase.CreatedByUsers.Name
	output.UpdatedBy = *fleetBase.UpdatedByUsers.Name

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *fleetModel.Field) (int, any) {
	_, err := m.FleetService.GetBySingle(&fleetModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.FleetService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *fleetModel.Update) (int, any) {
	fleetBase, err := m.FleetService.GetBySingle(&fleetModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.FleetService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, fleetBase.ID)
}
