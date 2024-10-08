create index idx_z1_data_date_time
    on z1_data (date_time desc);

create index idx_z1_data_server_time
    on z1_data (server_time desc);

create index idx_z1_data_sid
    on z1_data using gin (sid gin_trgm_ops);