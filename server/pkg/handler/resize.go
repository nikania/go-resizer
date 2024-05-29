package handler

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
)

func resizeImage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	query := r.URL.Query()
	filename := query.Get("name")
	Locallog.Info("Resizing image: ", filename)

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
	saveRatio := query.Get("save_ratio")
	name := strings.Split(filename, ".")

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
		Locallog.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var m image.Image
	if saveRatio != "" {
		m = resize.Thumbnail(uint(width), uint(height), img, resize.Bilinear)
	} else {
		m = resize.Resize(uint(width), uint(height), img, resize.Bilinear)
	}

	resized := name[0] + "resized." + format
	out, err := os.Create("res/" + resized)
	if err != nil {
		Locallog.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer out.Close()

	switch format {
	case "png":
		err = png.Encode(out, m)
	case "jpeg":
		err = jpeg.Encode(out, m, nil)
	case "gif":
		err = gif.Encode(out, m, nil)
	default:
		err = png.Encode(out, m)
	}
	if err != nil {
		Locallog.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	download(w, r, resized)
	file.Close()
	out.Close()
	deleteFile(filename)
	deleteFile(resized)
}
