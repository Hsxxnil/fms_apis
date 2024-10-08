package drivers

import (
	"fms/internal/interactor/models/page"
	"fms/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 車隊ID
	FleetID string `json:"fleet_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 姓名
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 身份證字號
	IDCardNumber string `json:"id_card_number,omitempty"`
	// 員工編號
	EmployeeNumber string `json:"employee_number,omitempty"`
	// 行動電話
	PhoneNumber string `json:"phone_number,omitempty"`
	// 電子郵件
	Email string `json:"email,omitempty"`
	// 連絡地址
	Address string `json:"address,omitempty"`
	// 每日成本
	DailyCost int64 `json:"daily_cost,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 表ID
	ID string `json:"id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 姓名
	Name *string `json:"name,omitempty" form:"name"`
	// 身份證字號
	IDCardNumber *string `json:"id_card_number,omitempty" form:"id_card_number"`
	// 員工編號
	EmployeeNumber *string `json:"employee_number,omitempty" form:"employee_number"`
	// 行動電話
	PhoneNumber *string `json:"phone_number,omitempty" form:"phone_number"`
	// 電子郵件
	Email *string `json:"email,omitempty" form:"email"`
	// 連絡地址
	Address *string `json:"address,omitempty" form:"address"`
	// 每日成本
	DailyCost *int64 `json:"daily_cost,omitempty" form:"daily_cost"`
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
	Drivers []*struct {
		// 表ID
		ID string `json:"id,omitempty"`
		// 姓名
		Name string `json:"name,omitempty"`
		// 身份證字號
		IDCardNumber string `json:"id_card_number,omitempty"`
		// 員工編號
		EmployeeNumber string `json:"employee_number,omitempty"`
		// 行動電話
		PhoneNumber string `json:"phone_number,omitempty"`
		// 電子郵件
		Email string `json:"email,omitempty"`
		// 連絡地址
		Address string `json:"address,omitempty"`
		// 每日成本
		DailyCost int64 `json:"daily_cost,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"drivers"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 表ID
	ID string `json:"id,omitempty"`
	// 姓名
	Name string `json:"name,omitempty"`
	// 身份證字號
	IDCardNumber string `json:"id_card_number,omitempty"`
	// 員工編號
	EmployeeNumber string `json:"employee_number,omitempty"`
	// 行動電話
	PhoneNumber string `json:"phone_number,omitempty"`
	// 電子郵件
	Email string `json:"email,omitempty"`
	// 連絡地址
	Address string `json:"address,omitempty"`
	// 每日成本
	DailyCost int64 `json:"daily_cost,omitempty"`
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
	DailyCost *int64 `json:"daily_cost,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}
