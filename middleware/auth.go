package middleware

import (
	"context"
	"net/http"
	"pharmacie-api/auth"
	apierrors "pharmacie-api/errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// authContextKey est un type de clé pour le contexte afin d'éviter les collisions avec d'autres valeurs de contexte.
type authContextKey string

// Les clés pour le rôle et l'identifiant de l'utilisateur dans le contexte
const (
	userRoleKey authContextKey = "role"
	userIDKey   authContextKey = "id"
)

// Auth est un middleware qui vérifie la présence et la validité d'un token JWT dans l'en-tête Authorization de la requête HTTP.
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		// Code exécuté avant le handler

		// Génération du token

		// Récupération du token JWT depuis l'en-tête Authorization
		authorization := request.Header.Get("Authorization")

		// Vérification de la présence du token
		if authorization == "" {
			apierrors.WriteHTTP(response, apierrors.Unauthorized("Token manquant"))
			return
		}

		// Vérification du format du token (doit commencer par "Bearer ")
		if !strings.HasPrefix(authorization, "Bearer ") {
			apierrors.WriteHTTP(response, apierrors.Unauthorized("Token invalide"))
			return
		}

		// Extraction du token JWT en supprimant le préfixe "Bearer "

		tokenString := strings.TrimPrefix(authorization, "Bearer ")

		// Vérification de la validité du token JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			// Retourne la clé secrète pour la vérification de la signature du token
			return auth.SecretKey(), nil
		})

		// Vérification de la validité du token et des claims

		if err != nil || !token.Valid {
			apierrors.WriteHTTP(response, apierrors.Unauthorized("Token non valide"))
			return
		}

		// Extraction des claims du token JWT

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			apierrors.WriteHTTP(response, apierrors.Unauthorized("Claims non valides"))
			return
		}

		// Extraction du rôle et de l'identifiant de l'utilisateur à partir des claims

		role, ok := claims["role"].(string)
		if !ok {
			apierrors.WriteHTTP(response, apierrors.Unauthorized("Rôle absent"))
			return
		}
		// Extraction de l'identifiant de l'utilisateur à partir des claims
		IDFloat, ok := claims["id"].(float64)
		if !ok {
			apierrors.WriteHTTP(response, apierrors.Unauthorized("Identifiant absent"))
			return
		}

		// Conversion de l'identifiant de l'utilisateur en entier
		userID := int(IDFloat)

		// Ajout du rôle et de l'identifiant de l'utilisateur au contexte de la requête
		ctx := context.WithValue(
			request.Context(), // Récupération du contexte de la requête HTTP
			userRoleKey,       // Clé pour le rôle de l'utilisateur dans le contexte
			role,              // Valeur du rôle de l'utilisateur extrait des claims du token JWT
		)
		// Ajout de l'identifiant de l'utilisateur au contexte de la requête

		ctx = context.WithValue(
			ctx,       // Contexte actuel contenant le rôle de l'utilisateur
			userIDKey, // Clé pour l'identifiant de l'utilisateur dans le contexte
			userID,    // Valeur de l'identifiant de l'utilisateur extrait des claims du token JWT
		)
		// Mise à jour de la requête avec le nouveau contexte contenant le rôle et l'identifiant de l'utilisateur
		request = request.WithContext(ctx)
		next.ServeHTTP(response, request)

		// Code exécuté après le handler
	})
}
