package main

import (
	"currency-exchanger-golang/internal/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.Check)

	handler.Check(mux, ":8080")

}
