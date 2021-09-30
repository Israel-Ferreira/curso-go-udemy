package controllers

import (
	"encoding/json"
	"fmt"
	"go-crud-api/models"
	"go-crud-api/repositories"
	"io/ioutil"
	"log"
	"net/http"
)

func InsertUser(rw http.ResponseWriter, r *http.Request) {
	var user models.Usuario
	
	

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(body, &user); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))

		return
	}

	db, err := repositories.ConnectDb("golang2", "golang2", "localhost", "33406", "devbook")

	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	usuarioRepo := repositories.NewUsuarioRepo(db)

	
	if err = usuarioRepo.CreateUser(user); err != nil {
		fmt.Println(err)
		rw.WriteHeader(500)
		return
	}



	fmt.Println("Conexão aberta com o DB")

	msg := map[string]string{
		"mesage": "usuario criado com sucesso",
	}

	msgInBytes, err := json.Marshal(msg)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))

		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write(msgInBytes)
}
