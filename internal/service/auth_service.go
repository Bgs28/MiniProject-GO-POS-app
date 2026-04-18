package service

import (
	"database/sql"
	"errors"
	"go-pos-app/internal/config"
	"go-pos-app/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	DB *sql.DB
	UserRepo *repository.UserRepository
}

func (s *AuthService) Login(username, password string) (string, error) {
	
	user, err := s.UserRepo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("Invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("Invalid username or password")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(config.JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}