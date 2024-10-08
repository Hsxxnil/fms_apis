package drivers

import (
	"fms/internal/entity/postgresql/db/users"
	"fms/internal/interactor/models/special"
)

// Table struct is drivers database table struct
type Table struct {
	// 姓名
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
	// 身份證字號
	IDCardNumber string `gorm:"column:id_card_number;type:text;" json:"id_card_number"`
	// 員工編號
	EmployeeNumber string `gorm:"column:employee_number;type:text;" json:"employee_number"`
	// 行動電話
	PhoneNumber string `gorm:"column:phone_number;type:text;" json:"phone_number"`
	// 電子郵件
	Email string `gorm:"column:email;type:text;" json:"email"`
	// 連絡地址
	Address string `gorm:"column:address;type:text;" json:"address"`
	// 每日成本
	DailyCost int64 `gorm:"column:daily_cost;type:int;" json:"daily_cost"`
	// 車隊ID
	FleetID string `gorm:"column:fleet_id;type:uuid;not null;" json:"fleet_id"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:ID;references:CreatedBy" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:ID;references:UpdatedBy" json:"updated_by_users,omitempty"`
	special.Table
}

// Base struct is corresponding to drivers table structure file
type Base struct {
	// 車隊ID
	FleetID *string `json:"fleet_id,omitempty"`
	// 姓名
	Name *string `json:"name,omitempty"`
	// 身份證字號
	IDCardNumber *string `json:"id_card_number,omitempty"`
	// 員工編號
	EmployeeNumber *string `json:"employee_number,omitempty"`
	// 行動電話
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 電子郵件
	Email *string `json:"email,omitempty"`
	// 連絡地址
	Address *string `json:"address,omitempty"`
	// 每日成本
	DailyCost *int64 `json:"daily_cost,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.Base
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "drivers"
}
