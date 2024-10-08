drop index idx_vehicles_style;
drop index idx_vehicles_billing_type;
drop index idx_vehicles_fuel_type;
drop index idx_vehicles_tax_id;

alter table vehicles
    drop column daily_cost;

alter table vehicles
    drop column weight;

alter table vehicles
    drop column style;

alter table vehicles
    drop column billing_type;

alter table vehicles
    drop column fuel_type;

alter table vehicles
    drop column fuel_consumption;

alter table vehicles
    drop column current_mileage;

alter table vehicles
    drop column tax_id;
