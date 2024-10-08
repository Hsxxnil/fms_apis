package transport_tasks

import (
	"fms/internal/interactor/models/page"
	"fms/internal/interactor/models/section"
	"fms/internal/interactor/models/transport_orders"
	"time"
)

// Create struct is used to create achieves
type Create struct {
	// 車隊ID
	FleetID string `json:"fleet_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 派工任務名稱
	Title string `json:"title,omitempty" binding:"required" validate:"required"`
	// 實際開始時間
	ActualStartTime *time.Time `json:"actual_start_time,omitempty"`
	// 預計開始時間
	ScheduledStartTime *time.Time `json:"scheduled_start_time,omitempty"`
	// 實際結束時間
	ActualEndTime *time.Time `json:"actual_end_time,omitempty"`
	// 預計結束時間
	ScheduledEndTime *time.Time `json:"scheduled_end_time,omitempty"`
	// 車輛ID
	VehicleID string `json:"vehicle_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 司機ID
	DriverID string `json:"driver_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// transport_order_ids
	TransportOrderIDs []string `json:"form,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 表ID
	ID string `json:"id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 派工任務名稱
	Title *string `json:"title,omitempty" form:"title"`
	// 實際開始時間
	ActualStartTime *time.Time `json:"actual_start_time,omitempty" form:"actual_start_time"`
	// 預計開始時間
	ScheduledStartTime *time.Time `json:"scheduled_start_time,omitempty" form:"scheduled_start_time"`
	// 實際結束時間
	ActualEndTime *time.Time `json:"actual_end_time,omitempty" form:"actual_end_time"`
	// 預計結束時間
	ScheduledEndTime *time.Time `json:"scheduled_end_time,omitempty" form:"scheduled_end_time"`
	// 車輛ID
	VehicleID *string `json:"vehicle_id,omitempty" form:"vehicle_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 司機ID
	DriverID *string `json:"driver_id,omitempty" form:"driver_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
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
	TransportTasks []*struct {
		// 表ID
		ID string `json:"id,omitempty"`
		// 派工任務名稱
		Title string `json:"title"`
		// 實際開始時間
		ActualStartTime *time.Time `json:"actual_start_time,omitempty"`
		// 預計開始時間
		ScheduledStartTime *time.Time `json:"scheduled_start_time,omitempty"`
		// 實際結束時間
		ActualEndTime *time.Time `json:"actual_end_time,omitempty"`
		// 預計結束時間
		ScheduledEndTime *time.Time `json:"scheduled_end_time,omitempty"`
		// 車輛ID
		VehicleID string `json:"vehicle_id,omitempty"`
		// 車輛名稱
		VehicleName string `json:"vehicle_name,omitempty"`
		// 司機ID
		DriverID string `json:"driver_id,omitempty"`
		// 司機名稱
		DriverName string `json:"driver_name,omitempty"`
		// transport_orders data
		TransportOrders []transport_orders.TaskSingle `json:"form,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"transport_tasks"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 表ID
	ID string `json:"id,omitempty"`
	// 派工任務名稱
	Title string `json:"title,omitempty"`
	// 實際開始時間
	ActualStartTime *time.Time `json:"actual_start_time,omitempty"`
	// 預計開始時間
	ScheduledStartTime *time.Time `json:"scheduled_start_time,omitempty"`
	// 實際結束時間
	ActualEndTime *time.Time `json:"actual_end_time,omitempty"`
	// 預計結束時間
	ScheduledEndTime *time.Time `json:"scheduled_end_time,omitempty"`
	// 車輛ID
	VehicleID string `json:"vehicle_id,omitempty"`
	// 車輛名稱
	VehicleName string `json:"vehicle_name,omitempty"`
	// 司機ID
	DriverID string `json:"driver_id,omitempty"`
	// 司機名稱
	DriverName string `json:"driver_name,omitempty"`
	// transport_orders data
	TransportOrders []transport_orders.TaskSingle `json:"form,omitempty"`
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
	VehicleID *string `json:"vehicle_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 司機ID
	DriverID *string `json:"driver_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// transport_order_ids
	TransportOrderIDs []string `json:"form,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}
