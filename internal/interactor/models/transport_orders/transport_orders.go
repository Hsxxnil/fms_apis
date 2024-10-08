package transport_orders

import (
	"fms/internal/interactor/models/page"
	"fms/internal/interactor/models/section"
	"fms/internal/interactor/models/transport_order_details"
	"time"
)

// Create struct is used to create achieves
type Create struct {
	// 車隊ID
	FleetID string `json:"fleet_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 托運訂單名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 托運訂單單號
	Code string `json:"code,omitempty"`
	// 起運點
	Origin string `json:"origin,omitempty" binding:"required" validate:"required"`
	// 卸貨點
	Destination string `json:"destination,omitempty" binding:"required" validate:"required"`
	// 託運人ID
	ClientID string `json:"client_id,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	// 指定抵達時間
	Deadline *time.Time `json:"deadline,omitempty"`
	// 派工任務ID
	TransportTaskID *string `json:"transport_task_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// transporter_order_details
	TransporterOrderDetails []*transport_order_details.Create `json:"shipping_list,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 表ID
	ID string `json:"id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 托運訂單名稱
	Name *string `json:"name,omitempty" form:"name"`
	// 托運訂單單號
	Code *string `json:"code,omitempty" form:"code"`
	// 起運點
	Origin *string `json:"origin,omitempty" form:"origin"`
	// 卸貨點
	Destination *string `json:"destination,omitempty" form:"destination"`
	// 託運人ID
	ClientID *string `json:"client_id,omitempty" form:"client_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 指定抵達時間
	Deadline *time.Time `json:"deadline,omitempty" from:"deadline"`
	// 派工任務ID
	TransportTaskID *string `json:"transport_task_id,omitempty" form:"transport_task_id" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 順序
	Sequence *int64 `json:"sequence,omitempty" form:"sequence"`
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
	TransportOrders []*struct {
		// 表ID
		ID string `json:"id,omitempty"`
		// 托運訂單名稱
		Name string `json:"name"`
		// 托運訂單單號
		Code string `json:"code,omitempty"`
		// 起運點
		Origin string `json:"origin,omitempty"`
		// 卸貨點
		Destination string `json:"destination,omitempty"`
		// 託運人ID
		ClientID string `json:"client_id,omitempty"`
		// 託運人名稱
		ClientName string `json:"client_name,omitempty"`
		// 指定抵達時間
		Deadline *time.Time `json:"deadline,omitempty"`
		// 派工任務ID
		TransportTaskID string `json:"transport_task_id,omitempty"`
		// 派工任務名稱
		TransportTaskName string `json:"transport_task_title,omitempty"`
		// 順序
		Sequence int64 `json:"sequence,omitempty"`
		// transporter_order_details data
		TransporterOrderDetails []*transport_order_details.Single `json:"shipping_list,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"transport_orders"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 表ID
	ID string `json:"id,omitempty"`
	// 托運訂單名稱
	Name string `json:"name,omitempty"`
	// 托運訂單單號
	Code string `json:"code,omitempty"`
	// 起運點
	Origin string `json:"origin,omitempty"`
	// 卸貨點
	Destination string `json:"destination,omitempty"`
	// 託運人ID
	ClientID string `json:"client_id,omitempty"`
	// 託運人名稱
	ClientName string `json:"client_name,omitempty"`
	// 指定抵達時間
	Deadline *time.Time `json:"deadline,omitempty"`
	// 派工任務ID
	TransportTaskID string `json:"transport_task_id,omitempty"`
	// 派工任務名稱
	TransportTaskName string `json:"transport_task_title,omitempty"`
	// 順序
	Sequence int64 `json:"sequence,omitempty"`
	// transporter_order_details data
	TransporterOrderDetails []*transport_order_details.Single `json:"shipping_list,omitempty"`
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
	// 托運訂單名稱
	Name *string `json:"name,omitempty"`
	// 托運訂單單號
	Code *string `json:"code,omitempty"`
	// 起運點
	Origin *string `json:"origin,omitempty"`
	// 卸貨點
	Destination *string `json:"destination,omitempty"`
	// 託運人ID
	ClientID *string `json:"client_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 指定抵達時間
	Deadline *time.Time `json:"deadline,omitempty"`
	// 派工任務ID
	TransportTaskID *string `json:"transport_task_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 順序
	Sequence *int64 `json:"sequence,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0" swaggerignore:"true"`
	// IDs (後端搜尋用)
	IDs []*string `json:"ids,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// transporter_order_details
	TransporterOrderDetails []*transport_order_details.Update `json:"shipping_list,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// TaskSingle return structure file for transport_tasks
type TaskSingle struct {
	// 表ID
	ID string `json:"id,omitempty"`
	// 車隊ID
	FleetID *string `json:"fleet_id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	// 托運訂單名稱
	Name string `json:"name,omitempty"`
	// 起運點
	Origin string `json:"origin,omitempty"`
	// 卸貨點
	Destination string `json:"destination,omitempty"`
	// 託運人ID
	ClientID string `json:"client_id,omitempty"`
	// 託運人名稱
	ClientName string `json:"client_name,omitempty"`
	// 指定抵達時間
	Deadline *time.Time `json:"deadline,omitempty"`
	// transporter_order_details data
	TransporterOrderDetails []*transport_order_details.Single `json:"shipping_list,omitempty"`
}
