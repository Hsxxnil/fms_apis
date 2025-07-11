alter table transport_order_details
    rename column goods_name to product_name;

alter index idx_transport_order_details_goods rename to idx_transport_order_details_product_name;

alter table transport_order_details
    add trailer_id uuid references trailers (id);

create index idx_transport_order_details_trailer_id
    on transport_order_details using hash (trailer_id);