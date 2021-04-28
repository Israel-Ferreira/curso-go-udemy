package main

import (
	"fmt"
)

type Pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint32
}

type Estudante struct {
	Pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Println("Herança #SQN")

	p1 := Pessoa{"Israel", "Souza", 21, 178}
	fmt.Println(p1)

	e1 := Estudante{p1, "Análise e Desenvolvimento de Sistemas", "Fatec"}
	fmt.Println(e1.altura)



}
