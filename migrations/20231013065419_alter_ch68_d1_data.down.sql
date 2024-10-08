drop index idx_ch68_d1_data_sid;

create index idx_ch68_d1_data_sid
    on ch68_d1_data using gin (sid gin_trgm_ops);