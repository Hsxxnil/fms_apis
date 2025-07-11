alter table avema_data
    rename to z1_data;

alter index avema_data_pkey rename to z1_data_pkey;

alter index idx_avema_data_date_time rename to idx_z1_data_date_time;

alter index idx_avema_data_server_time rename to idx_z1_data_server_time;

alter index idx_avema_data_sid rename to idx_z1_data_sid;

