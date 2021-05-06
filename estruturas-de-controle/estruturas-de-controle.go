package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estruturas de Controle")

	numero := -5

	if numero >= 15 {
		fmt.Println("O número é maior que 15")
	} else {
		fmt.Println("O número é menor  que 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que 0")
	} else if outroNumero == 0 {
		fmt.Println("O Número é  igual a 0")
	} else {
		fmt.Println("O número é menor que 0")
	}
}
