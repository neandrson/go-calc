package calculation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrExpressionValid = errors.New("Expression is not valid")

/*func stringToFloat64(str string) float64 {
	degree := float64(1)
	var res float64 = 0
	var invers bool = false
	for i := len(str); i > 0; i-- {
		if str[i-1] == '-' {
			invers = true
		} else {
			res += float64(9-int('9'-str[i-1])) * degree
			degree *= 10
		}
	}
	if invers {
		res = 0 - res
	}
	return res
}*/

func Calc(expression string) (float64, error) {

	expression = strings.ReplaceAll(expression, " ", "")
	if !isValid(expression) {
		return 0, fmt.Errorf("%w", ErrExpressionValid)
	}
	postfix := infixToPostfix(expression)
	result, err := evaluatePostfix(postfix)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func isValid(expression string) bool {

	validChars := "0123456789+-*/()"
	for _, char := range expression {
		if !strings.Contains(validChars, string(char)) {
			return false
		}
	}
	openCount := strings.Count(expression, "(")
	closeCount := strings.Count(expression, ")")
	//if openCount != closeCount {
	//	return false
	//}
	return openCount == closeCount
}

func infixToPostfix(expression string) []string {
	var postfix []string
	var stack []string

	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	tokens := strings.Split(expression, "")

	for _, token := range tokens {
		if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			for stack[len(stack)-1] != "(" {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		} else if _, isOperator := precedence[token]; isOperator {
			for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[token] {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		} else {
			postfix = append(postfix, token)
		}
	}

	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return postfix
}

func evaluatePostfix(postfix []string) (float64, error) {
	var stack []float64

	for _, token := range postfix {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, fmt.Errorf("%w", ErrExpressionValid)
			}

			num2 := stack[len(stack)-1]
			num1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				if num2 == 0 {
					return 0, fmt.Errorf("%w", ErrExpressionValid)
				}
				stack = append(stack, num1/num2)
			}
		}
	}
	//if len(stack) != 1 {
	//	return stack[0], nil //0, fmt.Errorf("%w", ErrExpressionValid)
	//}

	return stack[0], nil
}
