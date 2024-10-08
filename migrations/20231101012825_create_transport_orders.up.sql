CREATE TABLE transport_orders
(
    id                uuid      default uuid_generate_v4() not null
        primary key,                                                                         -- 表ID
    name              text                                 not null,                         -- 托運訂單名稱
    code              text                                 not null,                         -- 托運訂單單號
    shipper_id        uuid                                 not null references clients (id), -- 託運人ID
    origin            text                                 not null,                         -- 起運點
    destination       text                                 not null,                         -- 卸貨點
    transport_task_id uuid references transport_tasks (id),                                  -- 派工任務ID
    created_at        timestamp default now(),
    created_by        uuid,
    updated_at        timestamp,
    updated_by        uuid,
    deleted_at        timestamp
);

create index idx_transport_orders_name
    on transport_orders (name);

create index idx_transport_orders_code
    on transport_orders using gin (code gin_trgm_ops);

create index idx_transport_orders_shipper_id
    on transport_orders using hash (shipper_id);

create index idx_transport_orders_transport_task_id
    on transport_orders using hash (transport_task_id);

create index idx_transport_orders_origin
    on transport_orders using gin (origin gin_trgm_ops);

create index idx_transport_orders_destination
    on transport_orders using gin (destination gin_trgm_ops);

create index idx_transport_orders_created_at
    on transport_orders (created_at desc);

create index idx_transport_orders_created_by
    on transport_orders using hash (created_by);

create index idx_transport_orders_updated_at
    on transport_orders (updated_at desc);

create index idx_transport_orders_updated_by
    on transport_orders using hash (updated_by);