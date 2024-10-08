CREATE TABLE clients
(
    id           uuid      default uuid_generate_v4() not null
        primary key,                                            -- 表ID
    name         text                                 not null, -- 客戶名稱
    phone_number text,                                          -- 客戶電話
    created_at   timestamp default now(),
    created_by   uuid,
    updated_at   timestamp,
    updated_by   uuid,
    deleted_at   timestamp
);

create index idx_clients_name
    on clients using gin (name gin_trgm_ops);

create index idx_clients_phone_number
    on clients using gin (phone_number gin_trgm_ops);

create index idx_clients_created_at
    on clients (created_at desc);

create index idx_clients_created_by
    on clients using hash (created_by);

create index idx_clients_updated_at
    on clients (updated_at desc);

create index idx_clients_updated_by
    on clients using hash (updated_by);