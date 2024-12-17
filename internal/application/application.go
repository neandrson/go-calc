package application

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Addr string
}

type Application struct {
	config *Config
}

type Request struct {
	Expression string `json: "expression"`
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := calculation.Calc(request.Expression)
	if err != nil {
		fmt.Fprintf(w, "err^ %s", err.Error())
	} else {
		fmt.Fprintf(w, "result: %f", result)
	}
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

func (a Application) Run() {
	log.Println("input expression")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString("\n")
	if err != nil {
		log.Println("filed ti read expression from console")
	}

	text = strings.TrimSpace(text)
	if text == "exit" {
		log.Println("application was successfully closed")
		return nil
	}

	result, err := calculation.Calc(text)
	if err != nil {
		fmt.Println(text, "calculation failed wit error", err)
	} else {
		fmt.Println(text, "=", result)
	}
}

func (a *Application) RunServer() error {
	http.HandleFunc("/", CalcHandler)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
