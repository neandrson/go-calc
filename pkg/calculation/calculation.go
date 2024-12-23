package calculation

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// Функция для вычисления выражения
func Calc(expression string) (float64, error) {
	tokens, err := tokenize(expression)
	if err != nil {
		return 0, err
	}

	rpn, err := toRPN(tokens)
	if err != nil {
		return 0, err
	}

	return calculateRPN(rpn)
}

// Функция для токенизации выражения. Берет строку с выражением и возвращает массив токенов или ошибку.
func tokenize(expression string) ([]string, error) {
	var tokens []string
	var currentNumber string

	expression = strings.Replace(expression, " ", "", -1)

	for _, ch := range expression {
		if unicode.IsDigit(ch) || ch == '.' {
			currentNumber += string(ch)
			continue
		}

		if currentNumber != "" {
			tokens = append(tokens, currentNumber)
			currentNumber = ""
		}

		if strings.Contains("+-*/()", string(ch)) {
			tokens = append(tokens, string(ch))
		} else if !unicode.IsDigit(ch) && ch != '.' {
			return nil, errors.New("invalid character")
		}
	}

	if currentNumber != "" {
		tokens = append(tokens, currentNumber)
	}

	return tokens, nil
}

// Функция для преобразования токенов в обратную польскую нотацию
func toRPN(tokens []string) ([]string, error) {
	var rpn []string
	var stack []string

	priority := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	for _, token := range tokens {
		if _, err := strconv.ParseFloat(token, 64); err == nil {
			rpn = append(rpn, token)
		} else if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				rpn = append(rpn, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 || stack[len(stack)-1] != "(" {
				return nil, errors.New("unmatched parentheses")
			}
			stack = stack[:len(stack)-1]
		} else {
			for len(stack) > 0 && priority[stack[len(stack)-1]] >= priority[token] {
				rpn = append(rpn, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		}
	}
	for len(stack) > 0 {
		if stack[len(stack)-1] == "(" {
			return nil, errors.New("unmatched parentheses")
		}
		rpn = append(rpn, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return rpn, nil
}

// Функция для вычисления выражения в обратной польской нотации
func calculateRPN(rpn []string) (float64, error) {
	var stack []float64

	for _, token := range rpn {
		if value, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, value)
		} else {
			if len(stack) < 2 {
				return 0, errors.New("invalid expression")
			}
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, errors.New("division by zero")
				}
				stack = append(stack, a/b)
			}
		}
	}
	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}
	return stack[0], nil
}
