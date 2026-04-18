package middleware

import (
	"go-pos-app/internal/config"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Missing Token", http.StatusUnauthorized)
			return 
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{},error){
			return config.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return 
		}

		next.ServeHTTP(w,r)
	})
}