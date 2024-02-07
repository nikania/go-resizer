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

	width, err := strconv.ParseUint(query.Get("width"), 10, 32)
	if err != nil {
		fmt.Println(err)
		return
	}
	height, err := strconv.ParseUint(query.Get("height"), 10, 32)
	if err != nil {
		fmt.Println(err)
		return
	}
	saveRatio := query.Get("save_ratio")
	name := strings.Split(filename, ".")

	file, err := os.Open("res/" + filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	var m image.Image
	if saveRatio != "" {
		m = resize.Thumbnail(uint(width), uint(height), img, resize.Bilinear)
	} else {
		m = resize.Resize(uint(width), uint(height), img, resize.Bilinear)
	}

	out, err := os.Create("res/" + name[0] + "resized." + format)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	switch format {
	case "png":
		png.Encode(out, m)
	case "jpeg":
		jpeg.Encode(out, m, nil)
	case "gif":
		gif.Encode(out, m, nil)
	default:
		png.Encode(out, m)
	}

}