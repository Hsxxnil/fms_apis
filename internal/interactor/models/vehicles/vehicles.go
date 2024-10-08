package vehicles

import (
	"fms/internal/interactor/models/page"
	"fms/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 車隊ID
	FleetID string `json:"fleet_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 車輛名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 車輛司機
	Driver string `json:"driver,omitempty"`
	// 車輛牌照
	LicensePlate string `json:"license_plate,omitempty" binding:"required" validate:"required"`
	// 車機序號
	SID string `json:"sid,omitempty" binding:"required" validate:"required"`
	// 油耗量(公里/公升)
	FuelConsumption float64 `json:"fuel_consumption,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 目前里程
	CurrentMileage int64 `json:"current_mileage,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 車行統編
	TaxID string `json:"tax_id,omitempty"`
	// 油耗種類
	FuelType string `json:"fuel_type,omitempty"`
	// ETC計費車種
	BillingType string `json:"billing_type,omitempty"`
	// 車種
	Style string `json:"style,omitempty"`
	// 車重(噸)
	Weight string `json:"weight,omitempty"`
	// 每日成本
	DailyCost int64 `json:"daily_cost,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 表ID
	ID string `json:"id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 車隊ID
	FleetID *string `json:"fleet_id,omitempty" form:"fleet_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 車輛名稱
	Name *string `json:"name,omitempty" form:"name"`
	// 車輛司機
	Driver *string `json:"driver,omitempty" form:"driver"`
	// 車輛牌照
	LicensePlate *string `json:"license_plate,omitempty" form:"license_plate"`
	// 車機序號
	SID *string `json:"sid,omitempty" form:"sid"`
	// 油耗量(公里/公升)
	FuelConsumption *float64 `json:"fuel_consumption,omitempty" form:"fuel_consumption"`
	// 目前里程
	CurrentMileage *int64 `json:"current_mileage,omitempty" form:"current_mileage"`
	// 車行統編
	TaxID *string `json:"tax_id,omitempty" form:"tax_id"`
	// 油耗種類
	FuelType *string `json:"fuel_type,omitempty" form:"fuel_type"`
	// ETC計費車種
	BillingType *string `json:"billing_type,omitempty" form:"billing_type"`
	// 車種
	Style *string `json:"style,omitempty" form:"style"`
	// 車重(噸)
	Weight *string `json:"weight,omitempty" form:"weight"`
	// 每日成本
	DailyCost *int64 `json:"daily_cost,omitempty" form:"daily_cost"`
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
	Vehicles []*struct {
		// 表ID
		ID string `json:"id,omitempty"`
		// 車隊ID
		FleetID string `json:"fleet_id,omitempty"`
		// 車隊代碼
		FleetCode string `json:"fleet_code,omitempty"`
		// 車輛名稱
		Name string `json:"name,omitempty"`
		// 車輛司機
		Driver string `json:"driver,omitempty"`
		// 車輛牌照
		LicensePlate string `json:"license_plate,omitempty"`
		// 車機序號
		SID string `json:"sid,omitempty"`
		// 油耗量(公里/公升)
		FuelConsumption float64 `json:"fuel_consumption,omitempty"`
		// 目前里程
		CurrentMileage int64 `json:"current_mileage,omitempty"`
		// 車行統編
		TaxID string `json:"tax_id,omitempty"`
		// 油耗種類
		FuelType string `json:"fuel_type,omitempty"`
		// ETC計費車種
		BillingType string `json:"billing_type,omitempty"`
		// 車種
		Style string `json:"style,omitempty"`
		// 車重(噸)
		Weight string `json:"weight,omitempty"`
		// 每日成本
		DailyCost int64 `json:"daily_cost,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"vehicles"`
	// 分頁返回結構檔
	page.Total
}

// ListNoPagination is multiple return structure files (no pagination)
type ListNoPagination struct {
	// 多筆
	Vehicles []*struct {
		// 表ID
		ID string `json:"id,omitempty"`
		// 車輛名稱
		Name string `json:"name,omitempty"`
		// 車輛牌照
		LicensePlate string `json:"license_plate,omitempty"`
	}
}

// Single return structure file
type Single struct {
	// 表ID
	ID string `json:"id,omitempty"`
	// 車隊ID
	FleetID string `json:"fleet_id,omitempty"`
	// 車隊代碼
	FleetCode string `json:"fleet_code,omitempty"`
	// 車輛名稱
	Name string `json:"name,omitempty"`
	// 車輛司機
	Driver string `json:"driver,omitempty"`
	// 車輛牌照
	LicensePlate string `json:"license_plate,omitempty"`
	// 車機序號
	SID string `json:"sid,omitempty"`
	// 油耗量(公里/公升)
	FuelConsumption float64 `json:"fuel_consumption,omitempty"`
	// 目前里程
	CurrentMileage int64 `json:"current_mileage,omitempty"`
	// 車行統編
	TaxID string `json:"tax_id,omitempty"`
	// 油耗種類
	FuelType string `json:"fuel_type,omitempty"`
	// ETC計費車種
	BillingType string `json:"billing_type,omitempty"`
	// 車種
	Style string `json:"style,omitempty"`
	// 車重(噸)
	Weight string `json:"weight,omitempty"`
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
	// 車輛名稱
	Name *string `json:"name,omitempty"`
	// 車輛司機
	Driver *string `json:"driver,omitempty"`
	// 車機序號
	SID *string `json:"sid,omitempty"`
	// 油耗量(公里/公升)
	FuelConsumption *float64 `json:"fuel_consumption,omitempty"`
	// 目前里程
	CurrentMileage *int64 `json:"current_mileage,omitempty"`
	// 車行統編
	TaxID *string `json:"tax_id,omitempty"`
	// 油耗種類
	FuelType *string `json:"fuel_type,omitempty"`
	// ETC計費車種
	BillingType *string `json:"billing_type,omitempty"`
	// 車種
	Style *string `json:"style,omitempty"`
	// 車重(噸)
	Weight *string `json:"weight,omitempty"`
	// 每日成本
	DailyCost *int64 `json:"daily_cost,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}
