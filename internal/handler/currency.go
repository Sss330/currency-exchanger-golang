package handler

import (
	"currency-exchanger-golang/internal/service"
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetAllCurrencies(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currencyes := service.GetAll(r.Context(), db)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(currencyes); err != nil {
			return
		}
	}
}
