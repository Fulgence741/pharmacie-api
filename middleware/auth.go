package middleware

import (
	"context"
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

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return auth.SecretKey(), nil
		})

		if err != nil || !token.Valid {
			http.Error(response, "Token non valide", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(response, "Claims non valides", http.StatusUnauthorized)
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			http.Error(response, "Role absent", http.StatusUnauthorized)
			return
		}

		IDFloat, ok := claims["id"].(float64)
		if !ok {
			http.Error(response, "ID absent", http.StatusUnauthorized)
			return
		}
		userID := int(IDFloat)

		ctx := context.WithValue(
			request.Context(),
			"role",
			role,
		)

		ctx = context.WithValue(
			ctx,
			"id",
			userID,
		)
		request = request.WithContext(ctx)
		next.ServeHTTP(response, request)

		// code exécuté après le handlers ================
	})
}
