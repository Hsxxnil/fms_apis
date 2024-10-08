package vehicles

import (
	"fms/internal/entity/postgresql/db/fleets"
	"fms/internal/entity/postgresql/db/users"
	"fms/internal/interactor/models/special"
)

// Table struct is vehicles database table struct
type Table struct {
	// 車隊ID
	FleetID string `gorm:"column:fleet_id;type:uuid;not null;" json:"fleet_id"`
	// fleets data
	Fleets fleets.Table `gorm:"foreignKey:ID;references:FleetID" json:"fleets,omitempty"`
	// 車輛名稱
	Name string `gorm:"column:name;type:text;not null;" json:"name"`
	// 車輛司機
	Driver string `gorm:"column:driver;type:text;" json:"driver"`
	// 車輛牌照
	LicensePlate string `gorm:"column:license_plate;type:text;not null;" json:"license_plate"`
	// 車機序號
	SID string `gorm:"column:sid;type:text;not null;" json:"sid"`
	// 油耗量(公里/公升)
	FuelConsumption float64 `gorm:"column:fuel_consumption;type:numeric;" json:"fuel_consumption"`
	// 目前里程
	CurrentMileage int64 `gorm:"column:current_mileage;type:int;" json:"current_mileage"`
	// 車行統編
	TaxID string `gorm:"column:tax_id;type:text;" json:"tax_id"`
	// 油耗種類
	FuelType string `gorm:"column:fuel_type;type:text;" json:"fuel_type"`
	// ETC計費車種
	BillingType string `gorm:"column:billing_type;type:text;" json:"billing_type"`
	// 車種
	Style string `gorm:"column:style;type:text;" json:"style"`
	// 車重(噸)
	Weight string `gorm:"column:weight;type:text;" json:"weight"`
	// 每日成本
	DailyCost int64 `gorm:"column:daily_cost;type:int;" json:"daily_cost"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:ID;references:CreatedBy" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:ID;references:UpdatedBy" json:"updated_by_users,omitempty"`
	special.Table
}

// Base struct is corresponding to vehicles table structure file
type Base struct {
	// 車隊ID
	FleetID *string `json:"fleet_id,omitempty"`
	// fleets data
	Fleets fleets.Base `json:"fleets,omitempty"`
	// 車輛名稱
	Name *string `json:"name,omitempty"`
	// 車輛司機
	Driver *string `json:"driver,omitempty"`
	// 車輛牌照
	LicensePlate *string `json:"license_plate,omitempty"`
	// 車機序號
	SID *string `json:"sid,omitempty,omitempty"`
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
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.Base
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "vehicles"
}
