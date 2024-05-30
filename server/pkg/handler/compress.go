package handler

import (
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func compressImage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	// Get the value of the "name" parameter from the request
	filename := r.URL.Query().Get("name")
	rateStr := r.URL.Query().Get("rate")
	rate, err := strconv.Atoi(rateStr)
	if err != nil {
		Locallog.Error(err)
		http.Error(w, "Wrong compress rate: "+rateStr, http.StatusBadRequest)
		return
	}
	name := strings.Split(filename, ".")

	// Open the image file
	file, err := os.Open("res/" + filename)
	if err != nil {
		Locallog.Error(err)
		http.Error(w, "File does not exist, or already processed and deleted\n", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Create a new compressed image file
	compressedName := fmt.Sprintf("%scompressed%d.jpeg", name[0], rate)
	compressedFile, err := os.Create("res/" + compressedName)
	if err != nil {
		Locallog.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	defer compressedFile.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		Locallog.Error(err)
		http.Error(w, "Wrong image format\n", http.StatusBadRequest)
		return
	}

	// Compress the image using your preferred compression algorithm
	var opts jpeg.Options
	opts.Quality = rate // Quality ranges from 1 to 100 inclusive, higher is better.

	err = jpeg.Encode(compressedFile, img, &opts)
	if err != nil {
		Locallog.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	download(w, r, compressedName)
	file.Close()
	compressedFile.Close()
	deleteFile(filename)
	deleteFile(compressedName)
}
