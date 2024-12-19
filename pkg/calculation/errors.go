package calculation

import "errors"

var (
	ErrInvalidExpression = errors.New("Expression is not valid")
	ErrDivisionByZero    = errors.New("division by zero")
	ErrInternalServer    = errors.New("Internal server error")
)
