package controllers

import (
	"encoding/json"
	"fmt"
	"go-crud-api/repositories"
	"log"
	"net/http"
)

func BuscarUsuarios(rw http.ResponseWriter, r *http.Request) {
	repo, err := repositories.CreateUserRepo("golang2", "golang2", "localhost", "33406", "devbook")

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	defer repo.CloseDbConnection()

	usuarios, err := repo.FindAll()

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Erro ao buscar os usuários no banco de dados"))
		return
	}

	rw.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(rw).Encode(usuarios); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Erro ao serializar o json"))
		return
	}

}

func BuscarUsuario(rw http.ResponseWriter, r *http.Request) {
	ID, err := getIdParam(r)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	repo, err := repositories.CreateUserRepo("golang2", "golang2", "localhost", "33406", "devbook")

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer repo.CloseDbConnection()

	usuario, err := repo.FindById(ID)

	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Erro: Usuario não encontrado"))
		return
	}

	rw.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(rw).Encode(usuario); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	

}

func InsertUser(rw http.ResponseWriter, r *http.Request) {
	// var user models.Usuario

	defer r.Body.Close()

	usuario, err := deserializeUserBody(r)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
		return
	}

	usuarioRepo, err := repositories.CreateUserRepo("golang2", "golang2", "localhost", "33406", "devbook")

	if err != nil {
		log.Fatalln(err)
	}

	defer usuarioRepo.CloseDbConnection()

	if err = usuarioRepo.CreateUser(usuario); err != nil {
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

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	rw.Write(msgInBytes)
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	ID, err := getIdParam(r)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := deserializeUserBody(r)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	db, err := repositories.CreateUserRepo("golang2", "golang2", "localhost", "33406", "devbook")

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer db.CloseDbConnection()

	if err = db.FindByIdAndUpdate(int(ID), body); err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	ID, err := getIdParam(r)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	repo, err := repositories.CreateUserRepo("golang2", "golang2", "localhost", "33406", "devbook")

	if err != nil {
		rw.WriteHeader(http.StatusAccepted)
	}

	if err = repo.DeleteById(int(ID)); err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
