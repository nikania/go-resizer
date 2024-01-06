package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strings"

	// "strconv"
	"io"
	"net/http"

	// "net/url"

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
	aspectRatio := query.Get("aspect")
	name := strings.Split(filename, ".")

	file, err := os.Open("res/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	var m image.Image
	if aspectRatio != "" {

		m = resize.Thumbnail(100, 100, img, resize.Bilinear)
	} else {
		m = resize.Resize(200, 100, img, resize.Bilinear)

	}

	out, err := os.Create("res/" + name[0] + "SIZE.png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	png.Encode(out, m)

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

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	id := uuid.New()
	name := fmt.Sprintf("img-%v.png", id)
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

func deleteFile(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filename := query.Get("name")

	err := os.Remove("res/" + filename)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	fmt.Println("hello")
	http.HandleFunc("/", handle)
	http.HandleFunc("/download", download)
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/resize", resizeImage)
	http.HandleFunc("/delete", deleteFile)

	http.ListenAndServe(":80", nil)
}
