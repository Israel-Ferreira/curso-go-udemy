package controllers

import (
	"encoding/json"
	"fmt"
	"go-crud-api/models"
	"go-crud-api/repositories"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func BuscarUsuarios(rw http.ResponseWriter, r *http.Request) {
	db, err := repositories.ConnectDb("golang2", "golang2", "localhost", "33406", "devbook")

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	defer db.Close()

	repo := repositories.NewUsuarioRepo(db)

	usuarios, err := repo.FindAll()

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Erro ao buscar os usuários no banco de dados"))
		return
	}

	if err = json.NewEncoder(rw).Encode(usuarios); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Erro ao serializar o json"))
		return
	}

}

func BuscarUsuario(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	db, err := repositories.ConnectDb("golang2", "golang2", "localhost", "33406", "devbook")

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	repo := repositories.NewUsuarioRepo(db)

	usuario, err := repo.FindById(ID)

	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Erro: Usuario não encontrado"))
		return
	}

	if err = json.NewEncoder(rw).Encode(usuario); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

}

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
