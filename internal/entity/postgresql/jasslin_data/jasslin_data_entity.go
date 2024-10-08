package jasslin_data

import (
	model "fms/internal/entity/postgresql/db/jasslin_data"
	"fms/internal/interactor/pkg/util/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Entity interface {
	WithTrx(trx *gorm.DB) Entity
	GetByList(input *model.Base) (quantity int64, output []*model.Table, err error)
	GetBySIDList(input *model.Base) (output []*model.Table, err error)
	GetByLastData(input *model.Base) (output []*model.Table, err error)
	GetBySingle(input *model.Base) (output *model.Table, err error)
	GetByQuantity(input *model.Base) (quantity int64, err error)
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

func (s *storage) GetByList(input *model.Base) (quantity int64, output []*model.Table, err error) {
	query := s.db.Model(&model.Table{}).Preload(clause.Associations)
	if input.ID != nil {
		query.Where("id = ?", input.ID)
	}

	if input.SID != nil {
		query.Where("sid = ?", *input.SID)
	}

	err = query.Count(&quantity).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("date_time desc").Find(&output).Error
	if err != nil {
		log.Error(err)
		return 0, nil, err
	}

	return quantity, output, nil
}

func (s *storage) GetBySIDList(input *model.Base) (output []*model.Table, err error) {
	subQuery := s.db.Model(&model.Table{}).Select("distinct on (date_time) *").Order("date_time desc").Limit(10)
	subQuery1 := s.db.Table("(?)", subQuery).Select("min(date_time)")
	query := s.db.Model(&model.Table{}).Preload(clause.Associations)
	if input.ID != nil {
		query.Where("id = ?", input.ID)
	}

	if input.SID != nil {
		query.Where("sid = ?", *input.SID)
		subQuery.Where("sid = ?", *input.SID)
	}

	// filter
	filter := s.db.Model(&model.Table{})
	if input.FilterStartTime != nil && input.FilterEndTime != nil {
		subQuery.Where("date_time <= ?", input.FilterStartTime)
		filter.Where("date_time between (? as sub) and ?", subQuery1, input.FilterEndTime)
	}

	query.Where(filter)
	err = query.Select("distinct on (date_time) *").Order("date_time asc").Find(&output).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (s *storage) GetByLastData(input *model.Base) (output []*model.Table, err error) {
	subQuery := s.db.Model(&model.Table{}).Select("distinct on (date_time) *").Order("date_time desc").Limit(10)
	subQuery1 := s.db.Table("(?)", subQuery).Select("min(date_time)")
	query := s.db.Model(&model.Table{}).Preload(clause.Associations)

	if input.ID != nil {
		query.Where("id = ?", input.ID)
		subQuery.Where("id = ?", input.ID)
	}

	if input.SID != nil {
		query.Where("sid = ?", *input.SID)
		subQuery.Where("sid = ?", *input.SID)
	}

	err = query.Select("distinct on (date_time) *").Where("date_time >= (? as sub)", subQuery1).Order("date_time desc").Find(&output).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (s *storage) GetBySingle(input *model.Base) (output *model.Table, err error) {
	query := s.db.Model(&model.Table{}).Preload(clause.Associations)
	if input.ID != nil {
		query.Where("id = ?", input.ID)
	}

	if input.SID != nil {
		query.Where("sid = ?", *input.SID)
	}

	err = query.Order("date_time desc").First(&output).Error
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

	if input.SID != nil {
		query.Where("sid = ?", *input.SID)
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

	if input.Address != nil {
		data["address"] = input.Address
	}

	if input.Status != nil {
		data["status"] = input.Status
	}

	if input.StatusTime != nil {
		data["status_time"] = input.StatusTime
	}

	if input.ID != nil {
		query.Where("id = ?", input.ID)
	}

	if input.SID != nil {
		query.Where("sid = ?", *input.SID)
	}

	if input.ServerTime != nil {
		query.Where("server_time = ?", input.ServerTime)
	}

	err = query.Select("*").Updates(data).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
