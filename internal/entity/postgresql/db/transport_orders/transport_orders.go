package transport_orders

import (
	"fms/internal/entity/postgresql/db/clients"
	"fms/internal/entity/postgresql/db/transport_order_details"
	"fms/internal/entity/postgresql/db/users"
	"fms/internal/interactor/models/special"
	"time"
)

// Table struct is transport_orders database table struct
type Table struct {
	// 托運訂單名稱
	Name string `gorm:"column:name;type:int;not null;" json:"name"`
	// 托運訂單單號
	Code string `gorm:"column:code;type:text;not null;" json:"code"`
	// 託運人ID
	ClientID string `gorm:"column:client_id;type:uuid;not null;" json:"client_id"`
	// clients data
	Client clients.Table `gorm:"foreignKey:ID;references:ClientID" json:"client,omitempty"`
	// 起運點
	Origin string `gorm:"column:origin;type:text;not null;" json:"origin"`
	// 卸貨點
	Destination string `gorm:"column:destination;type:text;not null;" json:"destination"`
	// 指定抵達時間
	Deadline *time.Time `gorm:"column:deadline;type:timestamp;" json:"deadline"`
	// 派工任務ID
	TransportTaskID *string `gorm:"column:transport_task_id;type:uuid;" json:"transport_task_id"`
	// 順序
	Sequence *int64 `gorm:"column:sequence;type:int;" json:"sequence"`
	// 車隊ID
	FleetID string `gorm:"column:fleet_id;type:uuid;not null;" json:"fleet_id"`
	// transport_order_details data
	TransportOrderDetails []*transport_order_details.Table `gorm:"foreignKey:TransportOrderID;references:ID" json:"shipping_list,omitempty"`
	// create_users data
	CreatedByUsers users.Table `gorm:"foreignKey:ID;references:CreatedBy" json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Table `gorm:"foreignKey:ID;references:UpdatedBy" json:"updated_by_users,omitempty"`
	special.Table
}

// Base struct is corresponding to transport_orders table structure file
type Base struct {
	// 車隊ID
	FleetID *string `json:"fleet_id,omitempty"`
	// 托運訂單名稱
	Name *string `json:"name,omitempty"`
	// 托運訂單單號
	Code *string `json:"code,omitempty"`
	// 託運人ID
	ClientID *string `json:"client_id,omitempty"`
	// clients data
	Client *clients.Base `json:"client,omitempty"`
	// 起運點
	Origin *string `json:"origin,omitempty"`
	// 卸貨點
	Destination *string `json:"destination,omitempty"`
	// 指定抵達時間
	Deadline *time.Time `json:"deadline,omitempty"`
	// 派工任務ID
	TransportTaskID *string `json:"transport_task_id,omitempty"`
	// 順序
	Sequence *int64 `json:"sequence,omitempty"`
	// IDs (後端搜尋用)
	IDs []*string `json:"ids,omitempty"`
	// transport_order_details data
	TransportOrderDetails []*transport_order_details.Base `json:"shipping_list,omitempty"`
	// create_users data
	CreatedByUsers users.Base `json:"created_by_users,omitempty"`
	// update_users data
	UpdatedByUsers users.Base `json:"updated_by_users,omitempty"`
	special.Base
}

// TableName sets the insert table transport_orders for this struct type
func (t *Table) TableName() string {
	return "transport_orders"
}
