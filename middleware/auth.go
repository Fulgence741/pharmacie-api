package middleware

import (
	"context"
	"net/http"
	"pharmacie-api/auth"
	apierrors "pharmacie-api/errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type authContextKey string

const (
	userRoleKey authContextKey = "role"
	userIDKey   authContextKey = "id"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		// Code exécuté avant le handler ================

		// Génération du token

		authorization := request.Header.Get("Authorization")

		if !strings.HasPrefix(authorization, "Bearer ") {
			apierrors.WriteHTTP(response, apierrors.Unauthorized("Token invalide"))
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
			apierrors.WriteHTTP(response, apierrors.Unauthorized("Token non valide"))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			apierrors.WriteHTTP(response, apierrors.Unauthorized("Claims non valides"))
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			apierrors.WriteHTTP(response, apierrors.Unauthorized("Rôle absent"))
			return
		}

		IDFloat, ok := claims["id"].(float64)
		if !ok {
			apierrors.WriteHTTP(response, apierrors.Unauthorized("Identifiant absent"))
			return
		}
		userID := int(IDFloat)

		// Les clés typées évitent les collisions avec les valeurs ajoutées par d'autres middlewares.
		ctx := context.WithValue(
			request.Context(),
			userRoleKey,
			role,
		)

		ctx = context.WithValue(
			ctx,
			userIDKey,
			userID,
		)
		request = request.WithContext(ctx)
		next.ServeHTTP(response, request)

		// code exécuté après le handlers ================
	})
}
