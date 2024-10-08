package user_roles

import (
	"fms/internal/entity/postgresql/db/fleets"
	"fms/internal/entity/postgresql/db/roles"
	"fms/internal/interactor/models/special"
)

// Table struct is user_roles database table struct
type Table struct {
	// 使用者ID
	UserID string `gorm:"column:user_id;type:uuid;not null;" json:"user_id"`
	// 角色ID
	RoleID string `gorm:"column:role_id;type:uuid;not null;" json:"role_id"`
	// roles data
	Roles roles.Table `gorm:"foreignKey:ID;references:RoleID" json:"roles,omitempty"`
	// 車隊ID
	FleetID string `gorm:"column:fleet_id;type:uuid;not null;" json:"fleet_id"`
	// fleets data
	Fleets fleets.Table `gorm:"foreignKey:ID;references:FleetID" json:"fleets,omitempty"`
	special.Table
}

// Base struct is corresponding to user_roles table structure file
type Base struct {
	// 使用者ID
	UserID *string `json:"user_id,omitempty"`
	// 角色ID
	RoleID *string `json:"role_id,omitempty"`
	// roles data
	Roles roles.Base `json:"roles,omitempty"`
	// 車隊ID
	FleetID *string `json:"fleet_id,omitempty"`
	// fleets data
	Fleets fleets.Base `json:"fleets,omitempty"`
	special.Base
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "user_roles"
}
