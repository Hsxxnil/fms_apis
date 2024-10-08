create table vehicles
(
    id            uuid      default uuid_generate_v4() not null
        primary key,                                                                    -- 表ID
    fleet_id      uuid                                 not null references fleets (id), -- 車隊ID
    name          text      default '':: text          not null,                        -- 車輛名稱
    driver        text      default '':: text          not null,                        -- 司機名稱
    license_plate text      default '':: text          not null,                        -- 車牌號碼
    sid           text                                 not null,                        -- 車機序號
    created_at    timestamp default now()              not null,
    created_by    uuid                                 not null,
    updated_at    timestamp,
    updated_by    uuid,
    deleted_at    timestamp
);

create index idx_vehicles_id
    on vehicles using hash (id);

create index idx_vehicles_fleet_id
    on vehicles using hash (fleet_id);

create index idx_vehicles_name
    on vehicles using gin (name gin_trgm_ops);

create index idx_vehicles_driver
    on vehicles using gin (driver gin_trgm_ops);

create index idx_vehicles_license_plate
    on vehicles using gin (license_plate gin_trgm_ops);

create index idx_vehicles_sid
    on vehicles using gin (sid gin_trgm_ops);

create index idx_vehicles_created_at
    on vehicles (created_at desc);

create index idx_vehicles_created_by
    on vehicles using hash (created_by);

create index idx_vehicles_updated_at
    on vehicles (updated_at desc);

create index idx_vehicles_updated_by
    on vehicles using hash (updated_by);