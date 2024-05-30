package model

type User struct {
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
