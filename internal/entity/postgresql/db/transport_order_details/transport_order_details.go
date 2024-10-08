package transport_order_details

import (
	"fms/internal/entity/postgresql/db/trailers"
	"fms/internal/interactor/models/special"
)

// Table struct is transport_order_details database table struct
type Table struct {
	// 品項
	ProductName string `gorm:"column:product_name;type:int;not null;" json:"product_name"`
	// 單價
	UnitPrice float64 `gorm:"column:unit_price;type:numeric;" json:"unit_price"`
	// 數量
	Quantity int64 `gorm:"column:quantity;type:int;" json:"quantity"`
	// 噸數
	Tonnage int64 `gorm:"column:tonnage;type:int;" json:"tonnage"`
	// 托運訂單ID
	TransportOrderID string `gorm:"column:transport_order_id;type:uuid;not null;" json:"transport_order_id"`
	// 板台ID
	TrailerID *string `gorm:"column:trailer_id;type:uuid;" json:"trailer_id"`
	// trailers data
	Trailers trailers.Table `gorm:"foreignKey:ID;references:TrailerID" json:"trailers,omitempty"`
	special.Table
}

// Base struct is corresponding to transport_order_details table structure file
type Base struct {
	// 品項
	ProductName *string `json:"product_name,omitempty"`
	// 單價
	UnitPrice *float64 `json:"unit_price,omitempty"`
	// 數量
	Quantity *int64 `json:"quantity,omitempty"`
	// 噸數
	Tonnage *int64 `json:"tonnage,omitempty"`
	// 托運訂單ID
	TransportOrderID *string `json:"transport_order_id,omitempty"`
	// 板台ID
	TrailerID *string `json:"trailer_id,omitempty"`
	// trailers data
	Trailers *trailers.Base `json:"trailers,omitempty"`
	special.Base
}

// TableName sets the insert table transport_order_details for this struct type
func (t *Table) TableName() string {
	return "transport_order_details"
}
