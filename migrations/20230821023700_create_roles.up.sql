CREATE
    EXTENSION IF NOT EXISTS "pg_trgm";
CREATE
    EXTENSION IF NOT EXISTS "uuid-ossp";

create table roles
(
    id           uuid      default uuid_generate_v4() not null
        primary key,                                            -- 表ID
    display_name text      default ''::text           not null, -- 展示名稱
    name         text      default ''::text           not null, -- 角色名稱
    is_enable    boolean   default true               not null, -- 是否啟用
    created_at   timestamp default now(),
    created_by   uuid,
    updated_at   timestamp,
    updated_by   uuid,
    deleted_at   timestamp
);

create index idx_roles_id
    on roles using hash (id);

create index idx_roles_display_name
    on roles using gin (display_name gin_trgm_ops);

create index idx_roles_name
    on roles using gin (name gin_trgm_ops);

create index idx_roles_created_at
    on roles (created_at desc);

create index idx_roles_created_by
    on roles using hash (created_by);

create index idx_roles_updated_at
    on roles (updated_at desc);

create index idx_roles_updated_by
    on roles using hash (updated_by);

insert into roles(id, display_name, name)
values ('d56fc184-9441-4396-be6c-d48580650171', '管理員', 'admin')