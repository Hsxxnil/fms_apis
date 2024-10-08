create table fleets
(
    id         uuid      default uuid_generate_v4() not null
        primary key,                                          -- 表ID
    fleet_code text      default '':: text          not null, -- 車隊代號
    name       text      default '':: text          not null, -- 車隊名稱
    created_at timestamp default now(),
    created_by uuid,
    updated_at timestamp,
    updated_by uuid,
    deleted_at timestamp
);

create index idx_fleets_id
    on fleets using hash (id);

create index idx_fleets_fleet_code
    on fleets using gin (fleet_code gin_trgm_ops);

create index idx_fleets_name
    on fleets using gin (name gin_trgm_ops);

create index idx_fleets_created_at
    on fleets (created_at desc);

create index idx_fleets_created_by
    on fleets using hash (created_by);

create index idx_fleets_updated_at
    on fleets (updated_at desc);

create index idx_fleets_updated_by
    on fleets using hash (updated_by);

insert into fleets(id, fleet_code, name)
values ('c2d40ef0-341a-4793-b1b3-f4e4f82ba9f2', 'A12345', '管理員')