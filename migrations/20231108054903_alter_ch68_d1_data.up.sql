alter index ch68_d1_data_pkey rename to poison_data_pkey;

alter index idx_ch68_d1_data_date_time rename to idx_poison_data_date_time;

alter index idx_ch68_d1_data_server_time rename to idx_poison_data_server_time;

alter index idx_ch68_d1_data_sid rename to idx_poison_data_sid;

alter table ch68_d1_data
    rename to poison_data;

