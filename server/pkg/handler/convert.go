package handler

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"server/util"
	"strings"
)

// can decode many formats (not so many) but only to allowed ones
func convertImage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	query := r.URL.Query()
	filename := query.Get("name")
	convertTo := query.Get("to")
	name := strings.Split(filename, ".")
	path := "res/" + filename
	Locallog.Info("Converting image: ", filename, " to ", convertTo, " format")

	allowed := []string{"image/png", "image/jpeg", "image/gif"}
	if !util.Contains(allowed, convertTo) {
		Locallog.Error("Unsupported type to convert to")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, err := os.Open(path)
	if err != nil {
		Locallog.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "File does not exist, or already processed and deleted\n")
		return
	}
	defer file.Close()

	var img image.Image

	img, _, err = image.Decode(file)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Wrong image format: %v", err)
		return
	}

	var convName string

	switch convertTo {
	case "image/png":
		convName = "res/" + name[0] + "conv.png"
		out, err := os.Create(convName)
		if err != nil {
			Locallog.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer out.Close()

		png.Encode(out, img)
	case "image/jpeg":
		convName = "res/" + name[0] + "conv.jpeg"
		out, err := os.Create(convName)
		if err != nil {
			Locallog.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer out.Close()

		jpeg.Encode(out, img, nil)
	case "image/gif":
		convName = "res/" + name[0] + "conv.gif"
		out, err := os.Create(convName)
		if err != nil {
			Locallog.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer out.Close()

		gif.Encode(out, img, nil)
	}

	download(w, r, convName)
}
