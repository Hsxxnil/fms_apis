CREATE TABLE ch68_d1_data
(
    id          serial primary key      not null, -- 表ID
    sid         text                    not null, -- 車機序號
    date_time   timestamp,                        -- 時間
    lat         double precision,                 -- 緯度
    lon         double precision,                 -- 經度
    speed       real,                             -- 車速
    heading     int,                              -- 車頭方向
    sat_used    int,                              -- 衛星數
    data_type   int,                              -- 資料種類
    io1         int,                              -- io1
    io2         int,                              -- io2
    io3         int,                              -- io3
    reserved    text,                             -- 保留欄位
    barcode     text,                             -- 條碼
    server_time timestamp default now() not null  -- 回傳時間
);