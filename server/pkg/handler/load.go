package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"server/util"

	"github.com/google/uuid"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, you're requested %s\n", r.URL.Path)
}

func download(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	query := r.URL.Query()
	filename := query.Get("name")
	Locallog.Info("Downloading file: ", filename)

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, "res/"+filename)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	Locallog.Info("Uploading file...")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	if err != nil {
		Locallog.Error("Error Retrieving the File", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	Locallog.Info("Uploaded File: ", handler.Filename)
	Locallog.Info("File Size: ", handler.Size)
	Locallog.Info("MIME Header: ", handler.Header)
	Locallog.Info("MIME Header: ", handler.Header["Content-Type"])
	format := handler.Header["Content-Type"][0]

	allowed := []string{"image/jpeg", "image/png", "image/gif", "application/pdf"}
	if !util.Contains(allowed, format) {
		// not allowed
		fmt.Fprintf(w, "File format not allowed: %s\n", format)
		w.WriteHeader(http.StatusBadRequest)
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
			Locallog.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	tempFile, err := os.Create("res/" + name)
	if err != nil {
		Locallog.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		Locallog.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	Locallog.Info("Successfully uploaded file: ", name)
	// return file name in json format
	jsonResp, err := json.Marshal(struct {
		Name string `json:"name"`
	}{
		Name: name,
	})
	if err != nil {
		Locallog.Error(err)
	}

	fmt.Fprint(w, string(jsonResp))
}
