package controllers

import (
	"encoding/json"
	"go-crud-api/models"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func deserializeUserBody(r *http.Request) (models.Usuario, error) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return models.Usuario{}, err
	}

	var user models.Usuario

	if err = json.Unmarshal(body, &user); err != nil {
		return models.Usuario{}, err
	}

	return user, nil
}

func getIdParam(r *http.Request) (uint64, error) {
	params := mux.Vars(r)
	return strconv.ParseUint(params["id"], 10, 32)
}
