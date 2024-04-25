package main

import (
	"bufio"
	"fmt"
	"kata.calculator/Services"
	"os"
)

func main() {
	fmt.Println("Hello Kata!!!")

	fmt.Print("Введите выражение в формате 'a + b': ")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Ошибка при чтении ввода: %v\n", err)
		return
	}

	whatKindOfExpression := Services.WhatKindOfExpression(input)

	if whatKindOfExpression {
		operand1, operand2, operator, err := Services.ValidateExpression(input)
		result, err := Services.Calculate(operand1, operand2, operator)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	} else {
		operandRoman1, operandRoman2, operator, err := Services.SeparateRoman(input)

		if err != nil {
			panic(err)
		}

		number1, _ := Services.RomanToArabic(operandRoman1)
		number2, _ := Services.RomanToArabic(operandRoman2)

		result, err := Services.Calculate(number1, number2, operator)
		if err != nil {
			panic(err)
		}
		if result > 1 {
			resultRoman, _ := Services.ArabicToRoman(result)
			fmt.Println(resultRoman)
		} else {
			fmt.Println("Выдача паники, так как в римской системе нет отрицательных чисел.")
		}
	}
}
