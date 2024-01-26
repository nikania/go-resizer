package handler

import (
	"net/http"
	"fmt"
)

func enableCors(w *http.ResponseWriter, r *http.Request) {
	header:= r.Header;
	origin:= header["Origin"]
	fmt.Println(origin)
	if len(origin)!=0 && origin[0] == "http://localhost:5173" {
		fmt.Println("Access-Control-Allow-Origin")
		(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	} else if len(origin)==0  {

	} else {
		fmt.Println(403)
		(*w).WriteHeader(403)
	}
}