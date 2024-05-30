package handler

import (
	"net/http"
	"server/logger"
	"server/pkg/service"
)

var Locallog logger.Logger

type Handler struct {
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
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/resize", resizeImage)
	http.HandleFunc("/convert", convertImage)
	http.HandleFunc("/crop", cropImage)
	http.HandleFunc("/compress", compressImage)
	http.HandleFunc("/auth", h.auth)
	http.HandleFunc("/signup", h.signUp)
	http.HandleFunc("/signin", h.signIn)
	http.HandleFunc("/signout", h.signOut)
}
