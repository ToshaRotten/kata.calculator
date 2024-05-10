package main

import (
	"bufio"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"kata.calculator/services"
	"kata.calculator/services/calculate"
)

func main() {
	fmt.Println("Введите выражение в формате 'a + b': ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		slog.Error(err.Error())
	}

	whatKindOfExpression, err := services.WhatKindOfExpression(input)
	if err != nil {
		slog.Error(err.Error())
	}

	if whatKindOfExpression {
		operand1, operand2, operator := services.ValidateExpression(input)
		result, err := calculate.Calculate(operand1, operand2, operator)
		if err != nil {
			slog.Error(err.Error())
		}
		fmt.Println(result)

	} else {
		operandRoman1, operandRoman2, operator, err := services.SeparateRoman(input)
		if err != nil {
			slog.Error(err.Error())
		}

		number1 := services.RomanToArabic(operandRoman1)
		number2 := services.RomanToArabic(operandRoman2)
		result, err := calculate.Calculate(number1, number2, operator)
		if err != nil {
			slog.Error(err.Error())
		}

		if result > 1 {
			resultRoman := services.ArabicToRoman(result)
			fmt.Println(resultRoman)
		} else {
			slog.Error(errors.New("no negatives in roman").Error())
		}
	}
}
