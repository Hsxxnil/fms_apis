drop index idx_z1_data_sid;

create index idx_z1_data_sid
    on z1_data using gin (sid gin_trgm_ops);