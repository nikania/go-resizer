package handler

import (
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func cropImage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	query := r.URL.Query()
	filename := query.Get("name")
	name := strings.Split(filename, ".")
	x, err := strconv.ParseUint(query.Get("x"), 10, 32)
	if err != nil {
		Locallog.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Wrong starting x: %v", query.Get("x"))
		return
	}
	y, err := strconv.ParseUint(query.Get("y"), 10, 32)
	if err != nil {
		Locallog.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Wrong starting y: %v", query.Get("y"))
		return
	}
	width, err := strconv.ParseUint(query.Get("width"), 10, 32)
	if err != nil {
		Locallog.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Wrong width: %v", query.Get("width"))
		return
	}
	height, err := strconv.ParseUint(query.Get("height"), 10, 32)
	if err != nil {
		Locallog.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Wrong height: %v", query.Get("height"))
		return
	}
	Locallog.Info("Cropping image: ", filename)

	file, err := os.Open("res/" + filename)
	if err != nil {
		Locallog.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "File does not exist, or already processed and deleted\n")
		return
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Wrong image format: %v", err)
		return
	}

	// Define the rectangle for the portion of the image you want to keep
	// This example crops a 100x100 square from the top left corner of the image
	rect := image.Rect(int(x), int(y), int(width), int(height))

	// Create a new RGBA image of the required size
	croppedImg := image.NewRGBA(rect)

	// Draw the original image onto the cropped image
	draw.Draw(croppedImg, rect, img, rect.Min, draw.Src)

	// croppedImg now contains the cropped portion of the original image
	cropped := fmt.Sprintf("%scropped%d%d%d%d.%s", name[0], x, y, width, height, format)
	out, err := os.Create("res/" + cropped)
	if err != nil {
		Locallog.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer out.Close()

	switch format {
	case "png":
		err = png.Encode(out, croppedImg)
	case "jpeg":
		err = jpeg.Encode(out, croppedImg, nil)
	case "gif":
		err = gif.Encode(out, croppedImg, nil)
	default:
		err = png.Encode(out, croppedImg)
	}
	if err != nil {
		Locallog.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	download(w, r, cropped)
	file.Close()
	out.Close()
	deleteFile(filename)
	deleteFile(cropped)
}
