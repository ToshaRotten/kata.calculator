package calculator_test

import (
	"fmt"
	"testing"

	"kata.calculator/services/calculate"
)

func TestCalculate(t *testing.T) {
	testCases := []struct {
		a        int
		b        int
		operator string
		expected int
		errMsg   string
	}{
		{2, 3, "+", 5, ""},
		{10, 5, "-", 5, ""},
		{4, 2, "*", 8, ""},
		{15, 3, "/", 5, ""},
		{8, 0, "/", 0, "деление на ноль"},
		{6, 4, "%", 0, "неверная операция: %"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d %s %d", tc.a, tc.operator, tc.b), func(t *testing.T) {
			result := calculate.Calculate(tc.a, tc.b, tc.operator)

			// Проверяем ожидаемый результат
			if result != tc.expected {
				t.Errorf("Ожидалось %d, получено %d", tc.expected, result)
			}
		})
	}
}
