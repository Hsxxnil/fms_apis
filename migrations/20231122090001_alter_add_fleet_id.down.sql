alter table clients
    drop column fleet_id;

alter table trailers
    drop column fleet_id;

alter table drivers
    drop column fleet_id;

alter table transport_tasks
    drop column fleet_id;

alter table transport_orders
    drop column fleet_id;