package service

import (
	"fmt"
	"server/pkg/model"
	"server/pkg/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const (
	signingKey = "secret"
	tokenTTL   = time.Hour * 24
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

func (s *AuthService) GenerateToken(credentials model.LoginCredentials) (string, error) {
	userFromDB, err := s.repo.GetUserByLogin(credentials.Login)
	if err != nil {
		return "", err
	}
	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(credentials.Password))
	// nil means it is a match
	if err != nil {
		Locallog.Error(err)
		return "", err
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &BearerClaims{
		userFromDB.Id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		Locallog.Error(err)
		return "", err
	}

	return tokenString, nil
}

type BearerClaims struct {
	UserId int
	jwt.StandardClaims
}

func (s *AuthService) ParseToken(bearerToken string) (int, error) {
	token, err := jwt.ParseWithClaims(bearerToken, &BearerClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		Locallog.Error(err)
		return 0, err
	}
	claims, ok := token.Claims.(*BearerClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}
	return claims.UserId, nil
}

func (s *AuthService) CreateSession() {}

func (s *AuthService) DeleteSession() {}

func (s *AuthService) RefreshSession() {}

func (s *AuthService) CheckSession() {}

func (s *AuthService) GetSession() {}
