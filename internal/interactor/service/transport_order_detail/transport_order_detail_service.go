package transport_order_detail

import (
	"encoding/json"

	db "fms/internal/entity/postgresql/db/transport_order_details"
	store "fms/internal/entity/postgresql/transport_order_detail"
	model "fms/internal/interactor/models/transport_order_details"
	"fms/internal/interactor/pkg/util"
	"fms/internal/interactor/pkg/util/log"
	"fms/internal/interactor/pkg/util/uuid"

	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Create(input []*model.Create) (output []*db.Base, err error)
	GetByList(input *model.Fields) (quantity int64, output []*db.Base, err error)
	GetByListNoPagination(input *model.Field) (output []*db.Base, err error)
	GetBySingle(input *model.Field) (output *db.Base, err error)
	GetByQuantity(input *model.Field) (quantity int64, err error)
	Update(input *model.Update) (err error)
	Delete(input *model.Field) (err error)
}

type service struct {
	Repository store.Entity
}

func Init(db *gorm.DB) Service {
	return &service{
		Repository: store.Init(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {
	return &service{
		Repository: s.Repository.WithTrx(tx),
	}
}

func (s *service) Create(input []*model.Create) (output []*db.Base, err error) {
	var base []*db.Base
	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = json.Unmarshal(marshal, &base)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	for i, field := range base {
		field.ID = util.PointerString(uuid.CreatedUUIDString())
		field.CreatedAt = util.PointerTime(util.NowToUTC())
		field.UpdatedAt = util.PointerTime(util.NowToUTC())
		field.UpdatedBy = util.PointerString(input[i].CreatedBy)
	}

	err = s.Repository.Create(base)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	marshal, err = json.Marshal(base)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)

		return nil, err
	}

	return output, nil
}

func (s *service) GetByList(input *model.Fields) (quantity int64, output []*db.Base, err error) {
	field := &db.Base{}
	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		return 0, nil, err
	}

	err = json.Unmarshal(marshal, &field)
	if err != nil {
		log.Error(err)
		return 0, nil, err
	}

	quantity, fields, err := s.Repository.GetByList(field)
	if err != nil {
		log.Error(err)
		return 0, output, err
	}

	marshal, err = json.Marshal(fields)
	if err != nil {
		log.Error(err)
		return 0, nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return 0, nil, err
	}

	return quantity, output, nil
}

func (s *service) GetByListNoPagination(input *model.Field) (output []*db.Base, err error) {
	field := &db.Base{}
	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = json.Unmarshal(marshal, &field)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	fields, err := s.Repository.GetByListNoPagination(field)
	if err != nil {
		log.Error(err)
		return output, err
	}

	marshal, err = json.Marshal(fields)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (s *service) GetBySingle(input *model.Field) (output *db.Base, err error) {
	field := &db.Base{}
	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = json.Unmarshal(marshal, &field)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	single, err := s.Repository.GetBySingle(field)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	marshal, err = json.Marshal(single)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (s *service) Delete(input *model.Field) (err error) {
	field := &db.Base{}
	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		return err
	}

	err = json.Unmarshal(marshal, &field)
	if err != nil {
		log.Error(err)
		return err
	}

	err = s.Repository.Delete(field)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *service) Update(input *model.Update) (err error) {
	field := &db.Base{}
	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		return err
	}

	err = json.Unmarshal(marshal, &field)
	if err != nil {
		log.Error(err)
		return err
	}

	err = s.Repository.Update(field)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *service) GetByQuantity(input *model.Field) (quantity int64, err error) {
	field := &db.Base{}
	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	err = json.Unmarshal(marshal, &field)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	quantity, err = s.Repository.GetByQuantity(field)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return quantity, nil
}
