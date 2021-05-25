package main

import (
	"fmt"
	"strings"
)

func contarVogais(palavra string) (qtde int) {
	palavra = strings.ToLower(palavra)
	vogais := []string{"a", "e", "i", "o", "u"}

	for _, letra := range palavra {
		for _, vogal := range vogais {
			if vogal == string(letra) {
				qtde++
			}
		}
	}

	return
}

func main() {
	i := 0

	/**
	* No Go, não há While, Until, Loop do; Tudo é feito com o For
	 */
	for i < 10 {

		i++
		fmt.Printf("Número: %d \n", i)
	}

	// For Classico
	fmt.Println("For Clássico")
	for j := 0; j < 10; j += 2 {
		fmt.Printf("Num: %d \n", j)
	}

	arr := []string{"Thiago", "Rubens", "Santiago"}

	for _, elem := range arr {
		fmt.Printf("Valor %s \n", elem)
	}

	qtdeVogais := contarVogais("Gato")
	fmt.Println("Quantidade vogais na palavra gato é: ", qtdeVogais)

	usuario := map[string]string{
		"nome":      "Luke",
		"sobrenome": "Skywalker",
	}

	fmt.Println(usuario)

	for key, value := range usuario {
		fmt.Printf("Chave %s, Valor %s", key, value)
	}

}
