CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists todo
(
    id           uuid    default uuid_generate_v4() not null constraint todo_pk primary key,
    name         varchar default ''                 not null constraint todo_uk unique,
    is_completed boolean default false              not null
);

