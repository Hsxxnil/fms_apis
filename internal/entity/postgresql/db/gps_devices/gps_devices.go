package gps_devices

import (
	"fms/internal/entity/postgresql/db/users"
	"fms/internal/interactor/models/special"
)

// Table struct is gps_devices database table struct
type Table struct {
	// 車機序號
	SID string `gorm:"<-:create;column:sid;type:text;not null;" json:"sid"`
	// 廠商
	Firm string `gorm:"column:firm;type:text;" json:"firm"`
	// 型號
	Model string `gorm:"column:model;type:text;not null;" json:"model"`
	// 門號
	PhoneNumber string `gorm:"column:phone_number;type:text;" json:"phone_number"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:ID;references:CreatedBy" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:ID;references:UpdatedBy" json:"updated_by_users,omitempty"`
	special.Table
}

// Base struct is corresponding to gps_devices table structure file
type Base struct {
	// 車機序號
	SID *string `json:"sid,omitempty"`
	// 廠商
	Firm *string `json:"firm,omitempty"`
	// 型號
	Model *string `json:"model,omitempty"`
	// 門號
	PhoneNumber *string `json:"phone_number,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.Base
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "gps_devices"
}
