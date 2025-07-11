CREATE TABLE drivers
(
    id              uuid      default uuid_generate_v4() not null
        primary key,                                               -- 表ID
    name            text                                 not null, -- 姓名
    id_card_number  text,                                          -- 身份證字號
    employee_number text,                                          -- 員工編號
    phone_number    text,                                          -- 行動電話
    email           text,                                          -- 電子郵件
    address         text,                                          -- 連絡地址
    daily_cost      int,                                           -- 每日成本
    created_at      timestamp default now(),
    created_by      uuid,
    updated_at      timestamp,
    updated_by      uuid,
    deleted_at      timestamp
);

create index idx_drivers_sid
    on drivers using gin (name gin_trgm_ops);

create index idx_drivers_firm
    on drivers (id_card_number);

create index idx_drivers_model
    on drivers (employee_number);

create index idx_drivers_phone_number
    on drivers using gin (phone_number gin_trgm_ops);

create index idx_drivers_address
    on drivers using gin (address gin_trgm_ops);

create index idx_drivers_email
    on drivers using gin (email gin_trgm_ops);

create index idx_drivers_created_at
    on drivers (created_at desc);

create index idx_drivers_created_by
    on drivers using hash (created_by);

create index idx_drivers_updated_at
    on drivers (updated_at desc);

create index idx_drivers_updated_by
    on drivers using hash (updated_by);