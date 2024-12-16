package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Создаём обработчик для маршрута "/api/v1/calculate"
	hello := http.HandlerFunc(helloHandler)

	// Применяем logging middleware к обработчику "/api/v1/calculate"
	mux.Handle("/api/v1/calculate", loggingMiddleware(hello))

	// Запускаем сервер на порту 8080
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
