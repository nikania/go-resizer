package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/pkg/model"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	if r.Method != "POST" {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
		return
	}

	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		Locallog.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(user)

	_, err = h.services.Authorization.CreateUser(user)
	if err != nil {
		Locallog.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "\"id\": %v", id)
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	// query := r.URL.Query()
}

func (h *Handler) signOut(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	// query := r.URL.Query()
}

func (h *Handler) auth(w http.ResponseWriter, r *http.Request) {
	// check if the user is authenticated

}
