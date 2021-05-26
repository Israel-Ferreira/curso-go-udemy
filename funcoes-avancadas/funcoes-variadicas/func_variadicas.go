package main

import "fmt"

// Funções Variadicas
func soma(numeros ...int) int {
	var soma int

	for _, num := range numeros {
		soma += num
	}

	return soma
}

func escrever(texto string, numeros ...int) string {
	textoComNumeros := texto

	for _, num := range numeros {
		textoComNumeros += fmt.Sprintf(" %d", num)
	}

	return textoComNumeros
}

func main() {
	resultados := []int{soma(1, 2, 3, 4, 5, 6, 7), soma(909, 10)}
	fmt.Printf("%v \n", resultados)

	totalSoma := soma()
	fmt.Println(totalSoma)

	fmt.Println(escrever("Go é uma linguagem legal", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	fmt.Println(escrever("Olá Cambada", 1,2,3))
}
