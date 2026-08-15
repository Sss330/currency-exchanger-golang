-- +goose Up

create table if not exists currencies
(
    id        int primary key generated always as identity,
    code      varchar unique not null,
    full_name varchar        not null,
    sign      varchar
);

-- +goose Down

drop table if exists currencies;