package avema_data

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
	SID string `gorm:"->;column:sid;type:text;not null;index:idx_avema_data_sid" json:"sid"`
	// vehicles data
	Vehicles vehicles.Table `gorm:"foreignKey:SID;references:SID" json:"vehicles,omitempty"`
	// 時間
	DateTime time.Time `gorm:"->;column:date_time;type:timestamp;not null;index:idx_avema_data_date_time" json:"date_time"`
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
	// MSGID
	ReportID int64 `gorm:"->;column:report_id;type:int;" json:"report_id"`
	// GPS里程數
	Mileage float64 `gorm:"->;column:mileage;type:real;" json:"mileage"`
	// 輸入PORT二進制表示
	Inputs int64 `gorm:"->;column:inputs;type:bit;" json:"inputs"`
	// 類比輸入電壓
	AnalogInput string `gorm:"->;column:analog_input;type:text;" json:"analog_input"`
	// 主電源電壓
	MainPowerVol string `gorm:"->;column:main_power_vol;type:text;" json:"main_power_vol"`
	// 輸出PORT二進制表示
	Outputs int64 `gorm:"->;column:outputs;type:bit;" json:"outputs"`
	// 4G訊號強度
	GsmCsq int64 `gorm:"->;column:gsm_csq;type:int;" json:"gsm_csq"`
	// 4:4G系統 , 3:3G系統
	GsmMode int64 `gorm:"->;column:gsm_mode;type:int;" json:"gsm_mode"`
	// 溫度_司機ID
	TemperatureDriverID string `gorm:"->;column:temperature_driverid;type:text;" json:"temperature_driver_id"`
	// 地址
	Address string `gorm:"column:address;type:text;" json:"address"`
	// 狀態
	Status string `gorm:"column:status;type:text;" json:"status"`
	// 狀態累積時間
	StatusTime string `gorm:"column:status_time;type:text;" json:"status_time"`
	// 回傳時間
	ServerTime time.Time `gorm:"->;column:server_time;type:timestamp;index:idx_avema_data_server_time" json:"server_time"`
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
	// MSGID
	ReportID *int64 `json:"report_id,omitempty"`
	// GPS里程數
	Mileage *float64 `json:"mileage,omitempty"`
	// 輸入PORT二進制表示
	Inputs *int64 `json:"inputs,omitempty"`
	// 類比輸入電壓
	AnalogInput *string `json:"analog_input,omitempty"`
	// 主電源電壓
	MainPowerVol *string `json:"main_power_vol,omitempty"`
	// 輸出PORT二進制表示
	Outputs *int64 `json:"outputs,omitempty"`
	// 4G訊號強度
	GsmCsq *int64 `json:"gsm_csq,omitempty"`
	// 4:4G系統 , 3:3G系統
	GsmMode *int64 `json:"gsm_mode,omitempty"`
	// 溫度_司機ID
	TemperatureDriverID *string `json:"temperature_driver_id,omitempty"`
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
	// avema車機序號(後端搜尋用)
	AvemaSID []*string `json:"avema_sid,omitempty"`
}

// TableName sets the insert table date for this struct type
func (t *Table) TableName() string {
	return "avema_data"
}
