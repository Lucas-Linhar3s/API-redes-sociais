package controllers

import (
	"api/safety"
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/resposts"
	"encoding/json"
	"io/ioutil"
	"net/http"
)



func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resposts.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		resposts.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conect()
	if err != nil {
		resposts.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)
	userSaveDatabase, err := repository.SearchForEmail(user.Email)
	if err != nil {
		resposts.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = safety.VerifyPassword(userSaveDatabase.Password, user.Password); err != nil {
		resposts.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(userSaveDatabase.ID)
	if err != nil {
		resposts.Erro(w, http.StatusInternalServerError, err)
		return
	}
	w.Write([]byte(token))
}