package handler

import (
	"net/http"
	"os"
	"time"
)

func deleteFile(w http.ResponseWriter, r *http.Request) {
	enableCors(&w, r)
	query := r.URL.Query()
	filename := query.Get("name")

	err := os.Remove("res/" + filename)
	if err != nil {
		Locallog.Error(err)
		return
	}
}

func DeleteLoop() {
	Locallog.Info("entering dlete loop")
	for {
		Locallog.Info("entering dlete loop: for")

		time.Sleep(time.Hour * 8)
		dir, err := os.ReadDir("res/")
		if err != nil {
			Locallog.Error(err)
		}
		for i := 0; i < len(dir); i++ {
			Locallog.Info("Deleting ", dir[i].Name())
			err := os.RemoveAll("res/" + dir[i].Name())
			if err != nil {
				Locallog.Error(err)
			}
		}
		Locallog.Info("deleted files")
	}
}
