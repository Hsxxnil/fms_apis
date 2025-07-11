alter table z1_data
    rename to avema_data;

alter index z1_data_pkey rename to avema_data_pkey;

alter index idx_z1_data_date_time rename to idx_avema_data_date_time;

alter index idx_z1_data_server_time rename to idx_avema_data_server_time;

alter index idx_z1_data_sid rename to idx_avema_data_sid;

