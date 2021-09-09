package main

import (
	"encoding/json"
	"log"
	"net/http"
)


type HelloMsg struct {
	Msg string `json:"msg"`
}


func main() {
	port := ":4000"


	http.HandleFunc("/home", func(rw http.ResponseWriter, r *http.Request) {
		msg := HelloMsg{"Olá Championship"}

		jsonMsg, _ := json.Marshal(msg)

		rw.Write(jsonMsg)
	})



	http.HandleFunc("/users", func(rw http.ResponseWriter, r *http.Request) {
		user := map[string]string{
			"name": "John Doe",
			"idade": "22",
		}

		jsonMsg, _ :=  json.Marshal(user)

		rw.Write(jsonMsg)
	})

	log.Fatal(http.ListenAndServe(port, nil))
}