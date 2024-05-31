package repository

import (
	"database/sql"
	"server/pkg/model"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	res, err := r.db.Exec("INSERT INTO users (email, login, password_hash) values ($1, $2, $3)", user.Email, user.Login, user.Password)
	if err != nil {
		Locallog.Error(err)
		return 0, err
	}
	// LastInsertId is not supported by this driver
	id, err := res.LastInsertId()
	if err != nil {
		Locallog.Error(err)
	}

	return int(id), nil
}

func (r *AuthPostgres) UserExists(user model.User) (bool, error) {
	return false, nil
}

func (r *AuthPostgres) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT email, login, password_hash FROM users WHERE email = $1", email).Scan(&user.Email, &user.Login, &user.Password)
	if err != nil {
		Locallog.Error(err)
		return model.User{}, err
	}

	return user, nil
}

func (r *AuthPostgres) GetUserByLogin(login string) (model.User, error) {
	var user model.User
	err := r.db.QueryRow("select email, login, password_hash from users where login=$1", login).Scan(&user.Email, &user.Login, &user.Password)
	if err != nil {
		Locallog.Error(err)
		return model.User{}, err
	}

	return user, nil
}
