create table if not exists users
(
    id         bigserial primary key,
    email      text unique,
    password   varchar(100) not null,
    last_login timestamp with time zone default CURRENT_TIMESTAMP
);

create table if not exists user_preferences
(
    id         bigserial constraint user_preferences_pk primary key,
    country    varchar(100),
    user_id_fk bigserial constraint user_preferences_users_null_fk  references users
);

create table if not exists "ships-routes"
(
    _key      text,
    name      text,
    city      text,
    country   text,
    latitude  numeric,
    longitude numeric,
    province  text,
    timezone  text,
    unlocs_0  text,
    code      integer,
    alias_0   text,
    alias_1   text,
    alias_2   text,
    unlocs_1  text
);
