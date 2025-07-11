drop index idx_transport_orders_sequence;

alter table transport_orders
    drop column sequence;