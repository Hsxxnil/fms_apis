package gps_devices

import (
	"fms/internal/interactor/models/page"
	"fms/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 車機序號
	SID string `json:"sid,omitempty" binding:"required" validate:"required"`
	// 廠商
	Firm string `json:"firm,omitempty"`
	// 型號
	Model string `json:"model,omitempty" binding:"required" validate:"required"`
	// 門號
	PhoneNumber string `json:"phone_number,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 表ID
	ID string `json:"id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 車機序號
	SID *string `json:"sid,omitempty" form:"sid"`
	// 廠商
	Firm *string `json:"firm,omitempty" form:"firm"`
	// 型號
	Model *string `json:"model,omitempty" form:"model"`
	// 門號
	PhoneNumber *string `json:"phone_number,omitempty" form:"phone_number"`
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
	GpsDevices []*struct {
		// 表ID
		ID string `json:"id,omitempty"`
		// 車機序號
		SID string `json:"sid,omitempty"`
		// 廠商
		Firm string `json:"firm,omitempty"`
		// 型號
		Model string `json:"model,omitempty"`
		// 門號
		PhoneNumber string `json:"phone_number,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"gps_devices"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 表ID
	ID string `json:"id,omitempty"`
	// 車機序號
	SID string `json:"sid,omitempty"`
	// 廠商
	Firm string `json:"firm,omitempty"`
	// 型號
	Model string `json:"model,omitempty"`
	// 門號
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
	// 廠商
	Firm *string `json:"firm,omitempty"`
	// 型號
	Model *string `json:"model,omitempty"`
	// 門號
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}
