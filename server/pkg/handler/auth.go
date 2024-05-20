package handler

import "net/http"

func signUp(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	// query := r.URL.Query()

}

func signIn(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	// query := r.URL.Query()
}
