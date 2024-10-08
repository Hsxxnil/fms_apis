CREATE TABLE jasslin_data
(
    id           serial primary key      not null, -- 表ID
    imei         text                    not null, -- 車機IMEI
    sid          text                    not null, -- 車號
    imsi         text                    not null, -- SIM卡號
    driver_id    text,                             -- 司機ID
    date_time    timestamp,                        -- 時間
    sat_status   bool,                             -- 衛星定位狀態
    gsm_csq      int,                              -- 訊號強度
    sat_used     int,                              -- 衛星數
    mileage      real,                             -- 里程(公尺)
    packet_count int,                              -- 封包數
    lat          double precision,                 -- 緯度
    lon          double precision,                 -- 經度
    heading      int,                              -- 車頭方向
    speed        real,                             -- 時速
    rpm          real,                             -- 轉速
    io           int,                              -- I/O(ACC、煞車、左燈、右燈、IO/1、IO/2、IO/3、IO/4)
    gps_speed    real,                             -- GPS車速
    status       int,                              -- 狀態(USB連接狀態、保留、保留、保留、保留、保留、超速、怠速)
    crc          text,                             -- CRC checksum
    server_time  timestamp default now() not null  -- 回傳時間
);

create index idx_jasslin_data_date_time
    on jasslin_data (date_time desc);

create index idx_jasslin_data_server_time
    on jasslin_data (server_time desc);

create index idx_jasslin_data_sid
    on jasslin_data (sid);