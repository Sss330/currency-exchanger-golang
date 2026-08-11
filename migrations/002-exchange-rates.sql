create table exchange_rates(
    id int primary key,
    base_currency_id int REFERENCES id,
    target_currency_id int references id,
    rate decimal(6)
)


create unique index asd = (base_currency_id, target_currency)