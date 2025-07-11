alter table vehicles
    add tax_id text; -- 車行統編

alter table vehicles
    add current_mileage int; -- 目前里程

alter table vehicles
    add fuel_consumption numeric; -- 油耗量(公里/公升)

alter table vehicles
    add fuel_type text; -- 油耗種類

alter table vehicles
    add billing_type text; -- ETC計費車種

alter table vehicles
    add style text; -- 車種

alter table vehicles
    add weight text; -- 車重(噸)

alter table vehicles
    add daily_cost int; -- 每日成本

create index idx_vehicles_tax_id
    on vehicles (tax_id);

create index idx_vehicles_fuel_type
    on vehicles (fuel_type);

create index idx_vehicles_billing_type
    on vehicles (billing_type);

create index idx_vehicles_style
    on vehicles (style);
