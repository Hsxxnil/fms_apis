CREATE TABLE transport_order_details
(
    id                 uuid      default uuid_generate_v4() not null
        primary key,                                                                                   -- 表ID
    goods_name         text                                 not null,                                  -- 品項
    unit_price         numeric,                                                                        -- 單價
    quantity           int,                                                                            -- 數量
    tonnage            int,                                                                            -- 噸數
    transport_order_id uuid                                 not null references transport_orders (id), -- 托運訂單ID
    created_at         timestamp default now(),
    created_by         uuid,
    updated_at         timestamp,
    updated_by         uuid,
    deleted_at         timestamp
);

create index idx_transport_order_details_goods
    on transport_order_details using gin (goods_name gin_trgm_ops);

create index idx_transport_order_details_transport_order_id
    on transport_order_details using hash (transport_order_id);

create index idx_transport_order_details_created_at
    on transport_order_details (created_at desc);

create index idx_transport_order_details_created_by
    on transport_order_details using hash (created_by);

create index idx_transport_order_details_updated_at
    on transport_order_details (updated_at desc);

create index idx_transport_order_details_updated_by
    on transport_order_details using hash (updated_by);