package handler

import (
	"database/sql"
	"go-pos-app/internal/repository"
	"go-pos-app/internal/service"
)

func NewAuthHandler(db *sql.DB) *AuthHandler{
	repo := &repository.UserRepository{
		DB: db,
	}

	service := &service.AuthService{
	   DB : db,
	   UserRepo : repo,
	}

	return &AuthHandler{
		Service: service,
	}
}