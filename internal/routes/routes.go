package routes

import (
	"database/sql"
	"go-pos-app/internal/handler"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {
	UserHandler := handler.NewUserHandler(db)

	ProductHandler := handler.NewProductHandler(db)

	transactionHandler := handler.NewTransactionHandler(db)

	http.HandleFunc("/users", UserHandler.HandleUsers)
	http.HandleFunc("/products", ProductHandler.HandleProduct)
	http.HandleFunc("/transaction", transactionHandler.CreateTransaction)
}