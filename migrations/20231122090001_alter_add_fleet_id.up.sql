alter table clients
    add column fleet_id uuid references fleets (id);

alter table trailers
    add column fleet_id uuid references fleets (id);

alter table drivers
    add column fleet_id uuid references fleets (id);

alter table transport_tasks
    add column fleet_id uuid references fleets (id);

alter table transport_orders
    add column fleet_id uuid references fleets (id);