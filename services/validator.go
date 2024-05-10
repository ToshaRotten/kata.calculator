package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Функция для проверки формата введенного выражения
func ValidateExpression(input string) (int, int, string) {
	input = strings.TrimSpace(input)

	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		panic("неверный формат выражения")
	}

	operand1Str := parts[0]
	operator := parts[1]
	operand2Str := parts[2]

	operand1, err1 := strconv.Atoi(operand1Str)
	operand2, err2 := strconv.Atoi(operand2Str)
	if err1 != nil || err2 != nil {
		panic("неверный формат операндов")
	}

	if operand1 < 0 || operand1 > 10 || operand2 < 0 || operand2 > 10 {
		panic("операнды должны быть числами от 1 до 10")
	}

	return operand1, operand2, operator
}

// WhatKindOfExpression если true - значит арабские, иначе false римские
func WhatKindOfExpression(input string) (bool, error) {
	const op = "services.WhatKindOfExpression"
	parts := strings.Split(input, " ")

	if len(parts) == 1 {
		return false, fmt.Errorf(
			"%s: %w",
			op,
			errors.New(("строка не является математической операцией")),
		)
	}
	if len(parts) != 3 {
		return false, fmt.Errorf(
			"%s: %w",
			op,
			errors.New(
				("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)"),
			),
		)
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
		return false, fmt.Errorf("%s: %w", op, errors.New("разные системы счисления"))
	}

	// Возвращаем true, если это арабские цифры, иначе false (римские цифры)
	return hasArabic, nil
}
