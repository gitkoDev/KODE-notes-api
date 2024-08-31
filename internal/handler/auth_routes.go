package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gitkoDev/KODE-test-task/helpers"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var input, err = helpers.DecodeAuthJSON(r)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	if input.Name == "" || input.Password == "" {
		responseString := "please provide valid user name and password"
		helpers.RespondWithError(w, errors.New(responseString), http.StatusBadRequest)
		return
	}

	// id, err := h.storage.CreateUser(input)
	id, err := h.service.Auth.CreateUser(input)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	helpers.RespondWithJSON(w, map[string]int{"id": id}, http.StatusCreated)
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input, err = helpers.DecodeAuthJSON(r)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	if input.Name == "" || input.Password == "" {
		responseString := "please provide valid user name and password"
		helpers.RespondWithError(w, errors.New(responseString), http.StatusBadRequest)
		return
	}

	// id, err := h.storage.CheckForUserExistence(input.Name, input.Password)
	id, err := h.service.Auth.CheckForUserExistence(input.Name, input.Password)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	if id == -1 {
		responseString := fmt.Sprintf("user %s with this password doesn't exist in database\n", input.Name)
		helpers.RespondWithError(w, errors.New(responseString), http.StatusBadRequest)
		return
	}

	// token, err := h.storage.GenerateToken(input.Name, input.Password)
	token, err := h.service.Auth.GenerateToken(input.Name, input.Password)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	helpers.RespondWithJSON(w, map[string]string{"token": token}, http.StatusOK)

}
