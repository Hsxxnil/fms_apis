package client

import (
	"encoding/json"
	"errors"
	clientService "fms/internal/interactor/service/client"

	"fms/internal/interactor/pkg/util"

	clientModel "fms/internal/interactor/models/clients"

	"gorm.io/gorm"

	"fms/internal/interactor/pkg/util/code"
	"fms/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *clientModel.Create) (int, any)
	GetByList(input *clientModel.Fields) (int, any)
	GetBySingle(input *clientModel.Field) (int, any)
	Delete(input *clientModel.Field) (int, any)
	Update(input *clientModel.Update) (int, any)
}

type manager struct {
	ClientService clientService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		ClientService: clientService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *clientModel.Create) (int, any) {
	defer trx.Rollback()

	clientBase, err := m.ClientService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, clientBase.ID)
}

func (m *manager) GetByList(input *clientModel.Fields) (int, any) {
	output := &clientModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, clientBase, err := m.ClientService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	clientByte, err := json.Marshal(clientBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(clientByte, &output.Clients)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, client := range output.Clients {
		client.CreatedBy = *clientBase[i].CreatedByUsers.Name
		client.UpdatedBy = *clientBase[i].UpdatedByUsers.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *clientModel.Field) (int, any) {
	clientBase, err := m.ClientService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &clientModel.Single{}
	clientByte, _ := json.Marshal(clientBase)
	err = json.Unmarshal(clientByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.CreatedBy = *clientBase.CreatedByUsers.Name
	output.UpdatedBy = *clientBase.UpdatedByUsers.Name

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *clientModel.Field) (int, any) {
	_, err := m.ClientService.GetBySingle(&clientModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.ClientService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *clientModel.Update) (int, any) {
	clientBase, err := m.ClientService.GetBySingle(&clientModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.ClientService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, clientBase.ID)
}
