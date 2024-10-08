package poison_data

import (
	"fms/internal/entity/postgresql/db/vehicles"
	model "fms/internal/interactor/models/gps"
	"fms/internal/interactor/models/page"
	"time"
)

// Table struct is fleets database table struct
type Table struct {
	// 表ID
	ID int `gorm:"->;column:id;type:serial;not null;primaryKey;" json:"id"`
	// 車機序號
	SID string `gorm:"->;column:sid;type:text;not null;index:idx_poison_data_sid" json:"sid"`
	// vehicles data
	Vehicles vehicles.Table `gorm:"foreignKey:SID;references:SID" json:"vehicles,omitempty"`
	// 時間
	DateTime time.Time `gorm:"->;column:date_time;type:timestamp;not null;index:idx_poison_data_date_time" json:"date_time"`
	// 緯度
	Lat float64 `gorm:"->;column:lat;type:double precision;" json:"lat"`
	// 經度
	Lon float64 `gorm:"->;column:lon;type:double precision;" json:"lon"`
	// 車速
	Speed float64 `gorm:"->;column:speed;type:real;" json:"speed"`
	// 車頭方向
	Heading int64 `gorm:"->;column:heading;type:int;" json:"heading"`
	// 衛星數
	SatUsed int64 `gorm:"->;column:sat_used;type:int;" json:"sat_used"`
	// 資料種類
	DataType string `gorm:"->;column:data_type;type:text;" json:"data_type"`
	// IO1
	Io1 int64 `gorm:"->;column:io1;type:int;" json:"io1"`
	// IO2
	Io2 int64 `gorm:"->;column:io2;type:int;" json:"io2"`
	// IO3
	Io3 int64 `gorm:"->;column:io3;type:int;" json:"io3"`
	// 保留欄位
	Reserved string `gorm:"->;column:reserved;type:text;" json:"reserved"`
	// 條碼
	Barcode string `gorm:"->;column:barcode;type:text;" json:"barcode"`
	// 溫度1
	Temp1 string `gorm:"->;column:temp1;type:text;" json:"temp1"`
	// 溫度2
	Temp2 string `gorm:"->;column:temp2;type:text;" json:"temp2"`
	// 地址
	Address string `gorm:"column:address;type:text;" json:"address"`
	// 狀態
	Status string `gorm:"column:status;type:text;" json:"status"`
	// 狀態累積時間
	StatusTime string `gorm:"column:status_time;type:text;" json:"status_time"`
	// 回傳時間
	ServerTime time.Time `gorm:"->;column:server_time;type:timestamp;index:idx_poison_data_server_time" json:"server_time"`
}

// Base struct is corresponding to fleets table structure file
type Base struct {
	// 表ID
	ID *int `json:"id,omitempty"`
	// 車機序號
	SID *string `json:"sid,omitempty"`
	// vehicles data
	Vehicles vehicles.Base `json:"vehicles,omitempty"`
	// 時間
	DateTime time.Time `json:"date_time,omitempty"`
	// 緯度
	Lat *float64 `json:"lat,omitempty"`
	// 經度
	Lon *float64 `json:"lon,omitempty"`
	// 車速
	Speed *float64 `json:"speed,omitempty"`
	// 車頭方向
	Heading *int64 `json:"heading,omitempty"`
	// 衛星數
	SatUsed *int64 `json:"sat_used,omitempty"`
	// 資料種類
	DataType *string `json:"data_type,omitempty"`
	// IO1
	Io1 *int64 `json:"io1,omitempty"`
	// IO2
	Io2 *int64 `json:"io2,omitempty"`
	// IO3
	Io3 *int64 `json:"io3,omitempty"`
	// 保留欄位
	Reserved *string `json:"reserved,omitempty"`
	// 條碼
	Barcode *string `json:"barcode,omitempty"`
	// 溫度1
	Temp1 *string `json:"temp1,omitempty"`
	// 溫度2
	Temp2 *string `json:"temp2,omitempty"`
	// 地址
	Address *string `json:"address,omitempty"`
	// 狀態
	Status *string `json:"status,omitempty"`
	// 狀態累積時間
	StatusTime *string `json:"status_time,omitempty"`
	// 回傳時間
	ServerTime *time.Time `json:"server_time,omitempty"`
	// 引入page
	page.Pagination
	// 搜尋欄位
	model.Filter `json:"filter"`
	// poison車機序號(後端搜尋用)
	PoisonSID []*string `json:"poison_sid,omitempty"`
}

// TableName sets the insert table date for this struct type
func (t *Table) TableName() string {
	return "poison_data"
}
