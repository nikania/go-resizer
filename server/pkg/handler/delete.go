package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"
)


func deleteFile(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filename := query.Get("name")

	err := os.Remove("res/" + filename)
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteLoop() {
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