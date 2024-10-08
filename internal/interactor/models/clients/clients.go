package clients

import (
	"fms/internal/interactor/models/page"
	"fms/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 車隊ID
	FleetID string `json:"fleet_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 客戶名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 客戶電話
	PhoneNumber string `json:"phone_number,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 表ID
	ID string `json:"id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 客戶名稱
	Name *string `json:"name,omitempty" form:"name"`
	// 客戶電話
	PhoneNumber *string `json:"phone_number,omitempty" form:"phone_number"`
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
	Clients []*struct {
		// 表ID
		ID string `json:"id,omitempty"`
		// 客戶名稱
		Name string `json:"name"`
		// 客戶電話
		PhoneNumber string `json:"phone_number,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"clients"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 表ID
	ID string `json:"id,omitempty"`
	// 客戶名稱
	Name string `json:"name,omitempty"`
	// 客戶電話
	PhoneNumber string `json:"phone_number,omitempty"`
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
	// 車隊ID
	FleetID *string `json:"fleet_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 客戶名稱
	Name *string `json:"name,omitempty"`
	// 客戶電話
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}
