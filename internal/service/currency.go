package service

import (
	"context"
	"currency-exchanger-golang/internal/model"
	"currency-exchanger-golang/internal/repository"
	"database/sql"
)

func GetAll(ctx context.Context, db *sql.DB) []model.Currency {
	currencyRepo := repository.NewCurrencyPostgresRepository(db)

	currencyes, err := currencyRepo.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	if currencyes == nil {
		return currencyes
	}

	return nil
}

func GetById(ctx context.Context, db *sql.DB, id int) model.Currency {
	currencyRepo := repository.NewCurrencyPostgresRepository(db)
}

func Create(ctx context.Context, db *sql.DB, currency model.Currency) model.Currency {
	currencyRepo := repository.NewCurrencyPostgresRepository(db)
	currencyRepo.Create()
}
