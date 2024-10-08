package status_configurations

import (
	"fms/internal/entity/postgresql/db/statuses"
	"fms/internal/entity/postgresql/db/users"
	"fms/internal/interactor/models/special"
)

// Table struct is status_configurations database table struct
type Table struct {
	// 狀態ID
	StatusID int64 `gorm:"column:status_id;type:int;not null;" json:"status_id"`
	// statuses data
	Statuses statuses.Table `gorm:"foreignKey:ID;references:StatusID" json:"statuses,omitempty"`
	// 時間限制
	LimitTime int64 `gorm:"column:limit_time;type:int;" json:"limit_time"`
	// 車隊ID
	FleetID string `gorm:"column:fleet_id;type:uuid;not null;" json:"fleet_id"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:ID;references:CreatedBy" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:ID;references:UpdatedBy" json:"updated_by_users,omitempty"`
	special.Table
}

// Base struct is corresponding to status_configurations table structure file
type Base struct {
	// 狀態ID
	StatusID *int64 `json:"status_id,omitempty"`
	// statuses data
	Statuses statuses.Base `json:"statuses,omitempty"`
	// 時間限制
	LimitTime *int64 `json:"limit_time,omitempty"`
	// 車隊ID
	FleetID *string `json:"fleet_id,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.Base
}

// TableName sets the insert table status_configurations for this struct type
func (t *Table) TableName() string {
	return "status_configurations"
}
