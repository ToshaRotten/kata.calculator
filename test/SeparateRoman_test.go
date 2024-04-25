package main

import (
	"kata.calculator/Services"
	"testing"
)

func TestSeparateRoman(t *testing.T) {
	testCases := []struct {
		input            string
		expectedOperand1 string
		expectedOperand2 string
		expectedOperator string
		expectedError    string
	}{
		{"II + V", "II", "V", "+", ""},     // Корректное выражение
		{"X - IV", "X", "IV", "-", ""},     // Корректное выражение с пробелами
		{"III * VI", "III", "VI", "*", ""}, // Корректное выражение
		{"VII / IX", "VII", "IX", "/", ""}, // Корректное выражение
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			operand1, operand2, operator := Services.SeparateRoman(tc.input)

			// Проверяем ожидаемые значения операндов и оператора
			if operand1 != tc.expectedOperand1 || operand2 != tc.expectedOperand2 || operator != tc.expectedOperator {
				t.Errorf("Для входа '%s' ожидались операнды '%s', '%s' и оператор '%s'", tc.input, tc.expectedOperand1, tc.expectedOperand2, tc.expectedOperator)
			}

		})
	}
}
