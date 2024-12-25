package application

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/neandrson/go-calc/pkg/calculation"
)

// Структура приложения, содержит конфигурацию
type Application struct {
	Config *Config
}

// Структура для хранения конфигурации программы, содержит порт
type Config struct {
	Addr string
}

// Функция для получения конфигурации из переменных окружения
func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

// Функция для создания экземпляра приложения
func New() *Application {
	return &Application{
		Config: ConfigFromEnv(),
	}
}

// Структура для хранения запроса пользователя
type calculateRequest struct {
	Expression string `json:"expression"`
}

// Структура для хранения ответа пользователю
type calculateResponse struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

// Основная логика програамы. Принимает запросы, проверяет на ошибки и выдает ответ. Также добавлено логгирование для просмотра операций в терминале.
func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(calculateResponse{Error: "Method not allowed"})
		return
	}

	var req calculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(calculateResponse{Error: "Expression is not valid"})
		return
	}

	result, err := calculation.Calc(req.Expression)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(calculateResponse{Error: err.Error()})
		return
	}

	log.Println("New calculation. Expression:", req.Expression, " Result:", result)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(calculateResponse{Result: result})
}

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(calculateResponse{Error: "Internal Server Error"})
}
