package main

import (
	"fmt"
	"net/http"
)

type Config struct {
	Port string `json: "port"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет, мир!"))
}

func main() {
	http.HandleFunc("/calculate", calculateHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
