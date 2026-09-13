package main

import (
	"context"
	db2 "currency-exchanger-golang/internal/db"
	"currency-exchanger-golang/internal/db/migrations"
	"currency-exchanger-golang/internal/handler"
	"currency-exchanger-golang/internal/webapp"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()

	dsn := os.Getenv("DATABASE_URL")

	db, err := db2.Open(ctx, dsn)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := migrations.Up(ctx, db); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	webapp.Register(mux)

	mux.HandleFunc("GET /all-currencyes", handler.GetAllCurrencies(db))

	if err := http.ListenAndServe(":8086", mux); err != nil {
		log.Printf("HTTP server failed: %v", err)
	}
}
