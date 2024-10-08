alter table jasslin_data
    rename column driver_id to driver_code; -- 司機代號

alter table jasslin_data
    rename column io to io_status; -- IO狀態

alter table jasslin_data
    rename column status to gps_status; -- GPS狀態
