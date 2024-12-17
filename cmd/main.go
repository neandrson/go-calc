package main

import (
	"fmt"
	"net/http"
)

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет, мир!"))
}

func main() {
	http.HandleFunc("/api/v1/calculate", calculateHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
