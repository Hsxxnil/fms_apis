CREATE TABLE transport_tasks
(
    id                   uuid      default uuid_generate_v4() not null
        primary key,                                                                             -- 表ID
    name                 text                                 not null,                          -- 派工任務名稱
    actual_start_time    timestamp,                                                              -- 實際開始時間
    scheduled_start_time timestamp,                                                              -- 預計開始時間
    actual_end_time      timestamp,                                                              -- 實際結束時間
    scheduled_end_time   timestamp,                                                              -- 預計結束時間
    driver_id            uuid                                 not null references drivers (id),  -- 司機ID
    vehicle_id           uuid                                 not null references vehicles (id), -- 車輛ID
    trailer_id           uuid references trailers (id),                                          -- 板台ID
    created_at           timestamp default now(),
    created_by           uuid,
    updated_at           timestamp,
    updated_by           uuid,
    deleted_at           timestamp
);

create index idx_transport_tasks_name
    on transport_tasks (name);

create index idx_transport_tasks_actual_start_time
    on transport_tasks (actual_start_time);

create index idx_transport_tasks_scheduled_start_time
    on transport_tasks (scheduled_start_time);

create index idx_transport_tasks_actual_end_time
    on transport_tasks (actual_end_time);

create index idx_transport_tasks_scheduled_end_time
    on transport_tasks (scheduled_end_time);

create index idx_transport_tasks_driver_id
    on transport_tasks using hash (driver_id);

create index idx_transport_tasks_vehicle_id
    on transport_tasks using hash (vehicle_id);

create index idx_transport_tasks_trailer_id
    on transport_tasks using hash (trailer_id);

create index idx_transport_tasks_created_at
    on transport_tasks (created_at desc);

create index idx_transport_tasks_created_by
    on transport_tasks using hash (created_by);

create index idx_transport_tasks_updated_at
    on transport_tasks (updated_at desc);

create index idx_transport_tasks_updated_by
    on transport_tasks using hash (updated_by);