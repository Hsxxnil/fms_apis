create table user_roles
(
    id         uuid      default uuid_generate_v4() not null
        primary key,                                                                 -- 表ID
    user_id    uuid                                 not null references users (id),  -- 使用者ID
    role_id    uuid                                 not null references roles (id),  -- 角色ID
    fleet_id   uuid                                 not null references fleets (id), -- 車隊ID
    created_at timestamp default now(),
    created_by uuid,
    updated_at timestamp,
    updated_by uuid,
    deleted_at timestamp
);

create index idx_user_roles_id
    on user_roles using hash (id);

create index idx_user_roles_user_id
    on user_roles using hash (user_id);

create index idx_user_roles_role_id
    on user_roles using hash (role_id);

create index idx_user_roles_fleet_id
    on user_roles using hash (fleet_id);

create index idx_user_roles_created_at
    on user_roles (created_at desc);

create index idx_user_roles_created_by
    on user_roles using hash (created_by);

create index idx_user_roles_updated_at
    on user_roles (updated_at desc);

create index idx_user_roles_updated_by
    on user_roles using hash (updated_by);

insert into user_roles(id, user_id, role_id, fleet_id)
values ('561159b0-85a6-4c25-ab60-c2848cd57e00', 'a1bb0141-68e3-420c-8a92-9332fc21bd25',
        'd56fc184-9441-4396-be6c-d48580650171', 'c2d40ef0-341a-4793-b1b3-f4e4f82ba9f2')