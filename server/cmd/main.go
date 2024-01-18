package main

import (
	"fmt"
	"log"
	"server"
	"server/pkg/handler"
	"server/pkg/repository"
	"server/pkg/service"
)

func main() {
	fmt.Println("hello")
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	server := new(server.Server)
	if err := server.Run("8080", *handler); err != nil {
		log.Fatalf("error occured %s", err.Error())
	}
}