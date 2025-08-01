package special

import (
	"time"

	"fms/internal/interactor/models/page"
	"fms/internal/interactor/models/section"

	"gorm.io/gorm"
)

// Table is the common file of the backend table structure.
type Table struct {
	// 表ID
	ID string `gorm:"<-:create;column:id;type:uuid;not null;primaryKey;" json:"id"`
	// 創建時間
	CreatedAt time.Time `gorm:"<-:create;column:created_at;type:TIMESTAMP;" json:"created_at"`
	// 創建人
	CreatedBy string `gorm:"<-:create;column:created_by;type:uuid;" json:"created_by"`
	// 更新時間
	UpdatedAt *time.Time `gorm:"column:updated_at;type:TIMESTAMP;" json:"updated_at"`
	// 更新人
	UpdatedBy *string `gorm:"column:updated_by;type:uuid;" json:"updated_by"`
	// 刪除時間
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMP;" json:"deleted_at,omitempty"`
}

// Base is the common file of the backend base structure.
type Base struct {
	// 表ID
	ID *string `json:"id,omitempty"`
	// 基本時間
	section.TimeAt
	// 引入page
	page.Pagination
	// 開始結束時間
	section.StartEnd
	// 開始結束時間
	section.ManagementExclusive
	// SQL OrderBy 區段
	OrderBy *string
	// 創建者
	CreatedBy *string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty"`
}
