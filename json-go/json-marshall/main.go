package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade int    `json:"idade"`
}

func main() {
	c := Cachorro{"Genghis", "Dalmata", 2}
	fmt.Println(c)

	cJson, err := json.Marshal(c)

	if err != nil {
		log.Fatalln("Erro ao serializar o cachorro para json")
	}
	

	fmt.Println(bytes.NewBuffer(cJson))
}
