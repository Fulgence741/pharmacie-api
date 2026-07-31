package middleware

import (
	"net/http"
	"pharmacie-api/auth"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		// Code exécuté avant le handler ================

		// Génération du token

		authorization := request.Header.Get("Authorization")

		if !strings.HasPrefix(authorization, "Bearer ") {
			http.Error(response, "Token invalide", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authorization, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return auth.SecretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(response, "Token non valide", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(response, request)

		// code exécuté après le handlers ================
	})
}
