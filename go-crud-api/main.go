package main

import (
	"fmt"
	"go-crud-api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("CRUD Go Lang")

	router := mux.NewRouter()

	router.HandleFunc("/usuarios", controllers.InsertUser).Methods(http.MethodPost)

	router.HandleFunc("/usuarios", controllers.BuscarUsuarios).Methods(http.MethodGet)

	router.HandleFunc("/usuarios/{id}", controllers.BuscarUsuario).Methods(http.MethodGet)

	

	fmt.Println("Escutando na porta :5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
