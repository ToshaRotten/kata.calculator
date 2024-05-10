package main

import (
	"testing"

	"kata.calculator/Services"
)

func TestWhatKindOfExpression(t *testing.T) {
	testCases := []struct {
		input         string
		expected      bool
		expectedPanic bool
	}{
		{"2 + 3", true, false},              // Арабские цифры
		{"IV + VI", false, false},           // Римские цифры
		{"7 * III", true, false},            // Арабские и римские цифры (паника)
		{"II - IX", false, false},           // Римские цифры
		{"5 + V", true, true},               // Арабские и римские цифры (паника)
		{"8 / 2", true, false},              // Арабские цифры
		{"X * IV", false, false},            // Римские цифры
		{"3 * 5", true, false},              // Арабские цифры
		{"III + 7", false, false},           // Римские и арабские цифры (паника)
		{"10 - 2", true, false},             // Арабские цифры
		{"VII / IX", false, false},          // Римские цифры
		{"IV / 2", false, true},             // Римские и арабские цифры (паника)
		{"invalid expression", false, true}, // Некорректная строка (паника)
		{"3 + 4 + 5", false, true},          // Некорректный формат (паника)
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tc.expectedPanic {
					t.Errorf("Ожидалась паника: %v, получено: %v", tc.expectedPanic, r)
				}
			}()

			result := Services.WhatKindOfExpression(tc.input)

			if result != tc.expected {
				t.Errorf("Для входа '%s' ожидалось %t, получено %t", tc.input, tc.expected, result)
			}
		})
	}
}
