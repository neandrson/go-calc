package calculation

import (
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		name        string
		expression  string
		expected    float64
		expectError bool
	}{
		{"Simple addition", "2+2", 4, false},
		{"Simple multiplication", "2*3", 6, false},
		{"Mixed operators", "2+3*4", 14, false},
		{"Division", "10/2", 5, false},
		{"Division by zero", "10/0", 0, true},
		{"Parentheses", "(2+3)*4", 20, false},
		{"Unmatched parentheses", "3+(4-2))", 0, true},
		{"Unmatched parentheses", "(2+3", 0, true},
		{"Invalid characters", "2+abc", 0, true},
		{"Empty expression", "", 0, true},
		{"Invalid expression", "3 + 5 *", 0, true},
		{"Complex expression", "3+(6/2)*5-4", 14, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Calc(tt.expression)
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}
			if !tt.expectError && result != tt.expected {
				t.Errorf("expected result: %v, got: %v", tt.expected, result)
			}
		})
	}
}
