package jasslin_data

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
	// 車機IMEI
	Imei string `gorm:"->;column:imei;type:text;not null;" json:"imei"`
	// 車機序號
	SID string `gorm:"->;column:sid;type:text;not null;index:idx_jasslin_data_sid" json:"sid"`
	// vehicles data
	Vehicles vehicles.Table `gorm:"foreignKey:SID;references:SID" json:"vehicles,omitempty"`
	// SIM卡號
	Imsi string `gorm:"->;column:imsi;type:text;not null;" json:"imsi"`
	// 司機代號
	DriverCode string `gorm:"->;column:driver_code;type:text;" json:"driver_code"`
	// 時間
	DateTime time.Time `gorm:"->;column:date_time;type:timestamp;not null;index:idx_jasslin_data_date_time" json:"date_time"`
	// 衛星定位狀態 true:定位 false:未定位
	SatStatus bool `gorm:"->;column:sat_status;type:bool;" json:"sat_status"`
	// 訊號強度
	GsmCsq int64 `gorm:"->;column:gsm_csq;type:int;" json:"gsm_csq"`
	// 衛星數
	SatUsed int64 `gorm:"->;column:sat_used;type:int;" json:"sat_used"`
	// GPS里程數
	Mileage float64 `gorm:"->;column:mileage;type:real;" json:"mileage"`
	// 封包數
	PacketCount int64 `gorm:"->;column:packet_count;type:int;" json:"packet_count"`
	// 緯度
	Lat float64 `gorm:"->;column:lat;type:double precision;" json:"lat"`
	// 經度
	Lon float64 `gorm:"->;column:lon;type:double precision;" json:"lon"`
	// 車頭方向
	Heading int64 `gorm:"->;column:heading;type:int;" json:"heading"`
	// 車速
	Speed float64 `gorm:"->;column:speed;type:real;" json:"speed"`
	// RPM
	Rpm float64 `gorm:"->;column:rpm;type:real;" json:"rpm"`
	// I/O(ACC、煞車、左燈、右燈、IO/1、IO/2、IO/3、IO/4)
	IoStatus int64 `gorm:"->;column:io_status;type:int;" json:"io_status"`
	// GPS速度
	GpsSpeed float64 `gorm:"->;column:gps_speed;type:real;" json:"gps_speed"`
	// 狀態(USB連接狀態、保留、保留、保留、保留、保留、超速、怠速)
	GpsStatus int64 `gorm:"->;column:gps_status;type:int;" json:"gps_status"`
	// CRC checksum
	Crc string `gorm:"->;column:crc;type:text;" json:"crc"`
	// 地址
	Address *string `gorm:"column:address;type:text;" json:"address"`
	// 狀態
	Status *string `gorm:"column:status;type:text;" json:"status"`
	// 狀態累積時間
	StatusTime *string `gorm:"column:status_time;type:text;" json:"status_time"`
	// 回傳時間
	ServerTime time.Time `gorm:"->;column:server_time;type:timestamp;index:idx_jasslin_data_server_time" json:"server_time"`
}

// Base struct is corresponding to fleets table structure file
type Base struct {
	// 表ID
	ID *int `json:"id,omitempty"`
	// 車機IMEI
	Imei *string `json:"imei,omitempty"`
	// 車機序號
	SID *string `json:"sid,omitempty"`
	// vehicles data
	Vehicles vehicles.Base `json:"vehicles,omitempty"`
	// SIM卡號
	Imsi *string `json:"imsi,omitempty"`
	// 司機代號
	DriverCode *string `json:"driver_code,omitempty"`
	// 時間
	DateTime time.Time `json:"date_time,omitempty"`
	// 衛星定位狀態 true:定位 false:未定位
	SatStatus *bool `json:"sat_status,omitempty"`
	// 訊號強度
	GsmCsq *int64 `json:"gsm_csq,omitempty"`
	// 緯度
	Lat *float64 `json:"lat,omitempty"`
	// 經度
	Lon *float64 `json:"lon,omitempty"`
	// 車速
	Speed *float64 `json:"speed,omitempty"`
	// 車頭方向
	Heading *int64 `json:"heading,omitempty"`
	// GPS里程數
	Mileage *float64 `json:"mileage,omitempty"`
	// 衛星數
	SatUsed *int64 `json:"sat_used,omitempty"`
	// 封包數
	PacketCount *int64 `json:"packet_count,omitempty"`
	// RPM
	Rpm *float64 `json:"rpm,omitempty"`
	// I/O(ACC、煞車、左燈、右燈、IO/1、IO/2、IO/3、IO/4)
	IoStatus *int64 `json:"io_status,omitempty"`
	// GPS速度
	GpsSpeed *float64 `json:"gps_speed,omitempty"`
	// 狀態(USB連接狀態、保留、保留、保留、保留、保留、超速、怠速)
	GpsStatus *int64 `json:"gps_status,omitempty"`
	// CRC checksum
	Crc *string `json:"crc,omitempty"`
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
	// jasslin車機序號(後端搜尋用)
	JasslinSID []*string `json:"jasslin_sid,omitempty"`
}

// TableName sets the insert table date for this struct type
func (t *Table) TableName() string {
	return "jasslin_data"
}
