package handler

import (
	"net/http"
)

func enableCors(w *http.ResponseWriter, r *http.Request) {
	header := r.Header
	origin := header["Origin"]
	Locallog.Info("CORS origin: ", origin)
	if len(origin) != 0 && origin[0] == "http://localhost:5173" {
		Locallog.Info("Access-Control-Allow-Origin")
		(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	} else if len(origin) == 0 {

	} else {
		Locallog.Error(http.StatusForbidden, "Forbidden")
		(*w).WriteHeader(http.StatusForbidden)
	}
}
