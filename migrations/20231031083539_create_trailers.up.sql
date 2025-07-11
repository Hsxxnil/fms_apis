CREATE TABLE trailers
(
    id         uuid      default uuid_generate_v4() not null
        primary key,                                          -- 表ID
    code       text                                 not null, -- 板台號碼
    created_at timestamp default now(),
    created_by uuid,
    updated_at timestamp,
    updated_by uuid,
    deleted_at timestamp
);

create index idx_trailers_code
    on trailers using gin (code gin_trgm_ops);

create index idx_trailers_created_at
    on trailers (created_at desc);

create index idx_trailers_created_by
    on trailers using hash (created_by);

create index idx_trailers_updated_at
    on trailers (updated_at desc);

create index idx_trailers_updated_by
    on trailers using hash (updated_by);