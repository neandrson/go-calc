package calculation

import "errors"

var (
	ErrInvalidExpression = errors.New("expression is not valid")
	ErrDivisionByZero    = errors.New("division by zero")
	ErrInternalServer    = errors.New("internal server error")
)
