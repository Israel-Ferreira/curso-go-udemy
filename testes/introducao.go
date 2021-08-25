package main

import (
	"fmt"
	"introduction-tests/enderecos"
)

func main() {
	endereco := "Rodovia Dom Gabriel Paulino Couto, Km 50"

	tipo := enderecos.TipoDeEndereco(endereco)
	fmt.Println(tipo)
}
