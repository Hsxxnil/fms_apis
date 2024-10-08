package fleets

import (
	"fms/internal/entity/postgresql/db/users"
	"fms/internal/interactor/models/special"
)

// Table struct is fleets database table struct
type Table struct {
	// 車隊代碼
	FleetCode string `gorm:"column:fleet_code;type:text;not null;" json:"fleet_code"`
	// 車隊中文名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:ID;references:CreatedBy" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:ID;references:UpdatedBy" json:"updated_by_users,omitempty"`
	special.Table
}

// Base struct is corresponding to fleets table structure file
type Base struct {
	// 車隊代碼
	FleetCode *string `json:"fleet_code,omitempty"`
	// 車隊中文名稱
	Name *string `json:"name,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.Base
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "fleets"
}
