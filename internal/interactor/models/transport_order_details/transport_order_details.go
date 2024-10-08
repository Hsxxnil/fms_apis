package transport_order_details

import (
	"fms/internal/interactor/models/page"
)

// Create struct is used to create achieves
type Create struct {
	// 品項
	ProductName string `json:"product_name,omitempty" binding:"required" validate:"required"`
	// 單價
	UnitPrice *float64 `json:"unit_price,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 數量
	Quantity *int64 `json:"quantity,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 噸數
	Tonnage *int64 `json:"tonnage,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 托運訂單ID
	TransportOrderID string `json:"transport_order_id,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
	// 板台ID
	TrailerID *string `json:"trailer_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 表ID
	ID string `json:"id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 品項
	ProductName *string `json:"product_name,omitempty" form:"product_name"`
	// 單價
	UnitPrice *float64 `json:"unit_price,omitempty" form:"unit_price"`
	// 數量
	Quantity *int64 `json:"quantity,omitempty" form:"quantity"`
	// 噸數
	Tonnage *int64 `json:"tonnage,omitempty" form:"tonnage"`
	// 托運訂單ID
	TransportOrderID *string `json:"transport_order_id,omitempty" form:"transport_order_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 板台ID
	TrailerID *string `json:"trailer_id,omitempty" form:"trailer_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	// 搜尋結構檔
	Field
	// 分頁搜尋結構檔
	page.Pagination
}

// Single return structure file
type Single struct {
	// 表ID
	ID string `json:"id,omitempty"`
	// 品項
	ProductName string `json:"product_name,omitempty"`
	// 單價
	UnitPrice float64 `json:"unit_price,omitempty"`
	// 數量
	Quantity int64 `json:"quantity,omitempty"`
	// 噸數
	Tonnage int64 `json:"tonnage,omitempty"`
	// 板台ID
	TrailerID string `json:"trailer_id,omitempty"`
	// 板台號碼
	TrailerCode string `json:"trailer_code,omitempty"`
}

// Update struct is used to update achieves
type Update struct {
	// 表ID
	ID string `json:"id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 品項
	ProductName *string `json:"product_name,omitempty"`
	// 單價
	UnitPrice *float64 `json:"unit_price,omitempty"`
	// 數量
	Quantity *int64 `json:"quantity,omitempty"`
	// 噸數
	Tonnage *int64 `json:"tonnage,omitempty"`
	// 板台ID
	TrailerID *string `json:"trailer_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}
