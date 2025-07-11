CREATE TABLE z1_data
(
    id                   serial primary key      not null, -- 表ID
    sid                  text                    not null, -- 車機序號
    date_time            timestamp,                        -- 時間
    lat                  double precision,                 -- 緯度
    lon                  double precision,                 -- 經度
    speed                real,                             -- 車速
    heading              int,                              -- 車頭方向
    sat_used             int,                              -- 衛星數
    report_id            int,                              -- MSGID
    mileage              real,                             -- GPS里程數
    inputs               int,                              -- 輸入PORT二進制表示
    analog_input         text,                             -- 類比輸入電壓
    main_power_vol       text,                             -- 主電源電壓
    outputs              int,                              -- 輸出PORT二進制表示
    gsm_csq              int,                              -- 4G訊號強度
    gsm_mode             int,                              -- 4:4G系統 , 3:3G系統
    temperature_driverid text,                             -- 溫度_司機ID
    server_time          timestamp default now() not null  -- 回傳時間
);