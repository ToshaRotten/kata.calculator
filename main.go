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
		panic(err)
	}

	whatKindOfExpression := Services.WhatKindOfExpression(input)

	if whatKindOfExpression {
		operand1, operand2, operator := Services.ValidateExpression(input)
		result := Services.Calculate(operand1, operand2, operator)

		fmt.Println(result)
	} else {
		operandRoman1, operandRoman2, operator := Services.SeparateRoman(input)

		number1 := Services.RomanToArabic(operandRoman1)
		number2 := Services.RomanToArabic(operandRoman2)
		result := Services.Calculate(number1, number2, operator)

		if result > 1 {
			resultRoman := Services.ArabicToRoman(result)
			fmt.Println(resultRoman)
		} else {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
		}
	}
}
