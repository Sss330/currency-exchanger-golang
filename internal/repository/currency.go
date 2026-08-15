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

func (repo *CurrencyRepository) FindAll(ctx context.Context, r *CurrencyRepository) ([]model.Currency, error) {

	result, _ := r.db.QueryContext(ctx, "select * from currencies")

	result.Scan()

	return nil, nil
}

func delete(id int) {

}

func Create(currency model.Currency) {

}

func GetBuCode() model.Currency {
	return model.Currency{}
}

type Currency interface {
	FindAll(ctx context.Context) ([]model.Currency, error)
	Delete(id int)
	Create(currency model.Currency)
	GetByCode(code string) model.Currency
}
