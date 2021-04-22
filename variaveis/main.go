package main

import "fmt"

func main() {

	// Declaração de variável com tipo explicito
	var variavel1 string = "Variável 1"

	// Declaração da variável com tipo implicito, usando inferência de tipo
	variavel2 := "Variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// Declaração multipla  com tipo explicita
	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)


	// Declaração multipla com tipo implicito
	variavel5, variavel6 := "Variável 5", "Variável 6"

	fmt.Println(variavel3, variavel4)

	fmt.Println(variavel5, variavel6)

	const PI float64 = 3.14

	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)
}
