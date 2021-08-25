package enderecos

import "strings"



func TipoDeEndereco(endereco string) (typeElem string) {
	tiposValidos := []string{"avenida", "rua",  "estrada", "rodovia"}

	for _, tipo := range tiposValidos {
		endereco = strings.ToLower(endereco);
		contains := Contains(endereco, tipo)

		if contains {
			typeElem = strings.Title(tipo) 
		}
	}


	if typeElem == "" {
		typeElem = "Tipo Inválido"
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