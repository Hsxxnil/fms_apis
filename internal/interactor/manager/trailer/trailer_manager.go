package trailer

import (
	"encoding/json"
	"errors"
	trailerService "fms/internal/interactor/service/trailer"

	"fms/internal/interactor/pkg/util"

	trailerModel "fms/internal/interactor/models/trailers"

	"gorm.io/gorm"

	"fms/internal/interactor/pkg/util/code"
	"fms/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *trailerModel.Create) (int, any)
	GetByList(input *trailerModel.Fields) (int, any)
	GetBySingle(input *trailerModel.Field) (int, any)
	Delete(input *trailerModel.Field) (int, any)
	Update(input *trailerModel.Update) (int, any)
}

type manager struct {
	TrailerService trailerService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		TrailerService: trailerService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *trailerModel.Create) (int, any) {
	defer trx.Rollback()

	// 判斷板台號碼是否重複
	quantity, _ := m.TrailerService.GetByQuantity(&trailerModel.Field{
		Code: util.PointerString(input.Code),
	})

	if quantity > 0 {
		log.Info("Code already exists. Code: ", input.Code)
		return code.BadRequest, code.GetCodeMessage(code.BadRequest, "Trailer already exists.")
	}

	trailerBase, err := m.TrailerService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, trailerBase.ID)
}

func (m *manager) GetByList(input *trailerModel.Fields) (int, any) {
	output := &trailerModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, trailerBase, err := m.TrailerService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	trailerByte, err := json.Marshal(trailerBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(trailerByte, &output.Trailers)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, trailer := range output.Trailers {
		trailer.CreatedBy = *trailerBase[i].CreatedByUsers.Name
		trailer.UpdatedBy = *trailerBase[i].UpdatedByUsers.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *trailerModel.Field) (int, any) {
	trailerBase, err := m.TrailerService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &trailerModel.Single{}
	trailerByte, _ := json.Marshal(trailerBase)
	err = json.Unmarshal(trailerByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.CreatedBy = *trailerBase.CreatedByUsers.Name
	output.UpdatedBy = *trailerBase.UpdatedByUsers.Name

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *trailerModel.Field) (int, any) {
	_, err := m.TrailerService.GetBySingle(&trailerModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.TrailerService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *trailerModel.Update) (int, any) {
	trailerBase, err := m.TrailerService.GetBySingle(&trailerModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.TrailerService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, trailerBase.ID)
}
