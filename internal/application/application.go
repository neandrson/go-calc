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

/*type Request struct {
	Expression string `json: "expression"`
}*/

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	expression := data["expression"]

	result, err := calculation.Calc(expression)
	if err != nil {
		/*if err.Is(err, calculation.ErrInvalidExpression) {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				fmt.Fprintf(w, "error: %s", err.Error())
				w.WriteHeader(500)
			} else {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				//fmt.Fprintln(w, "error: unknow err", err.Error())
				//w.WriteHeader(404)
			}
		} else {*/
		//fmt.Fprintf(w, "result: %f", result)
		http.Error(w, "Internal error", http.StatusUnauthorized)
		return
	}

	response := map[string]string{"result": strconv.FormatFloat(result, 'f', 6, 64)}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
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

func (a *Application) RunServer() error {
	fmt.Println("Starting server on port 8080...")
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	err := http.ListenAndServe(":"+a.config.Addr, nil) // curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{\"expression\": \"2+2*2\"}"
	if err != nil {
		fmt.Println("Error starting server:", err)
		return err
	}
	return nil
}
