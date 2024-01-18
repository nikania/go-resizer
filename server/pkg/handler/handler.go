package handler

import ("net/http"
"server/pkg/service")

type Handler struct{
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
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