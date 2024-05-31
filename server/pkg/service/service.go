package service

import (
	"server/logger"
	"server/pkg/model"
	"server/pkg/repository"
)

var Locallog logger.Logger

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(credentials model.LoginCredentials) (string, error)
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
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
