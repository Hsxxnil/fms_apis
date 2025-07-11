alter table transport_orders
    add column sequence int; --順序

create index idx_transport_orders_sequence
    on transport_orders (sequence);
