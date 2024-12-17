package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type response struct {
	Name string `json:"name"`
}

func serveHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost/"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	//w.Write([]byte("Привет, мир!"))
}

func main() {

	http.HandleFunc("/api/v1/calculate", serveHandler)
	http.ListenAndServe(":8080", nil)
}
