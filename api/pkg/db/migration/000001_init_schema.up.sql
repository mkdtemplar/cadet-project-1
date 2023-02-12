CREATE TABLE IF NOT EXISTS public.users
(
    id    uuid         not null primary key,
    email varchar(100) not null unique,
    name  varchar(100)
);
CREATE TABLE IF NOT EXISTS public.user_preferences
(
    id     uuid not null primary key,
    user_country text,
    user_id      uuid CONSTRAINT fk_users_user_pref REFERENCES public.users on DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS public.vehicles
(
    id      uuid not null primary key,
    name    varchar(100),
    model   varchar(100),
    mileage double precision,
    user_id uuid CONSTRAINT fk_users_user_vehicle REFERENCES public.users on DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS public.ship_ports
(
    _key      varchar(6),
    name      text,
    city      text,
    country   text,
    latitude  double precision,
    longitude double precision,
    province  varchar(40),
    timezone  varchar(40),
    unlocks_0 varchar(6),
    code      integer,
    alias_0   varchar(50),
    alias_1   varchar(50),
    alias_2   varchar(50),
    unlocks_1 varchar(50)
);
