package calculate

import (
	"errors"

	"kata.calculator/services/calculate/operations"
)

func Calculate(a int, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return operations.Add(a, b), nil
	case "-":
		return operations.Sub(a, b), nil
	case "*":
		return operations.Mul(a, b), nil
	case "/":
		return operations.Div(a, b)
	default:
		return 0, errors.New("wrong operator")
	}
}
