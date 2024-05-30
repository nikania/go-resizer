package repository

import (
	"database/sql"
	"server/logger"
	"server/pkg/model"
)

var Locallog logger.Logger

type Authorization interface {
	CreateUser(user model.User) (int, error)
	UserExists()
	AuthenticateUser()
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
