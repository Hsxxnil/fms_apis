package fleets

import (
	"fms/internal/interactor/models/page"
	"fms/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 車隊代碼
	FleetCode string `json:"fleet_code,omitempty" binding:"required" validate:"required"`
	// 車隊中文名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 表ID
	ID string `json:"id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 車隊代碼
	FleetCode *string `json:"fleet_code,omitempty" form:"fleet_code"`
	// 車隊中文名稱
	Name *string `json:"name,omitempty" form:"name"`
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
	Fleets []*struct {
		// 表ID
		ID string `json:"id,omitempty"`
		// 車隊代碼
		FleetCode string `json:"fleet_code,omitempty"`
		// 車隊中文名稱
		Name string `json:"name,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"fleets"`
	// 分頁返回結構檔
	page.Total
}

// ListNoPagination is multiple return structure files (no pagination)
type ListNoPagination struct {
	// 多筆
	Fleets []*struct {
		// 表ID
		ID string `json:"id,omitempty"`
		// 車隊代碼
		FleetCode string `json:"fleet_code,omitempty"`
		// 車隊中文名稱
		Name string `json:"name,omitempty"`
	} `json:"fleets"`
}

// Single return structure file
type Single struct {
	// 表ID
	ID string `json:"id,omitempty"`
	// 車隊代碼
	FleetCode string `json:"fleet_code,omitempty"`
	// 車隊中文名稱
	Name string `json:"name,omitempty"`
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
	// 車隊代碼
	FleetCode *string `json:"fleet_code,omitempty"`
	// 車隊中文名稱
	Name *string `json:"name,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}
