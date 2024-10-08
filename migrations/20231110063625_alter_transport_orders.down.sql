alter table transport_orders
    rename column client_id to shipper_id;

alter index idx_transport_orders_client_id rename to idx_transport_orders_shipper_id;

alter table transport_orders
    drop column deadline;

drop index idx_transport_orders_deadline;