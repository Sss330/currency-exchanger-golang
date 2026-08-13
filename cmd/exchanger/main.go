package main

import (
	"context"
	db2 "currency-exchanger-golang/internal/db"
	"currency-exchanger-golang/internal/db/migrations"
	"currency-exchanger-golang/internal/handler"
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

	mux.HandleFunc("GET /health", handler.Check)

	http.ListenAndServe(":8086", mux)
}
