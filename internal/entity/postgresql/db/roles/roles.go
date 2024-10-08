package roles

import (
	"fms/internal/entity/postgresql/db/users"
	"fms/internal/interactor/models/special"
)

// Table struct is roles database table struct
type Table struct {
	// 角色名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
	// 角色顯示名稱
	DisplayName string `gorm:"column:display_name;type:text;not null;" json:"display_name"`
	// 角色是否啟用
	IsEnable bool `gorm:"column:is_enable;type:bool;not null;" json:"is_enable"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:ID;references:CreatedBy" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:ID;references:UpdatedBy" json:"updated_by_users,omitempty"`
	special.Table
}

// Base struct is corresponding to roles table structure file
type Base struct {
	// 角色名稱
	Name *string `json:"name,omitempty"`
	// 角色顯示名稱
	DisplayName *string `json:"display_name,omitempty"`
	// 角色是否啟用
	IsEnable *bool `json:"is_enable,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.Base
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "roles"
}
