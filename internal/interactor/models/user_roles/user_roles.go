package user_roles

import (
	"fms/internal/interactor/models/page"
	"fms/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 使用者ID
	UserID string `json:"user_id,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
	// 角色ID
	RoleID string `json:"role_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 車隊ID
	FleetID string `json:"fleet_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 表ID
	ID string `json:"id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 使用者ID
	UserID *string `json:"user_id,omitempty" form:"user_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 角色ID
	RoleID string `json:"role_id,omitempty" form:"role_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 車隊ID
	FleetID string `json:"fleet_id,omitempty" form:"fleet_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
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
	UserRoles []*struct {
		// 表ID
		ID string `json:"id,omitempty"`
		// 使用者ID
		UserID string `json:"user_id,omitempty"`
		// 車隊ID
		FleetID string `json:"fleet_id,omitempty"`
		// 角色ID
		RoleID string `json:"role_id,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"user_roles"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 表ID
	ID string `json:"id,omitempty"`
	// 使用者ID
	UserID string `json:"user_id,omitempty"`
	// 車隊ID
	FleetID string `json:"fleet_id,omitempty"`
	// 角色ID
	RoleID string `json:"role_id,omitempty"`
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
	// 使用者ID
	UserID *string `json:"user_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 角色ID
	RoleID *string `json:"role_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 車隊ID
	FleetID *string `json:"fleet_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}
