package enderecos

import "strings"



func TipoDeEndereco(endereco string) (typeElem string) {
	tiposValidos := []string{"Avenida", "Rua",  "Estrada", "Rodovia"}

	for _, tipo := range tiposValidos {
		contains := Contains(endereco, tipo)

		if contains {
			typeElem = tipo
		}
	}


	return 
}






func Contains(quote, strFind string) bool {
	quoteWords := strings.Split(quote, " ")


	for _, elem := range quoteWords {
		if elem == strFind {
			return true
		}
	}


	return false
}