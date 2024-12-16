package main

import (
	"net/http"
)

func main() {
	//mux := http.NewServeMux()

	// Создаём обработчик для маршрута "/api/v1/calculate"
	//hello := http.HandlerFunc(helloHandler)

	// Применяем logging middleware к обработчику "/api/v1/calculate"
	http.Handle("/api/v1/calculate", http.HandlerFunc())

	// Запускаем сервер на порту 8080
	http.ListenAndServe(":8080", nil)
}
