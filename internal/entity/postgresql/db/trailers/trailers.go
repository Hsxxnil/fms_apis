package trailers

import (
	"fms/internal/entity/postgresql/db/users"
	"fms/internal/interactor/models/special"
)

// Table struct is trailers database table struct
type Table struct {
	// 板台號碼
	Code string `gorm:"column:code;type:text;not null;" json:"code"`
	// 車隊ID
	FleetID string `gorm:"column:fleet_id;type:uuid;not null;" json:"fleet_id"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:ID;references:CreatedBy" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:ID;references:UpdatedBy" json:"updated_by_users,omitempty"`
	special.Table
}

// Base struct is corresponding to trailers table structure file
type Base struct {
	// 車隊ID
	FleetID *string `json:"fleet_id,omitempty"`
	// 板台號碼
	Code *string `json:"code,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.Base
}

// TableName sets the insert table trailers for this struct type
func (t *Table) TableName() string {
	return "trailers"
}
