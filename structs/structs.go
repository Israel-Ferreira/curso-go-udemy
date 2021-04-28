package main

import "fmt"

type Usuario struct {
	nome  string
	idade uint8
	endereco Endereco
}

type Endereco struct {
	logradouro string
	numero     uint32
}

func main() {
	// Instanciando uma struct com constructor com argumentos

	endereco1 := Endereco{"Av. Paulista", 12345}
	user := Usuario{"Israel",21, endereco1}


	fmt.Println("Structs: ")

	fmt.Println(user)

	// Instanciando uma struct com constructor vázio
	user1 := Usuario{}

	// Valor zero por padrão em instâncias de structs inicializadas vazias
	fmt.Println(user1.nome)
	fmt.Println(user1.idade)

	endereco := Endereco{"Rua dos Bobos",32}

	user1.nome = "David"
	user1.idade = 32
	user.endereco = endereco

	fmt.Println(user1)

}
