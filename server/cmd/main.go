package main

import (
	"server"
	"server/logger"
	"server/pkg/handler"
	"server/pkg/repository"
	"server/pkg/service"
)

var locallog logger.Logger

func init() {
	locallog = logger.NewConsoleLogger(true, true)
	handler.Locallog = locallog
	repository.Locallog = locallog
	service.Locallog = locallog
}

func main() {
	conf, _ := ReadConfiguration()

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	server := new(server.Server)
	if err := server.Run(conf.Port, *handler); err != nil {
		locallog.Error("Error running server", err)
	}
}
