package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ActionType string

const (
	Resize   ActionType = "resize"
	Crop     ActionType = "crop"
	Convert  ActionType = "convert"
	Compress ActionType = "compress"
)

type Params struct {
	Format              string `json:"format,omitempty"`
	Width               int    `json:"width,omitempty"`
	Height              int    `json:"height,omitempty"`
	MaintainAspectRatio bool   `json:"maintainAspectRatio,omitempty"`
	X                   int    `json:"x,omitempty"`
	Y                   int    `json:"y,omitempty"`
	Quality             int    `json:"quality,omitempty"`
}

type Action struct {
	Name   ActionType `json:"name"`
	Params Params     `json:"params"`
}

type ImageProcessingPipeline struct {
	Actions []Action `json:"actions"`
}

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
	var pipeline ImageProcessingPipeline
	err = json.NewDecoder(r.Body).Decode(&pipeline)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// perform chain of actions

	// Process the actions in the given order
	for _, action := range pipeline.Actions {
		switch action.Name {
		case Convert:
			h.services.Images.Convert(file, action.Params.Format)
			fmt.Printf("Converting to %s format\n", action.Params.Format)
		case Resize:
			h.services.Images.Resize(file, action.Params.Width, action.Params.Height, action.Params.MaintainAspectRatio)
			fmt.Printf("Resizing to %dx%d (maintain aspect ratio: %t)\n", action.Params.Width, action.Params.Height, action.Params.MaintainAspectRatio)
		case Crop:
			h.services.Images.Crop(file, action.Params.X, action.Params.Y, action.Params.Width, action.Params.Height)
			fmt.Printf("Cropping from (%d, %d) with dimensions %dx%d\n", action.Params.X, action.Params.Y, action.Params.Width, action.Params.Height)
		case Compress:
			h.services.Images.Compress(file, action.Params.Quality)
			fmt.Printf("Compressing with quality %d\n", action.Params.Quality)
		default:
			fmt.Printf("Unknown action: %v\n", action.Name)
		}
	}

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
