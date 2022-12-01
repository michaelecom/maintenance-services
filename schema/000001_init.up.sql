-- Active: 1669879880761@@91.220.109.55@5434@postgres
-- create tables
CREATE TABLE ongoing_maintenance (
    id serial not null unique,
    service_market_id int not null,
    order_number varchar(30) not null unique,
    car_brand varchar(128),
    car_model varchar(128),
    car_number varchar(10)
);
CREATE TABLE service_markets (
    id serial not null unique,
    service_market_id int not null unique,
    service_market_name varchar(128) not null
);
CREATE TABLE logging (
    id serial not null unique,
    datetime timestamp not null,
    service_market_id int not null,
    event_type text
);