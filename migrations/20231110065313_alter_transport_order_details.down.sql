alter table transport_order_details
    rename column product_name to goods_name;

alter index idx_transport_order_details_product_name rename to idx_transport_order_details_goods_name;

alter table transport_order_details
    drop column trailer_id;

drop index idx_transport_order_details_trailer_id;