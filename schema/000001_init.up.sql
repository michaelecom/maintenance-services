CREATE TABLE ongoing_maintenance (
    id serial not null unique,
    service_market_id int not null,
    order_number varchar(15) not null unique,
    car_brand varchar(128),
    car_model varchar(128),
    car_number varchar(10)
);