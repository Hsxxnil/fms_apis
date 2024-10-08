CREATE TABLE status_configurations
(
    id         uuid      default uuid_generate_v4() not null
        primary key,                                                                   -- 表ID
    status_id  int                                  not null references statuses (id), -- 狀態ID
    limit_time int,                                                                    -- 時間限制
    fleet_id   uuid                                 not null references fleets (id),   -- 車隊ID
    created_at timestamp default now(),
    created_by uuid,
    updated_at timestamp,
    updated_by uuid,
    deleted_at timestamp
);

create index idx_status_configurations_status
    on status_configurations (status_id);

create index idx_status_configurations_fleet_id
    on status_configurations using hash (fleet_id);

create index idx_status_configurations_created_at
    on status_configurations (created_at desc);

create index idx_status_configurations_created_by
    on status_configurations using hash (created_by);

create index idx_status_configurations_updated_at
    on status_configurations (updated_at desc);

create index idx_status_configurations_updated_by
    on status_configurations using hash (updated_by);