package repository

import (
	"context"
	"currency-exchanger-golang/internal/model"
	"database/sql"
)

type CurrencyRepository struct {
	db *sql.DB
}

func NewCurrencyPostgresRepository(db *sql.DB) *CurrencyRepository {
	return &CurrencyRepository{db: db}
}

func (repo *CurrencyRepository) FindAll(ctx context.Context) ([]model.Currency, error) {

	rows, err := repo.db.QueryContext(ctx, "select * from currencies")

	if err != nil {
		return nil, err
	}
	rows.Scan()

	currencies := make([]model.Currency, 0)

	for rows.Next() {
		var currency model.Currency

		err := rows.Scan(
			&currency.Id,
			&currency.Code,
			&currency.FullName,
			&currency.Sign,
		)
		if err != nil {
			return nil, err
		}

		currencies = append(currencies, currency)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return currencies, nil
}

func (repo *CurrencyRepository) Create(ctx context.Context, currency model.Currency) (model.Currency, error) {
	result, _ := repo.db.QueryContext(ctx, "insert into currencies VALUES currency")
	result.Scan()
	if result.Err() != nil {
		return model.Currency{}, result.Err()
	}

	var currency model.Currency

	for result.Next() {

		err := result.Scan(
			&currency.Id,
			&currency.Code,
			&currency.FullName,
			&currency.Sign,
		)
		if err != nil {
			return model.Currency{}, err
		}

		result.Scan()
	}
	return result.Next().(model.Currency)
}

func GetBuCode() model.Currency {
	return model.Currency{}
}

func delete(id int) {

}

type Currency interface {
	FindAll(ctx context.Context) ([]model.Currency, error)
	Delete(id int)
	Create(currency model.Currency)
	GetByCode(code string) model.Currency
}
