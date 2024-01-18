package main

import (
	"fmt"
	"log"
	"server"
	"server/pkg/handler"
)

func main() {
	fmt.Println("hello")
	h := new(handler.Handler)

	server := new(server.Server)
	if err := server.Run("8080", *h); err != nil {
		log.Fatalf("error occured %s", err.Error())
	}
}