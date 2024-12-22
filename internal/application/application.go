package application

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/neandrson/go-calc/pkg/calculation"
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

// конфигурация сервера
func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080" // порт сервера
	}
	return config
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	//var data map[string]string
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer func() {
		err := recover()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("internal server error")
		}
	}()
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("expression is not valid")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//expression := data["expression"]

	result, err := calculation.Calc(req.Expression)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err.Error())
		http.Error(w, "Expression is not valid", http.StatusUnprocessableEntity)
		return
	}

	response := map[string]string{"result": strconv.FormatFloat(result, 'f', 0, 64)}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	log.Println("Expression:", req.Expression, " Result:", result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (a Application) Run() {
	log.Println("input expression")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Println("filed to read expression from console")
	}

	text = strings.TrimSpace(text)
	if text == "exit" {
		log.Println("application was successfully closed")
		return
	}

	result, err := calculation.Calc(text)
	if err != nil {
		fmt.Println(text, "calculation failed with error", err)
	} else {
		fmt.Println(text, "=", result)
	}
}

func responseHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

func (a *Application) RunServer() error {
	//fmt.Println("Starting server on port 8080...")
	log.Println("Starting server on port 8080...")
	http.HandleFunc("/", responseHandler)
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	err := http.ListenAndServe(":"+a.config.Addr, nil) // curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{\"expression\": \"2+2*2\"}"
	if err != nil {
		log.Println("Internal server error")
		return err
	}
	return nil
}
