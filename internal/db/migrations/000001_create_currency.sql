-- +goose Up

create table currencies
(
    id        int primary key generated always as identity,
    code      varchar unique not null,
    full_name varchar not null,
    sign      varchar
);

-- +goouse Down

drop table currencies;