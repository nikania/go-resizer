package handler

import (
	"fmt"
	"net/http"
	"os"
)

type Action int

const (
	Resize Action = iota
	Crop
	Convert
	Compress
)

func (h *Handler) chain(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	h.auth(w, r)
	if r.Method != "POST" {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
		return
	}
	filename := r.URL.Query().Get("name")
	file, err := os.Open("res/" + filename)
	if err != nil {
		Locallog.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "File does not exist, or already processed and deleted\n")
		return
	}
	defer file.Close()

	// decode body

	// perform chain of actions

	// encode response

	// fmt.Fprint(w, string(jsonResp))
}

const (
	AuthorizationHeader = "Authorization"
)

func (h *Handler) auth(w http.ResponseWriter, r *http.Request) {
	// check if the user is authenticated
	header := r.Header.Get(AuthorizationHeader)
	if header == "" {
		Locallog.Error("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// check if the token is valid
	headerToken := header[len("Bearer "):]
	userId, err := h.services.Authorization.ParseToken(headerToken)
	if err != nil {
		Locallog.Error(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	Locallog.Info("User id: ", userId)

}
