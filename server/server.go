package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, you're requested %s\n", r.URL.Path)
}

func download(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filename := query.Get("name")
	// w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filename))
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, "res/"+filename)
}

func resizeImage(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filename := query.Get("name")
	width, err := strconv.ParseUint(query.Get("width"), 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	height, err := strconv.ParseUint(query.Get("height"), 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	saveRatio := query.Get("save_ratio")
	name := strings.Split(filename, ".")

	file, err := os.Open("res/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	var m image.Image
	if saveRatio != "" {
		m = resize.Thumbnail(uint(width), uint(height), img, resize.Bilinear)
	} else {
		m = resize.Resize(uint(width), uint(height), img, resize.Bilinear)
	}

	out, err := os.Create("res/" + name[0] + "resized." + format)
	if err != nil {
		log.Fatal(err)
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

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 2 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	fmt.Printf("MIME Header: %+v\n", handler.Header["Content-Type"])
	format := handler.Header["Content-Type"][0]

	allowed := []string{"image/jpeg", "image/png", "image/gif", "application/pdf"}
	if !Contains(allowed, format) {
		log.Print("not allowed")
		return
	}

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	id := uuid.New()
	var name string
	switch format {
	case "image/jpeg":
		name = fmt.Sprintf("img-%v.jpeg", id)
	case "image/png":
		name = fmt.Sprintf("img-%v.png", id)
	case "image/gif":
		name = fmt.Sprintf("img-%v.gif", id)
	case "application/pdf":
		name = fmt.Sprintf("%v.pdf", id)
	default:
		name = fmt.Sprintf("%v", id)
	}

	if _, err := os.Stat("res"); os.IsNotExist(err) {

		err = os.Mkdir("res", os.ModeDir)
		if err != nil {
			fmt.Println(err)
		}
	}
	tempFile, err := os.Create("res/" + name)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
	fmt.Fprintf(w, "name %s\n", name)
}

// can decode many formats (not so many) but only to allowed ones
func convertImage(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filename := query.Get("name")
	convertTo := query.Get("to")
	name := strings.Split(filename, ".")
	path := "res/" + filename

	allowed := []string{"image/png", "image/jpeg", "image/gif"}
	if !Contains(allowed, convertTo) {
		log.Fatal("Unsupported type to convert to")
		return
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var img image.Image

	img, _, err = image.Decode(file)
	if err != nil {
		log.Print(err)
		fmt.Fprintf(w, "Wrong image format\n")
		return
	}

	var convName string

	switch convertTo {
	case "image/png":
		convName = "res/" + name[0] + "conv.png"
		out, err := os.Create(convName)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		png.Encode(out, img)
	case "image/jpeg":
		convName = "res/" + name[0] + "conv.jpeg"
		out, err := os.Create(convName)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		jpeg.Encode(out, img, nil)
	case "image/gif":
		convName = "res/" + name[0] + "conv.gif"
		out, err := os.Create(convName)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		gif.Encode(out, img, nil)
	}

	fmt.Fprintf(w, "Successfully converted File\n")
	fmt.Fprintf(w, "name %s\n", convName)
}

func deleteFile(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filename := query.Get("name")

	err := os.Remove("res/" + filename)
	if err != nil {
		fmt.Println(err)
	}
}

func deleteLoop() {
	fmt.Println("entering dlete loop")
	for {
		fmt.Println("entering dlete loop: for")

		time.Sleep(time.Hour * 8)
		dir, err := os.ReadDir("res/")
		if err != nil {
			fmt.Println(err)
		}
		for i := 0; i < len(dir); i++ {
			fmt.Println(dir[i].Name())
			err := os.RemoveAll("res/" + dir[i].Name())
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println("deleted files")
	}
}

func main() {
	fmt.Println("hello")
	go deleteLoop()

	http.HandleFunc("/", handle)
	http.HandleFunc("/download", download)
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/resize", resizeImage)
	http.HandleFunc("/delete", deleteFile)
	http.HandleFunc("/convert", convertImage)

	http.ListenAndServe(":80", nil)

}
