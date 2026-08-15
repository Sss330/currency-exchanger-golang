package handler

import "currency-exchanger-golang/internal/model"

func GetAllCurrencies() []model.Currency {

	asd := []model.Currency{{Id: 1, Code: "asd", FullName: "Asd", Sign: "asd"}}
	return asd
}
