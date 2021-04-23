package main

import "fmt"

func main() {
	a, b := 30, 20

	// Operadores Aritméticos (+, -, *, /, %)

	soma := a + b
	subtracao := a - b
	multiplicacao := a * b
	divisao := float64(a) / float64(b)

	resto := float64(a % b)

	fmt.Printf("Soma: %d  \nSubstração: %d \nMultiplicação: %d \nDivisão: %f \nResto da Divisão: %f \n", soma, subtracao, multiplicacao, divisao, resto)

	// Operadores de Atribuição

	//Atribuição com a declaração do tipo da variável (=)
	var texto string = "The quick brown fox jumps over the lazy dog"

	// Atribuição com inferência de tipo (:=)
	texto2 := "The quick brown lynx jumps over the lazy cat"

	fmt.Println(texto, texto2)

	// OPERADORES RELACIONAIS (>, <, >=, <=, ==, !=)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a == b)

	fmt.Println(a != b)

	fmt.Println("-------------------")

	// OPERADORES LÓGICOS (&&, ||, !)

	fmt.Println(a != b && a >= b) // true
	fmt.Println(a == b && a >= b) // false

	fmt.Println(a != b || a <= b) // true

	fmt.Println("Negação ")
	fmt.Println(!true)  // false
	fmt.Println(!false) // true

	// OPERADORES Unários

	numero := 10
	numero++

	fmt.Println(numero)

	numero += 15

	fmt.Println(numero)

	var resultadoNumero string

	if numero > 5 {
		resultadoNumero = "Maior que 5"
	} else {
		resultadoNumero = "Menor que 5"
	}

	fmt.Println(resultadoNumero)
}
