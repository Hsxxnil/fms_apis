package transport_tasks

import (
	"fms/internal/entity/postgresql/db/drivers"
	"fms/internal/entity/postgresql/db/transport_orders"
	"fms/internal/entity/postgresql/db/users"
	"fms/internal/entity/postgresql/db/vehicles"
	"fms/internal/interactor/models/special"
	"time"
)

// Table struct is transport_tasks database table struct
type Table struct {
	// 派工任務名稱
	Title string `gorm:"column:title;type:int;not null;" json:"title"`
	// 實際開始時間
	ActualStartTime *time.Time `gorm:"column:actual_start_time;type:timestamp;" json:"actual_start_time"`
	// 預計開始時間
	ScheduledStartTime *time.Time `gorm:"column:scheduled_start_time;type:timestamp;" json:"scheduled_start_time"`
	// 實際結束時間
	ActualEndTime *time.Time `gorm:"column:actual_end_time;type:timestamp;" json:"actual_end_time"`
	// 預計結束時間
	ScheduledEndTime *time.Time `gorm:"column:scheduled_end_time;type:timestamp;" json:"scheduled_end_time"`
	// 車輛ID
	VehicleID string `gorm:"column:vehicle_id;type:uuid;not null;" json:"vehicle_id"`
	// vehicles data
	Vehicles vehicles.Table `gorm:"foreignKey:VehicleID;references:ID" json:"vehicles,omitempty"`
	// 司機ID
	DriverID string `gorm:"column:driver_id;type:uuid;not null;" json:"driver_id"`
	// drivers data
	Drivers drivers.Table `gorm:"foreignKey:DriverID;references:ID" json:"drivers,omitempty"`
	// 車隊ID
	FleetID string `gorm:"column:fleet_id;type:uuid;not null;" json:"fleet_id"`
	// transport_orders data
	TransportOrders []transport_orders.Table `gorm:"foreignKey:TransportTaskID;references:ID" json:"transport_orders,omitempty"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:ID;references:CreatedBy" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:ID;references:UpdatedBy" json:"updated_by_users,omitempty"`
	special.Table
}

// Base struct is corresponding to transport_tasks table structure file
type Base struct {
	// 車隊ID
	FleetID *string `json:"fleet_id,omitempty"`
	// 派工任務名稱
	Title *string `json:"title,omitempty"`
	// 實際開始時間
	ActualStartTime *time.Time `json:"actual_start_time,omitempty"`
	// 預計開始時間
	ScheduledStartTime *time.Time `json:"scheduled_start_time,omitempty"`
	// 實際結束時間
	ActualEndTime *time.Time `json:"actual_end_time,omitempty"`
	// 預計結束時間
	ScheduledEndTime *time.Time `json:"scheduled_end_time,omitempty"`
	// 車輛ID
	VehicleID *string `json:"vehicle_id,omitempty"`
	// vehicles data
	Vehicles vehicles.Base `json:"vehicles,omitempty"`
	// 司機ID
	DriverID *string `json:"driver_id,omitempty"`
	// drivers data
	Drivers drivers.Base `json:"drivers,omitempty"`
	// transport_orders data
	TransportOrders []transport_orders.Base `json:"transport_orders,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.Base
}

// TableName sets the insert table transport_tasks for this struct type
func (t *Table) TableName() string {
	return "transport_tasks"
}
