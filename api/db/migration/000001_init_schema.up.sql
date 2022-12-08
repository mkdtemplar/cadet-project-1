create table if not exists users
(
    id    bigserial primary key,
    email text unique,
    name  varchar(100)
);
create table if not exists user_preferences
(
    id      bigserial
        constraint user_preferences_pk
            primary key,
    country varchar(100),
    user_id integer not null constraint user_preferences_users_null_fk references public.users
);
create table if not exists "ships-routes"
(
    _key      varchar(6),
    name      varchar(50),
    city      varchar(50),
    country   varchar(50),
    latitude  double precision,
    longitude double precision,
    province  varchar(40),
    timezone  varchar(40),
    unlocs_0  varchar(6),
    code      integer,
    alias_0   varchar(50),
    alias_1   varchar(50),
    alias_2   varchar(50),
    unlocs_1  varchar(50)
);
