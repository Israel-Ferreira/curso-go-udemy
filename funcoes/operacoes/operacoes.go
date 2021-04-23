package operacoes

import "errors"

func Soma(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Multi(a int, b int) int {
	return a * b
}

func Div(a int, b int) (float64, error) {
	num1, num2 := float64(a), float64(b)

	if num2 == 0 {
		return 0, errors.New("erro: Divisão por zero")
	}

	return float64(num1/num2), nil
}
