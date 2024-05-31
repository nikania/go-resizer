package repository

import (
	"database/sql"
	"server/logger"
	"server/pkg/model"
)

var Locallog logger.Logger

type Authorization interface {
	CreateUser(user model.User) (int, error)
	UserExists(user model.User) (bool, error)
	GetUserByEmail(email string) (model.User, error)
	GetUserByLogin(login string) (model.User, error)
	GetUserById(id int) (model.User, error)
}

type Images interface {
}

type Documents interface{}

type Repository struct {
	Authorization
	Images
	Documents
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
