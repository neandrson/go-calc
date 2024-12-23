package main

import (
	"log"
	"net/http"

	"github.com/neandrson/go-calc/internal/application"
)

func main() {
	app := application.New()

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/calculate", application.CalcHandler)

	log.Println("Starting server on port", app.Config.Addr)
	if err := http.ListenAndServe(":"+app.Config.Addr, mux); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
