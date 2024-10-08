create index idx_ch68_d1_data_date_time
    on ch68_d1_data (date_time desc);

create index idx_ch68_d1_data_server_time
    on ch68_d1_data (server_time desc);

create index idx_ch68_d1_data_sid
    on ch68_d1_data using gin (sid gin_trgm_ops);