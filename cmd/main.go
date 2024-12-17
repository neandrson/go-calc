package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Name string `json:"name"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	result := response{}
	result.Name = r.URL.Query().Get("name")
	jsonOut, _ := json.Marshal(result)
	log.Print(string([]byte(jsonOut)))
	json.NewEncoder(w).Encode(jsonOut)
	w.Write([]byte("Привет, мир!"))
}

func main() {
	apiURL := "https://localhost/api/v1/calculate"

	resp, err := http.Post(apiURL, "application/json", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
    fmt.Println(resp)
	defer resp.Body.Close()
}
