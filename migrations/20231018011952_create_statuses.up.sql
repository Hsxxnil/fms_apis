CREATE TABLE statuses
(
    id         serial not null primary key, -- 表ID
    status     text   not null,             -- 狀態
    created_at timestamp default now(),
    created_by uuid,
    updated_at timestamp,
    updated_by uuid,
    deleted_at timestamp
);

create index idx_statuses_status
    on statuses (status);

create index idx_statuses_created_at
    on statuses (created_at desc);

create index idx_statuses_created_by
    on statuses using hash (created_by);

create index idx_statuses_updated_at
    on statuses (updated_at desc);

create index idx_statuses_updated_by
    on statuses using hash (updated_by);