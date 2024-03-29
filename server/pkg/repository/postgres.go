package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func NewPostgresDb(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	Locallog.Info("Connected to db")
	return db, nil
}
