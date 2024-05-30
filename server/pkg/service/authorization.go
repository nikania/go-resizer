package service

import (
	"server/pkg/model"
	"server/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		Locallog.Error(err)
	}
	return string(hashedPassword)
}

func (s *AuthService) UserExists() {}

func (s *AuthService) AuthenticateUser(user model.User) {
	// // Comparing the password with the hash
	// err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	// fmt.Println(err) // nil means it is a match
}

func (s *AuthService) CreateSession() {}

func (s *AuthService) DeleteSession() {}

func (s *AuthService) RefreshSession() {}

func (s *AuthService) CheckSession() {}

func (s *AuthService) GetSession() {}
