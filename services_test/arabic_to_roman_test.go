package arabic

import (
	"fmt"
	"testing"

	"kata.calculator/services"
)

func TestArabicToRoman(t *testing.T) {
	testCases := []struct {
		input          int
		expectedOutput string
		expectedError  string
	}{
		{1, "I", ""},            // Минимальное значение
		{3999, "MMMCMXCIX", ""}, // Максимальное значение
		{476, "CDLXXVI", ""},    // Пример: 476 = CDLXXVI
		{10, "X", ""},
		{11, "XI", ""},
		{50, "L", ""},
		{51, "LI", ""},
		{2000, "MM", ""},
		{2050, "MML", ""},
		{2051, "MMLI", ""},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d to Roman", tc.input), func(t *testing.T) {
			output := services.ArabicToRoman(tc.input)

			// Проверяем ожидаемый результат
			if output != tc.expectedOutput {
				t.Errorf(
					"Для числа %d ожидалось '%s', получено '%s'",
					tc.input,
					tc.expectedOutput,
					output,
				)
			}
		})
	}
}
