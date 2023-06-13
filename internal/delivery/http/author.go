package http_handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/giicoo/golang_CRUD/internal/models"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (h *Handler) addAuthor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Error("Body error: ", err)
		http.Error(w, ErrorOther.ErrorText, http.StatusBadRequest)
	}

	au := models.Author{}
	err = json.Unmarshal(body, &au)
	if err != nil {
		logrus.Error("Json error: ", err)
		http.Error(w, ErrorJson.ErrorText, http.StatusBadRequest)
	}

	author_id, err := h.services.AuthorServices.NewAuthor(au.Name)
	if err != nil {
		logrus.Error("Service error: ", err)
		http.Error(w, ErrorService.ErrorText, http.StatusInternalServerError)
	}

	logrus.Info(r.URL, http.StatusOK, " Status Code")
	w.WriteHeader(http.StatusOK)

	response := fmt.Sprintf("Author add successful with %v id", author_id)
	w.Write([]byte(response))
}

func (h *Handler) getAuthor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Error("Body error: ", err)
		http.Error(w, ErrorOther.ErrorText, http.StatusBadRequest)
		return
	}

	author := models.Author{}

	if err = json.Unmarshal(body, &author); err != nil {
		logrus.Error("Json error: ", err)
		http.Error(w, ErrorJson.ErrorText, http.StatusBadRequest)
		return
	}

	res, err := h.services.AuthorServices.GetAuthor(author)
	if err != nil {
		logrus.Error("Service error: ", err)
		http.Error(w, ErrorService.ErrorText, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		logrus.Error(err)
		http.Error(w, ErrorOther.ErrorText, http.StatusInternalServerError)
		return
	}

	logrus.Info(r.URL, http.StatusOK, " Status Code")
}

func (h *Handler) deleteAuthor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Error("Body error: ", err)
		http.Error(w, ErrorOther.ErrorText, http.StatusBadRequest)
		return
	}

	au := models.Author{}

	if err := json.Unmarshal(body, &au); err != nil {
		logrus.Error("Json error: ", err)
		http.Error(w, ErrorJson.ErrorText, http.StatusBadRequest)
		return
	}

	if err := h.services.AuthorServices.DeleteAuthor(au.Author_id); err != nil {
		logrus.Error("Service error: ", err)
		http.Error(w, ErrorService.ErrorText, http.StatusInternalServerError)
		return
	}

	res := fmt.Sprintf("Author with %v id delete successful", au.Author_id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}
