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
	err := r.db.QueryRow("INSERT INTO users (email, login, password_hash) values ($1, $2, $3) RETURNING id", user.Email, user.Login, user.Password).Scan(&user.Id)
	if err != nil {
		Locallog.Error(err)
		return 0, err
	}

	return int(user.Id), nil
}

func (r *AuthPostgres) UserExists(user model.User) (bool, error) {
	return false, nil
}

func (r *AuthPostgres) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT id, email, login, password_hash FROM users WHERE email = $1", email).Scan(&user.Id, &user.Email, &user.Login, &user.Password)
	if err != nil {
		Locallog.Error(err)
		return model.User{}, err
	}

	return user, nil
}

func (r *AuthPostgres) GetUserByLogin(login string) (model.User, error) {
	var user model.User
	err := r.db.QueryRow("select id, email, login, password_hash from users where login=$1", login).Scan(&user.Id, &user.Email, &user.Login, &user.Password)
	if err != nil {
		Locallog.Error(err)
		return model.User{}, err
	}

	return user, nil
}

func (r *AuthPostgres) GetUserById(id int) (model.User, error) {
	var user model.User
	err := r.db.QueryRow("select id, email, login, password_hash from users where id=$1", id).Scan(&user.Id, &user.Email, &user.Login, &user.Password)
	if err != nil {
		Locallog.Error(err)
		return model.User{}, err
	}

	return user, nil
}
