package service

import (
	"server/logger"
	"server/pkg/repository"
)

var Locallog logger.Logger

type Authorization interface {
}

type Images interface {
}

type Documents interface{}

type Service struct {
	Authorization
	Images
	Documents
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
