create schema if not exists core;

create table if not exists core.danger_zones
(
    device_id  varchar(64)      not null
    constraint dangerzones_pk
    primary key,
    company_id varchar(64)      not null,
    radius     double precision not null,
    longitude  double precision not null,
    latitude   double precision not null,
    end_ts     integer
    );

create index if not exists dangerzones_company_id_index
    on core.danger_zones (company_id);

create index if not exists dangerzones_device_id_index
    on core.danger_zones (device_id);

