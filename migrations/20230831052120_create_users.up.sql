create table users
(
    id            uuid      default uuid_generate_v4() not null
        primary key,                                             -- 表ID
    user_name     text      default ''::text           not null, -- 帳號
    password      text      default ''::text           not null, -- 密碼
    name          text      default ''::text           not null, -- 姓名
    phone_number1 text,                                          -- 電話1
    phone_number2 text,                                          -- 電話2
    email         text,                                          -- 電子郵件
    created_at    timestamp default now(),
    created_by    uuid,
    updated_at    timestamp,
    updated_by    uuid,
    deleted_at    timestamp
);

create index idx_users_id
    on users using hash (id);

create index idx_users_user_name
    on users using gin (user_name gin_trgm_ops);

create index idx_users_name
    on users using gin (name gin_trgm_ops);

create index idx_users_phone_number1
    on users using gin (phone_number1 gin_trgm_ops);

create index idx_users_phone_number2
    on users using gin (phone_number2 gin_trgm_ops);

create index idx_users_email
    on users using gin (email gin_trgm_ops);

create index idx_users_created_at
    on users (created_at desc);

create index idx_users_created_by
    on users using hash (created_by);

create index idx_users_updated_at
    on users (updated_at desc);

create index idx_users_updated_by
    on users using hash (updated_by);

insert into users(id, user_name, password, name)
values ('a1bb0141-68e3-420c-8a92-9332fc21bd25', 'admin',
        '9HXSglPqDWrOyA29croTTu8O8ahmj2EMHhxrsfzrEpJBVykaIkDJ211tJ03aq25Q2iHvkECACPDI/yJXiDsRQDojG1iLqTMQp3nUSmfV/9Yhc3i+ovXLuiRoapCluqw4oxkiuLtqlQMivNTnphmOF+iHnu6sz8N6aouA3mOS89aSoPpHwbWbo4ilh3sPIyEnwLT9npq3ICQwP7FxXPFxaw==',
        '管理員')