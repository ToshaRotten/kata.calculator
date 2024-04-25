package Services

import (
	"fmt"
	"strconv"
	"strings"
)

// Функция для проверки формата введенного выражения
func ValidateExpression(input string) (int, int, string, error) {
	input = strings.TrimSpace(input)

	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		return 0, 0, "", fmt.Errorf("неверный формат выражения")
	}

	operand1Str := parts[0]
	operator := parts[1]
	operand2Str := parts[2]

	operand1, err1 := strconv.Atoi(operand1Str)
	operand2, err2 := strconv.Atoi(operand2Str)
	if err1 != nil || err2 != nil {
		return 0, 0, "", fmt.Errorf("неверный формат операндов")
	}

	return operand1, operand2, operator, nil
}

// если true - значит арабские, иначе false римские
func WhatKindOfExpression(input string) bool {
	parts := strings.Split(input, " ")
	if len(parts) == 1 {
		panic("Выдача паники, так как строка не является математической операцией.")
	}
	if len(parts) != 3 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	hasArabic := false
	hasRoman := false
	// Проверяем наличие цифр и римских цифр во входной строке
	for _, char := range input {
		if char >= '0' && char <= '9' {
			hasArabic = true
		} else if strings.ContainsRune("IVXLCDM", char) {
			hasRoman = true
		}
	}

	if hasArabic && hasRoman {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}

	// Возвращаем true, если это арабские цифры, иначе false (римские цифры)
	return hasArabic
}
