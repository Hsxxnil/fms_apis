package vehicle

import (
	"encoding/json"
	"errors"
	vehicleService "fms/internal/interactor/service/vehicle"

	"fms/internal/interactor/pkg/util"

	vehicleModel "fms/internal/interactor/models/vehicles"

	"gorm.io/gorm"

	"fms/internal/interactor/pkg/util/code"
	"fms/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *vehicleModel.Create) (int, any)
	GetByList(input *vehicleModel.Fields) (int, any)
	GetByListNoPagination(input *vehicleModel.Field) (int, any)
	GetBySingle(input *vehicleModel.Field) (int, any)
	Delete(input *vehicleModel.Field) (int, any)
	Update(input *vehicleModel.Update) (int, any)
}

type manager struct {
	VehicleService vehicleService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		VehicleService: vehicleService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *vehicleModel.Create) (int, any) {
	defer trx.Rollback()

	vehicleBase, err := m.VehicleService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, vehicleBase.ID)
}

func (m *manager) GetByList(input *vehicleModel.Fields) (int, any) {
	output := &vehicleModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, vehicleBase, err := m.VehicleService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	vehicleByte, err := json.Marshal(vehicleBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(vehicleByte, &output.Vehicles)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, vehicle := range output.Vehicles {
		vehicle.CreatedBy = *vehicleBase[i].CreatedByUsers.Name
		vehicle.UpdatedBy = *vehicleBase[i].UpdatedByUsers.Name
		vehicle.FleetCode = *vehicleBase[i].Fleets.FleetCode
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetByListNoPagination(input *vehicleModel.Field) (int, any) {
	output := &vehicleModel.ListNoPagination{}
	vehicleBase, err := m.VehicleService.GetByListNoPagination(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	vehicleByte, err := json.Marshal(vehicleBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = json.Unmarshal(vehicleByte, &output.Vehicles)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *vehicleModel.Field) (int, any) {
	vehicleBase, err := m.VehicleService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &vehicleModel.Single{}
	vehicleByte, _ := json.Marshal(vehicleBase)
	err = json.Unmarshal(vehicleByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.CreatedBy = *vehicleBase.CreatedByUsers.Name
	output.UpdatedBy = *vehicleBase.UpdatedByUsers.Name
	output.FleetCode = *vehicleBase.Fleets.FleetCode

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *vehicleModel.Field) (int, any) {
	_, err := m.VehicleService.GetBySingle(&vehicleModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.VehicleService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *vehicleModel.Update) (int, any) {
	vehicleBase, err := m.VehicleService.GetBySingle(&vehicleModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.VehicleService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, vehicleBase.ID)
}
