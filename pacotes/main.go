package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)


func main(){
	fmt.Println("Escrevendo do Arquivo Main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("code@example.com")

	fmt.Println(erro)
}