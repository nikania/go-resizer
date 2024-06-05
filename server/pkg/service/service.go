package service

import (
	"os"
	"server/logger"
	"server/pkg/model"
	"server/pkg/repository"
)

var Locallog logger.Logger

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(credentials model.LoginCredentials) (string, error)
	ParseToken(token string) (int, error)
}

type Images interface {
	Resize(file *os.File, width, height int, saveRatio bool) (*os.File, error)
	Crop(file *os.File, x, y, width, height int) (*os.File, error)
	Convert(file *os.File, format string) (*os.File, error)
	Compress(file *os.File, quality int) (*os.File, error)
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
		Images:        NewImagesService(),
		Documents:     nil,
	}
}
