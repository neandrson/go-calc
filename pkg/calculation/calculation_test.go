package calculation_test

import (
	"testing"

	"github.com/neandrson/go-calc/pkg/calculation"
)

func TestCalculation(t *testing.T) {
	testCases := []struct {
		expression string
		expected   float64
	}{
		{"1+1", 1 + 1},
		{"3+3*6", 3 + 3*6},
		{"1+8/2*4", 1 + 8/2*4},
		{"(1+1) *2", (1 + 1) * 2},
		{"((1+4) * (1+2) +1) *4", ((1+4)*(1+2) + 1) * 4},
		{"(4+3+2)/(1+2) * 1 / 3", (4 + 3 + 2) / (1 + 2) * 1 / 3},
		{"((7+1) / (2+2) * 4) / 8 * (3 - ((4+1)*2)) -1", ((7+1)/(2+2)*4)/8*(3-((4+1)*2)) - 1},
		{"-1", -1},
		{"+5", 5},
		{"5+5+5+5+5", 5 + 5 + 5 + 5 + 5},
		{"(1)", 1},
		{"(1+2*(1) + 1)", (1 + 2*(1) + 1)},
		{"((1+2)*(5*(7+3) - 7 / (3+4) * (1+2)) - (8-1)) + (1 * (5-1 * (2+3)))", ((1+2)*(5*(7+3)-7/(3+4)*(1+2)) - (8 - 1)) + (1 * (5 - 1*(2+3)))},
		{"-1+2", -1 + 2},
		{"5+ -1", 5 + -1},
		{"5+ -5 + 7 - -6", 5 + -5 + 7 - -6},
		{"-(5+5)", -(5 + 5)},
		{"-9+9", -9 + 9},
		{"9*-1", 9 * -1},
		{"1*(1/1*-1)", 1 * 1 / 1 * -1},
		{"1*-1", 1 * -1},
	}
	for _, testCase := range testCases {
		t.Run(testCase.expression, func(t *testing.T) {
			result, err := calculation.Calc(testCase.expression)
			if err != nil {
				t.Errorf("Calc(%s) error: %v", testCase.expression, err)
			} else if result != testCase.expected {
				t.Errorf("Calc(%s) = %v, want %v", testCase.expression, result, testCase.expected)
			}
		})
	}
}
func TestCalculationErrors(t *testing.T) {
	testCases := []string{
		"1/0",
		"2*(1+9",
		"not numbs",
		"2r+1b",
		"1*(1+2*(1+2*(3+4) + 3 * (1+3) + 8 )",
		"1**2",
		"6^2",
		"((((((((((1)))))))))",
		"",
		"()",
		"*1",
		"-+",
		"-",
		"'1",
	}
	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			_, err := calculation.Calc(testCase)
			if err == nil {
				t.Errorf("Calc(%s) error is not nil", testCase)
			}
		})
	}
}
