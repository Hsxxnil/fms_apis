package gps

import (
	"fms/internal/interactor/models/page"
	"time"
)

// Field is structure file for search
type Field struct {
	// 表ID
	ID int `json:"id,omitempty" swaggerignore:"true"`
	// 車機序號
	SID *string `json:"sid,omitempty" form:"sid"`
	// 車輛牌照
	LicensePlate *string `json:"license_plate,omitempty" form:"license_plate"`
	//車隊ID
	FleetID string `json:"fleet_id,omitempty" form:"fleet_id"`
	// 搜尋欄位
	Filter `json:"filter"`
	// avema車機序號(後端搜尋用)
	AvemaSID []*string `json:"avema_sid,omitempty" swaggerignore:"true"`
	// poison車機序號(後端搜尋用)
	PoisonSID []*string `json:"poison_sid,omitempty" swaggerignore:"true"`
	// jasslin車機序號(後端搜尋用)
	JasslinSID []*string `json:"jasslin_sid,omitempty" swaggerignore:"true"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	// 搜尋結構檔
	Field
	// 分頁搜尋結構檔
	page.Pagination
}

// Filter struct is used to store the search field
type Filter struct {
	// 搜尋開始時間
	FilterStartTime *time.Time `json:"start_time,omitempty"`
	// 搜尋結束時間
	FilterEndTime *time.Time `json:"end_time,omitempty"`
}

// RealTimeList is multiple return real-time structure files
type RealTimeList struct {
	// 多筆
	Gps []*RealTimeSingle `json:"gps"`
	// 狀態累積時間
	StatusTime *StatusTime `json:"status_time,omitempty"`
	// 分頁返回結構檔
	page.Total
}

// RealTimeSingle return real-time structure file
type RealTimeSingle struct {
	// 表ID
	ID int `json:"id"`
	// 車機序號
	SID string `json:"sid"`
	// 車輛司機
	Driver string `json:"driver,omitempty"`
	// 車輛名稱
	VehicleName string `json:"vehicle_name,omitempty"`
	// 車輛牌照
	LicensePlate string `json:"license_plate,omitempty"`
	// 時間
	DateTime time.Time `json:"date_time"`
	// 緯度
	Lat float64 `json:"lat"`
	// 經度
	Lon float64 `json:"lon"`
	// 經度
	Lng float64 `json:"lng"`
	// 車速
	Speed float64 `json:"speed"`
	// 車頭方向
	Heading int64 `json:"heading"`
	// 衛星數
	SatUsed int64 `json:"sat_used"`
	// GPS里程數
	Mileage float64 `json:"mileage"`
	// 地址
	Address string `json:"address"`
	// 狀態
	Status string `json:"status"`
	// 狀態累積時間
	StatusTime string `json:"status_time"`
	// 回傳時間
	ServerTime time.Time `json:"server_time"`
}

// StatusTime is the structure file of the status time.
type StatusTime struct {
	// 怠停時間
	IdleTime string `json:"idle_time,omitempty"`
	// 久停時間
	LongIdleTime string `json:"long_idle_time,omitempty"`
	// 行駛時間
	DrivingTime string `json:"driving_time,omitempty"`
	// 熄火時間
	EngineOffTime string `json:"engine_off_time,omitempty"`
	// 失聯時間
	LostConnTime string `json:"lost_conn_time,omitempty"`
}

// Update struct is used to update achieves
type Update struct {
	// 車機序號
	SID string `json:"sid,omitempty" swaggerignore:"true"`
	// 回傳時間
	ServerTime time.Time `json:"server_time" swaggerignore:"true"`
	// 地址
	Address *string `json:"address,omitempty"`
	// 狀態
	Status *string `json:"status,omitempty"`
	// 狀態累積時間
	StatusTime *string `json:"status_time,omitempty"`
}

// DetermineStatus struct is used to determine status
type DetermineStatus struct {
	// 表ID
	ID int `json:"id"`
	// 車機序號
	SID string `json:"sid"`
	// 車輛司機
	Driver string `json:"driver,omitempty"`
	// 車輛名稱
	VehicleName string `json:"vehicle_name,omitempty"`
	// 車輛牌照
	LicensePlate string `json:"license_plate,omitempty"`
	// 時間
	DateTime time.Time `json:"date_time"`
	// 緯度
	Lat float64 `json:"lat"`
	// 經度
	Lon float64 `json:"lon,omitempty"`
	// 經度
	Lng float64 `json:"lng,omitempty"`
	// 車速
	Speed float64 `json:"speed"`
	// 車頭方向
	Heading int64 `json:"heading"`
	// 衛星數
	SatUsed int64 `json:"sat_used"`
	// GPS里程數
	Mileage float64 `json:"mileage"`
	// 地址
	Address string `json:"address"`
	// 狀態
	Status string `json:"status"`
	// 狀態累積時間
	StatusTime string `json:"status_time"`
	// IO1 (ch68機型判斷熄火)
	Io1 int64 `json:"io1,omitempty"`
	// MSGID (avema傳入資料類型)
	ReportID int64 `json:"report_id,omitempty"`
	// 輸入PORT二進制表示 (at35機型判斷熄火)
	Inputs int64 `json:"inputs,omitempty"`
	// 回傳時間
	ServerTime time.Time `json:"server_time"`
}
