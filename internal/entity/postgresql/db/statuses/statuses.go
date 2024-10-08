package statuses

import (
	"fms/internal/entity/postgresql/db/users"
	"fms/internal/interactor/models/page"
	"fms/internal/interactor/models/section"
	"gorm.io/gorm"
	"time"
)

// Table struct is statuses database table struct
type Table struct {
	// 表ID
	ID int `gorm:"->;column:id;type:serial;not null;primaryKey;" json:"id"`
	// 狀態
	Status string `gorm:"column:status;type:uuid;not null;" json:"status"`
	// 創建時間
	CreatedAt time.Time `gorm:"<-:create;column:created_at;type:TIMESTAMP;" json:"created_at"`
	// 創建人
	CreatedBy string `gorm:"<-:create;column:created_by;type:uuid;" json:"created_by"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:ID;references:CreatedBy" json:"created_by_users,omitempty"`
	// 更新時間
	UpdatedAt *time.Time `gorm:"column:updated_at;type:TIMESTAMP;" json:"updated_at"`
	// 更新人
	UpdatedBy *string `gorm:"column:updated_by;type:uuid;" json:"updated_by"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:ID;references:UpdatedBy" json:"updated_by_users,omitempty"`
	// 刪除時間
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMP;" json:"deleted_at,omitempty"`
}

// Base struct is corresponding to statuses table structure file
type Base struct {
	// 表ID
	ID *int `json:"id,omitempty"`
	// 狀態
	Status *string `json:"status,omitempty"`
	// 創建者
	CreatedBy *string `json:"created_by,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	// 基本時間
	section.TimeAt
	// 引入page
	page.Pagination
}

// TableName sets the insert table statuses for this struct type
func (t *Table) TableName() string {
	return "statuses"
}
