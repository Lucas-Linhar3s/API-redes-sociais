package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/resposts"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// FUNC FOR CRIATED USERS!
func CreateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		resposts.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
	}

	if erro = user.Prepare("cadastro"); erro != nil {
		resposts.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conect()
	if erro != nil {
		resposts.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)
	user.ID, erro = repository.Create(user)
	if erro != nil {
		resposts.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	resposts.JSON(w, http.StatusCreated, user)
}

// FUNC FOR SEARCH ALL USERS!
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Conect()
	if err != nil {
		resposts.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)
	users, err := repository.Search(nameOrNick)
	if err != nil {
		resposts.Erro(w, http.StatusInternalServerError, err)
		return
	}

	resposts.JSON(w, http.StatusOK, users)
}

// FUNC FOR SEARCH ONE USER!
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user_ID := mux.Vars(r)
	user_ID_conv, err := strconv.ParseUint(user_ID["userId"], 10, 64)
	if err != nil {
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
	user, err := repository.SearchForID(user_ID_conv)
	if err != nil {
		resposts.Erro(w, http.StatusInternalServerError, err)
		return
	}

	resposts.JSON(w, http.StatusOK, user)
}

// FUNC FOR UPDATE USERS!
func UptadeUsers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		resposts.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

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

	if err = user.Prepare("edicao"); err != nil {
		resposts.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Conect()
	if err != nil {
		resposts.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)
	if err = repository.Update(userID, user); err != nil {
		resposts.Erro(w, http.StatusInternalServerError, err)
		return
	}

	resposts.JSON(w, http.StatusNoContent, nil)

}

// FUNC FOR DELETE USERS!
func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
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
	if err = repository.DeleteUser(userID); err != nil {
		resposts.Erro(w, http.StatusInternalServerError, err)
		return
	}

	resposts.JSON(w, http.StatusNoContent, nil)
}
