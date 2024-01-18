package handler

import "net/http"

type Handler struct{

}

func (h *Handler) Init() {
	go DeleteLoop()

	http.HandleFunc("/", handle)
	http.HandleFunc("/download", download)
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/resize", resizeImage)
	http.HandleFunc("/delete", deleteFile)
	http.HandleFunc("/convert", convertImage)
}