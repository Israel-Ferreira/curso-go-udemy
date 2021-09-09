package main

import (
	"encoding/json"
	"fmt"
	"log"
)


type Cachorro struct {
	Nome  string `json:"name"`
	Raca  string `json:"raca"`
	Idade int `json:"idade"`
}

type Gato struct {
	Nome  string `json:"name"`
	Raca  string `json:"raca"`
	Idade int `json:"idade"`
}





func main() {
	gato1 := `{"name":"Anakin","raca":"SRD","idade":3}`
	catioro1 := `{"name":"Ximira","raca":"Pinscher","idade":4}`


	var g Gato
	var c Cachorro


	if errDog := json.Unmarshal([]byte(catioro1), &c); errDog != nil {
		log.Fatalln("Erro ao deserializar o json")
	}

	if errCat := json.Unmarshal([]byte(gato1), &g); errCat != nil {
		log.Fatalln("Erro ao deserializar o json")
	}
	
	
	fmt.Println(g)
	fmt.Println(c)
}
