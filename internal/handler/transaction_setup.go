package handler

import (
	"database/sql"
	"go-pos-app/internal/repository"
	"go-pos-app/internal/service"
)

func NewTransactionHandler(db *sql.DB) *TransactionHandler{
	repo := &repository.TransactionRepository{
		DB: db,
	}

	service := &service.TransactionService{
		DB: db,
		Repo: repo,
	}

	return &TransactionHandler{
		Service: service,
	}
}