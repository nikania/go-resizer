package repository

import (
	"database/sql"
	"server/logger"
)

var Locallog logger.Logger

type Authorization interface {
	CreateUser()
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
	return &Repository{}
}
