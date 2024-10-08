package users

import (
	"fms/internal/interactor/models/special"
)

// Table struct is users database table struct
type Table struct {
	// 使用者名稱
	UserName string `gorm:"column:user_name;type:text;not null;" json:"user_name"`
	// 使用者中文名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
	// 使用者密碼
	Password string `gorm:"column:password;type:text;not null;" json:"password"`
	// 使用者電話1
	PhoneNumber1 string `gorm:"column:phone_number1;type:text;" json:"phone_number1"`
	// 使用者電話2
	PhoneNumber2 string `gorm:"column:phone_number2;type:text;" json:"phone_number2"`
	// 使用者電子郵件
	Email string `gorm:"column:email;type:text;" json:"email"`
	// create_users data
	CreatedByUsers *Table `gorm:"foreignKey:CreatedBy" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers *Table `gorm:"foreignKey:UpdatedBy" json:"updated_by_users,omitempty"`
	special.Table
}

// Base struct is corresponding to users table structure file
type Base struct {
	// 使用者名稱
	UserName *string `json:"user_name,omitempty"`
	// 使用者中文名稱
	Name *string `json:"name,omitempty"`
	// 使用者密碼
	Password *string `json:"password,omitempty"`
	// 使用者電話1
	PhoneNumber1 *string `json:"phone_number1,omitempty"`
	// 使用者電話2
	PhoneNumber2 *string `json:"phone_number2,omitempty"`
	// 使用者電子郵件
	Email *string `json:"email,omitempty"`
	// create_users data
	CreatedByUsers *Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers *Base `json:"updated_by_users,omitempty"`
	special.Base
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "users"
}
