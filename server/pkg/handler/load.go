package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"server/util"

	"github.com/google/uuid"
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


func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
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
	if !util.Contains(allowed, format) {
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