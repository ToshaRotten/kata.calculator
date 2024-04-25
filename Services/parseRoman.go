package Services

import (
	"fmt"
	"strings"
)

func RomanToArabic(roman string) (int, error) {
	romanValues := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var prevValue int
	var total int

	for i := len(roman) - 1; i >= 0; i-- {
		currentValue := romanValues[rune(roman[i])]

		if currentValue < prevValue {
			total -= currentValue
		} else {
			total += currentValue
		}
		prevValue = currentValue
	}
	return total, nil
}

func ArabicToRoman(num int) (string, error) {
	if num <= 0 || num > 3999 {
		return "", fmt.Errorf("неверное значение: %d, арабские числа должны быть от 1 до 3999", num)
	}

	// Массив символов римских цифр и их соответствующих значения
	romanDigits := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	romanNum := ""
	for _, digit := range romanDigits {
		for num >= digit.value {
			romanNum += digit.digit
			num -= digit.value
		}
	}

	return romanNum, nil
}

func SeparateRoman(input string) (string, string, string, error) {
	input = strings.TrimSpace(input)

	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("неверный формат выражения")
	}

	return parts[0], parts[2], parts[1], nil
}
