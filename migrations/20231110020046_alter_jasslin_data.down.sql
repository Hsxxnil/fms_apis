alter table jasslin_data
    rename column driver_code to driver_id;

alter table jasslin_data
    rename column io_status to io;

alter table jasslin_data
    rename column gps_status to status;