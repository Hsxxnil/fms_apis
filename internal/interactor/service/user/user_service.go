package user

import (
	"encoding/json"

	db "fms/internal/entity/postgresql/db/users"
	store "fms/internal/entity/postgresql/user"
	"fms/internal/interactor/pkg/util/encryption"
	"fms/internal/interactor/pkg/util/hash"

	model "fms/internal/interactor/models/users"
	"fms/internal/interactor/pkg/util"
	"fms/internal/interactor/pkg/util/log"
	"fms/internal/interactor/pkg/util/uuid"

	"gorm.io/gorm"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Create(input *model.Create) (output *db.Base, err error)
	GetByList(input *model.Fields) (quantity int64, output []*db.Base, err error)
	GetBySingle(input *model.Field) (output *db.Base, err error)
	GetByQuantity(input *model.Field) (quantity int64, err error)
	Update(input *model.Update) (err error)
	Delete(input *model.Field) (err error)
	AcknowledgeUser(input *model.Field) (acknowledge bool, output *db.Base, err error)
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

func (s *service) Create(input *model.Create) (output *db.Base, err error) {
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

	key := "423CD5C09F7DD58950F1E494099EB075"
	input.Password = hash.HmacSha512(input.Password, key)
	password, err := encryption.AesEncryptOFB([]byte(input.Password), []byte(key))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	field.Password = util.PointerString(hash.Base64BydEncode(password))
	field.ID = util.PointerString(uuid.CreatedUUIDString())
	field.CreatedAt = util.PointerTime(util.NowToUTC())
	field.UpdatedAt = util.PointerTime(util.NowToUTC())
	field.UpdatedBy = util.PointerString(input.CreatedBy)
	err = s.Repository.Create(field)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	marshal, err = json.Marshal(field)
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

	key := "423CD5C09F7DD58950F1E494099EB075"
	if input.Password != nil {
		input.Password = util.PointerString(hash.HmacSha512(*input.Password, key))
		password, err := encryption.AesEncryptOFB([]byte(*input.Password), []byte(key))
		if err != nil {
			log.Error(err)
			return err
		}
		field.Password = util.PointerString(hash.Base64BydEncode(password))
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

func (s *service) AcknowledgeUser(input *model.Field) (acknowledge bool, output *db.Base, err error) {
	field := &db.Base{}
	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		return false, nil, err
	}

	err = json.Unmarshal(marshal, &field)
	if err != nil {
		log.Error(err)
		return false, nil, err
	}

	userBase, err := s.Repository.GetBySingle(field)
	if err != nil {
		log.Error(err)
		return false, output, err
	}

	if userBase == nil {
		//return false, nil, errors.New("Incorrect user_name or fleet_code")
		return false, nil, nil
	}

	key := "423CD5C09F7DD58950F1E494099EB075"
	input.Password = util.PointerString(hash.HmacSha512(*input.Password, key))
	password, err := encryption.AesDecryptOFB(hash.Base64StdDecode(userBase.Password), []byte(key))
	if err != nil {
		log.Error(err)
		return false, nil, err
	}

	if string(password) != *input.Password {
		//return false, nil, errors.New("Incorrect password")
		return false, nil, nil
	}

	marshal, err = json.Marshal(userBase)
	if err != nil {
		log.Error(err)
		return false, nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return false, nil, err
	}

	return true, output, nil
}
