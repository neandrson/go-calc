package application

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalcHandler(t *testing.T) {
	tests := []struct {
		name       string
		payload    string
		wantStatus int
		method     string
		wantBody   map[string]interface{}
	}{
		{
			name:       "valid expression",
			payload:    `{"expression": "2+2"}`,
			wantStatus: http.StatusOK,
			method:     http.MethodPost,
			wantBody:   map[string]interface{}{"result": 4.0},
		},
		{
			name:       "invalid character in expression",
			payload:    `{"expression": "2+abc"}`,
			wantStatus: http.StatusUnprocessableEntity,
			method:     http.MethodPost,
			wantBody:   map[string]interface{}{"error": "Expression is not valid"},
		},
		{
			name:       "invalid json",
			payload:    `{"expr`,
			wantStatus: http.StatusUnprocessableEntity,
			method:     http.MethodPost,
			wantBody:   map[string]interface{}{"error": "Expression is not valid"},
		},
		{
			name:       "empty request body",
			payload:    `{}`,
			wantStatus: http.StatusUnprocessableEntity,
			method:     http.MethodPost,
			wantBody:   map[string]interface{}{"error": "Expression is not valid"},
		},
		{
			name:       "incorrect close-brackets",
			payload:    `{"expression": "1+(1))"}`,
			wantStatus: http.StatusUnprocessableEntity,
			method:     http.MethodPost,
			wantBody:   map[string]interface{}{"error": "unmatched parentheses"},
		},
		{
			name:       "incorrect open-brackets",
			payload:    `{"expression": "(2*3"}`,
			wantStatus: http.StatusUnprocessableEntity,
			method:     http.MethodPost,
			wantBody:   map[string]interface{}{"error": "unmatched parentheses"},
		},
		{
			name:       "incorrect method",
			payload:    `{"expression": "1+1"}`,
			wantStatus: http.StatusMethodNotAllowed,
			method:     http.MethodGet,
			wantBody:   map[string]interface{}{"error": "Method not allowed"},
		},
		{
			name:       "complex expression",
			payload:    `{"expression": "1+1*2"}`,
			wantStatus: http.StatusOK,
			method:     http.MethodPost,
			wantBody:   map[string]interface{}{"result": 3.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/api/v1/calculate", bytes.NewBuffer([]byte(tt.payload)))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			calculateHandler := http.HandlerFunc(CalcHandler)
			calculateHandler.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("expected status %d, got %d", tt.wantStatus, w.Code)
			}

			var response map[string]interface{}
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatalf("failed to parse response: %v", err)
			}

			for key, wantValue := range tt.wantBody {
				if gotValue, ok := response[key]; !ok || gotValue != wantValue {
					t.Errorf("expected %s = %v, got %v", key, wantValue, gotValue)
				}
			}
		})
	}
}

func TestConfigFromEnv(t *testing.T) {
	config := ConfigFromEnv()
	if config.Addr != "8080" {
		t.Errorf("expected port 8080, got %s", config.Addr)
	}
}

func TestNew(t *testing.T) {
	app := New()
	if app.config.Addr != "8080" {
		t.Errorf("expected port 8080, got %s", app.config.Addr)
	}
}
