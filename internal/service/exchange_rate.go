package service

import (
	"context"
	"currency-exchanger-golang/internal/model"
	"currency-exchanger-golang/internal/repository"
)

func FindAll(ctx context.Context) []model.Currency {
	repo := repository.NewCurrencyPostgresRepository()

	curns, err := repo.FindAll(ctx, repo)
	if err != nil {

	}

	return curns
}
