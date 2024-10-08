CREATE TABLE gps_devices
(
    id           uuid      default uuid_generate_v4() not null
        primary key,                                            -- 表ID
    sid          text                                 not null, -- 車機序號
    firm         text,                                          -- 廠商
    model        text                                 not null, -- 型號
    phone_number text,                                          -- 門號
    created_at   timestamp default now(),
    created_by   uuid,
    updated_at   timestamp,
    updated_by   uuid,
    deleted_at   timestamp
);

create index idx_gps_devices_sid
    on gps_devices using gin (sid gin_trgm_ops);

create index idx_gps_devices_firm
    on gps_devices using gin (firm gin_trgm_ops);

create index idx_gps_devices_model
    on gps_devices using gin (model gin_trgm_ops);

create index idx_gps_devices_phone_number
    on gps_devices using gin (phone_number gin_trgm_ops);

create index idx_gps_devices_created_at
    on gps_devices (created_at desc);

create index idx_gps_devices_created_by
    on gps_devices using hash (created_by);

create index idx_gps_devices_updated_at
    on gps_devices (updated_at desc);

create index idx_gps_devices_updated_by
    on gps_devices using hash (updated_by);