package main

import (
	"currency-exchanger-golang/internal/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.Check)

	http.ListenAndServe(":8086", mux)
}
