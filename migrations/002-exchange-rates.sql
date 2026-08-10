create table exchange_rates(
    id int primary key,
    base_currency_id int REFERENCES id,
    target_currency int references id,
    rate decimal(6)
)