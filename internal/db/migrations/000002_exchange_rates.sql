create table exchange_rates
(
    id                 int primary key generated always as identity,
    base_currency_id   int references currencies(id) not null,
    target_currency_id int references currencies(id) not null,
    rate               decimal(3) not null
)


create unique index asd on exchange_rates (base_currency_id, target_currency_id)