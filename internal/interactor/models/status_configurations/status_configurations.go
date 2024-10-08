package status_configurations

import (
	"fms/internal/interactor/models/page"
	"fms/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 狀態ID
	StatusID int64 `json:"status_id,omitempty" binding:"required" validate:"required"`
	// 時間限制
	LimitTime int64 `json:"limit_time,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 車隊ID
	FleetID string `json:"fleet_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 表ID
	ID string `json:"id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 狀態ID
	StatusID *int64 `json:"status_id,omitempty" form:"status_id"`
	// 時間限制
	LimitTime *int64 `json:"limit_time,omitempty" form:"limit_time"`
	// 車隊ID
	FleetID *string `json:"fleet_id,omitempty" form:"fleet_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	// 搜尋結構檔
	Field
	// 分頁搜尋結構檔
	page.Pagination
}

// List is multiple return structure files
type List struct {
	// 多筆
	StatusConfigurations []*struct {
		// 表ID
		ID string `json:"id,omitempty"`
		// 狀態ID
		StatusID int64 `json:"status_id"`
		// 狀態
		Status string `json:"status"`
		// 時間限制
		LimitTime int64 `json:"limit_time,omitempty"`
		// 車隊ID
		FleetID string `json:"fleet_id,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"status_configurations"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 表ID
	ID string `json:"id,omitempty"`
	// 狀態ID
	StatusID int64 `json:"status_id,omitempty"`
	// 狀態
	Status string `json:"status"`
	// 時間限制
	LimitTime int64 `json:"limit_time"`
	// 車隊ID
	FleetID string `json:"fleet_id,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	// 表ID
	ID string `json:"id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 狀態ID
	StatusID *int64 `json:"status_id,omitempty"`
	// 時間限制
	LimitTime *int64 `json:"limit_time,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}
