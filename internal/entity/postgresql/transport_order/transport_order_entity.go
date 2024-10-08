package transport_order

import (
	"encoding/json"

	model "fms/internal/entity/postgresql/db/transport_orders"
	"fms/internal/interactor/pkg/util/log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Entity interface {
	WithTrx(trx *gorm.DB) Entity
	Create(input *model.Base) (err error)
	GetByList(input *model.Base) (quantity int64, output []*model.Table, err error)
	GetByListNoPagination(input *model.Base) (output []*model.Table, err error)
	GetBySingle(input *model.Base) (output *model.Table, err error)
	GetByQuantity(input *model.Base) (quantity int64, err error)
	Delete(input *model.Base) (err error)
	Update(input *model.Base) (err error)
}

type storage struct {
	db *gorm.DB
}

func Init(db *gorm.DB) Entity {
	return &storage{
		db: db,
	}
}

func (s *storage) WithTrx(trx *gorm.DB) Entity {
	return &storage{
		db: trx,
	}
}

func (s *storage) Create(input *model.Base) (err error) {
	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		return err
	}

	data := &model.Table{}
	err = json.Unmarshal(marshal, data)
	if err != nil {
		log.Error(err)
		return err
	}

	err = s.db.Model(&model.Table{}).Omit(clause.Associations).Create(&data).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *storage) GetByList(input *model.Base) (quantity int64, output []*model.Table, err error) {
	query := s.db.Model(&model.Table{}).
		Preload("TransportOrderDetails.Trailers").
		Preload(clause.Associations)

	if input.ID != nil {
		query.Where("id = ?", input.ID)
	}

	if input.FleetID != nil {
		query.Where("fleet_id = ?", input.FleetID)
	}

	err = query.Count(&quantity).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("created_at desc").Find(&output).Error
	if err != nil {
		log.Error(err)
		return 0, nil, err
	}

	return quantity, output, nil
}

func (s *storage) GetByListNoPagination(input *model.Base) (output []*model.Table, err error) {
	query := s.db.Model(&model.Table{})
	if input.ID != nil {
		query.Where("id = ?", input.ID)
	}

	if input.FleetID != nil {
		query.Where("fleet_id = ?", input.FleetID)
	}

	err = query.Order("created_at desc").Find(&output).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (s *storage) GetBySingle(input *model.Base) (output *model.Table, err error) {
	query := s.db.Model(&model.Table{}).
		Preload("TransportOrderDetails.Trailers").
		Preload(clause.Associations)

	if input.ID != nil {
		query.Where("id = ?", input.ID)
	}

	if input.FleetID != nil {
		query.Where("fleet_id = ?", input.FleetID)
	}

	err = query.First(&output).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (s *storage) GetByQuantity(input *model.Base) (quantity int64, err error) {
	query := s.db.Model(&model.Table{})
	if input.ID != nil {
		query.Where("id = ?", input.ID)
	}

	if input.FleetID != nil {
		query.Where("fleet_id = ?", input.FleetID)
	}

	err = query.Count(&quantity).Select("*").Error
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return quantity, nil
}

func (s *storage) Update(input *model.Base) (err error) {
	query := s.db.Model(&model.Table{}).Omit(clause.Associations)
	data := map[string]any{}

	if input.Name != nil {
		data["name"] = input.Name
	}

	if input.Code != nil {
		data["code"] = input.Code
	}

	if input.Origin != nil {
		data["origin"] = input.Origin
	}

	if input.Destination != nil {
		data["destination"] = input.Destination
	}

	if input.ClientID != nil {
		data["client_id"] = input.ClientID
	}

	if input.Deadline != nil {
		data["deadline"] = input.Deadline
	}

	if input.TransportTaskID != nil {
		data["transport_task_id"] = input.TransportTaskID
	} else {
		data["transport_task_id"] = nil
	}

	if input.Sequence != nil {
		data["sequence"] = input.Sequence
	} else {
		data["sequence"] = nil
	}

	if input.UpdatedBy != nil {
		data["updated_by"] = input.UpdatedBy
	}

	if input.ID != nil {
		query.Where("id = ?", input.ID)
	}

	if input.IDs != nil {
		query.Where("id in (?)", input.IDs)
	}

	err = query.Select("*").Updates(data).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *storage) Delete(input *model.Base) (err error) {
	query := s.db.Model(&model.Table{}).Omit(clause.Associations)
	if input.ID != nil {
		query.Where("id = ?", input.ID)
	}

	err = query.Delete(&model.Table{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
