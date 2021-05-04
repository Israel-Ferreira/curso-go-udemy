package main

import "fmt"

func main() {
	fmt.Println("Maps")
	usuario := map[string]string{
		"nome":      "Israel",
		"sobrenome": "Souza",
		"idade":     "21",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"usuario": {
			"nome":      "Israel",
			"sobrenome": "Souza",
		},

		"curso": {
			"nome":        "Análise e Desenvolvimento de Sistemas",
			"instituicao": "Fatec Jundiaí",
		},

		"endereco": {
			"logradouro": "Av. Jundiaí",
			"numero":     "1099",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "curso")

	fmt.Println(usuario2)

	usuario2["gostos"] = map[string]string{
		"comida":         "Lasanha",
		"bebida":         "Café",
		"estilo_musical": "Rock",
	}

	fmt.Println(usuario2)

}
