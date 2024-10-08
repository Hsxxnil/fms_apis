package clients

import (
	"fms/internal/entity/postgresql/db/users"
	"fms/internal/interactor/models/special"
)

// Table struct is clients database table struct
type Table struct {
	// 客戶名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
	// 客戶電話
	PhoneNumber string `gorm:"column:phone_number;type:text;" json:"phone_number"`
	// 車隊ID
	FleetID string `gorm:"column:fleet_id;type:uuid;not null;" json:"fleet_id"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:ID;references:CreatedBy" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:ID;references:UpdatedBy" json:"updated_by_users,omitempty"`
	special.Table
}

// Base struct is corresponding to clients table structure file
type Base struct {
	// 車隊ID
	FleetID *string `json:"fleet_id,omitempty"`
	// 客戶名稱
	Name *string `json:"name,omitempty"`
	// 客戶電話
	PhoneNumber *string `json:"phone_number,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.Base
}

// TableName sets the insert table clients for this struct type
func (t *Table) TableName() string {
	return "clients"
}
