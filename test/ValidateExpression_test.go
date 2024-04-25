package main

import (
	"kata.calculator/Services"
	"testing"
)

func TestValidateExpression(t *testing.T) {
	testCases := []struct {
		input            string
		expectedOperand1 int
		expectedOperand2 int
		expectedOperator string
		expectedError    string
	}{
		{"2 + 3", 2, 3, "+", ""},      // Корректное выражение
		{"  10 - 5 ", 10, 5, "-", ""}, // Корректное выражение с пробелами
		{"4 * 6", 4, 6, "*", ""},      // Корректное выражение
		{"7 / 0", 0, 0, "", "неверный формат операндов"},     // Деление на ноль
		{"abc + 5", 0, 0, "", "неверный формат операндов"},   // Некорректный формат операнда
		{"8 / 2 / 4", 0, 0, "", "неверный формат выражения"}, // Некорректный формат выражения
		{"6", 0, 0, "", "неверный формат выражения"},         // Некорректный формат выражения
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			operand1, operand2, operator, err := Services.ValidateExpression(tc.input)

			// Проверяем ожидаемые значения операндов и оператора
			if operand1 != tc.expectedOperand1 || operand2 != tc.expectedOperand2 || operator != tc.expectedOperator {
				t.Errorf("Для входа '%s' ожидались операнды %d, %d и оператор '%s'", tc.input, tc.expectedOperand1, tc.expectedOperand2, tc.expectedOperator)
			}

			// Проверяем ошибку
			if err != nil {
				if err.Error() != tc.expectedError {
					t.Errorf("Для входа '%s' ожидалась ошибка '%s', получена ошибка: '%s'", tc.input, tc.expectedError, err.Error())
				}
			} else {
				if tc.expectedError != "" {
					t.Errorf("Для входа '%s' ожидалась ошибка '%s', получена ошибка: <nil>", tc.input, tc.expectedError)
				}
			}
		})
	}
}
