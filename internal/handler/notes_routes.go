package handler

import (
	"errors"
	"net/http"

	"github.com/gitkoDev/KODE-test-task/helpers"
)

func (h *Handler) ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notes API for KODE v1"))
}

func (h *Handler) addNote(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id")
	userIdInt, ok := userId.(int)
	if !ok {
		helpers.RespondWithError(w, errors.New("error getting user id"), http.StatusBadRequest)
		return
	}

	var input, err = helpers.DecodeNotesJSON(r)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	// note_id, err := h.storage.AddNote(userIdInt, input)
	note_id, err := h.service.Notes.AddNote(userIdInt, input)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	helpers.RespondWithJSON(w, note_id, http.StatusCreated)
}

func (h *Handler) getNotes(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id")
	userIdInt, ok := userId.(int)
	if !ok {
		helpers.RespondWithError(w, errors.New("error getting user id"), http.StatusBadRequest)
		return
	}

	// notes, err := h.storage.GetAllNotes(userIdInt)
	notes, err := h.service.Notes.GetAllNotes(userIdInt)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusBadRequest)
	}

	helpers.RespondWithJSON(w, notes, http.StatusOK)
}
