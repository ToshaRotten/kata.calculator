package Services

func Calculate(a int, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("деление на ноль")
		}
		return a / b
	default:
		panic("неверная операция")
	}
}
