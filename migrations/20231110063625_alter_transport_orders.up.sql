alter table transport_orders
    rename column shipper_id to client_id;

alter index idx_transport_orders_shipper_id rename to idx_transport_orders_client_id;

alter table transport_orders
    add deadline timestamp;

create index idx_transport_orders_deadline
    on transport_orders (deadline);