package routes

import (
	"database/sql"
	"go-pos-app/internal/handler"
	"go-pos-app/internal/middleware"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {

	// >> API Routes
	UserHandler := handler.NewUserHandler(db)

	ProductHandler := handler.NewProductHandler(db)

	transactionHandler := handler.NewTransactionHandler(db)

	authHandler := handler.NewAuthHandler(db)

	http.Handle("/users", 		middleware.JWTAuth(http.HandlerFunc(UserHandler.HandleUsers)))
	http.Handle("/products", 	middleware.JWTAuth(http.HandlerFunc(ProductHandler.HandleProduct)))

	// routes transactions
	http.Handle("/transaction", middleware.JWTAuth(http.HandlerFunc(transactionHandler.CreateTransaction)))
	http.Handle("/transaction/detail", 		 middleware.JWTAuth(http.HandlerFunc(transactionHandler.GetTransactionDetail)))
	http.Handle("/transaction/detail/all", 	 middleware.JWTAuth(http.HandlerFunc(transactionHandler.GetAllTransactions)))
	http.Handle("/transaction/report/sales", middleware.JWTAuth(http.HandlerFunc(transactionHandler.GetSalesReport)))
	http.Handle("/dashboard", 				 middleware.JWTAuth(http.HandlerFunc(transactionHandler.GetDashboardStats)))

	// login
	http.HandleFunc("/login", authHandler.Login)


	// >>  WEB Routes
	http.HandleFunc("/login-page", handler.RenderLogin)
	http.HandleFunc("/dashboard-page", handler.RenderDashboard)
	http.HandleFunc("/pos-page", handler.RenderPOS)
	http.HandleFunc("/product-page", handler.RenderProducts)


	// >> Static
	http.Handle("/static/" ,
	http.StripPrefix("/static/",
	http.FileServer(http.Dir("web/static"))))
}