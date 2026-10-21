package middleware

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"pharmacie-api/config"
	apierrors "pharmacie-api/errors"
	"time"
)

// RateLimiter est une structure qui gère la limitation du nombre de requêtes par utilisateur ou par IP.
type RateLimiter struct {
	limit  int           // le nombre maximum de requêtes autorisées dans la fenêtre de temps
	window time.Duration // la durée de la fenêtre de temps pour la limitation
}

// NewRateLimiter crée une nouvelle instance de RateLimiter avec les paramètres spécifiés.
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:  limit,  // le nombre maximum de requêtes autorisées dans la fenêtre de temps
		window: window, // la durée de la fenêtre de temps pour la limitation
	}
}

// Allow vérifie si une requête est autorisée en fonction de la clé (IP ou ID utilisateur) et de la limite définie.

func (limiter *RateLimiter) Allow(
	ctx context.Context, // le contexte de la requête HTTP
	key string, // la clé unique pour identifier l'utilisateur ou l'IP (par exemple, "user:123" ou "ip:192.168.1.1")
) (bool, error) {

	// Script Lua pour incrémenter la clé dans Redis et définir une expiration si c'est la première requête

	const scrypt = `
local current = redis.call("INCR", KEYS[1])
if current == 1 then 
				redis.call("PEXPIRE", 
				 KEYS[1], ARGV[1])
				end 
				return current
`

	// Construction de la clé Redis pour la limitation de débit

	rediskey := "rate_limit:" + key // On préfixe la clé pour éviter les collisions avec d'autres clés dans Redis

	windowMilliseconds := limiter.window.Milliseconds() // On convertit la durée de la fenêtre en millisecondes pour Redis

	count, err := config.RedisClient.Eval( // On exécute le script Lua dans Redis pour incrémenter la clé et obtenir le nombre actuel de requêtes
		ctx,                // le contexte de la requête HTTP
		scrypt,             // le script Lua à exécuter
		[]string{rediskey}, // la clé Redis à utiliser pour la limitation de débit
		windowMilliseconds, // la durée de la fenêtre en millisecondes pour Redis
	).Int64() // On récupère le résultat du script Lua en tant qu'entier

	// Vérifier si Redis a rencontré une erreur lors de l'exécution du script Lua

	if err != nil {
		return false, fmt.Errorf(
			"erreur redis pendant le rate limitting : %w",
			err,
		)
	}

	// Vérifier si la limite est dépassée

	if count > int64(limiter.limit) {
		return false, nil // La requête n'est pas autorisée car la limite est dépassée
	}

	// Autoriser la requête car elle est dans les limites définies

	return true, nil
}

// RateLimit applique une limitation par utilisateur ou par IP.

func RateLimit(
	ipLimiter *RateLimiter,
	userLimiter *RateLimiter,
) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(
			response http.ResponseWriter,
			request *http.Request,
		) {

			// Code exécuté avant le handler
			ctx := request.Context() // Récupération du contexte de la requête HTTP

			// 1. Limitation par ID utilisateur.

			// Récupération de l'identifiant de l'utilisateur depuis le contexte
			userID, ok := ctx.Value(userIDKey).(int)

			// Vérification si l'identifiant de l'utilisateur est présent et si le limiteur d'utilisateur est défini
			if ok && userLimiter != nil {

				// Construction de la clé Redis associée à l'utilisateur
				userKey := fmt.Sprintf("user:%d", userID)

				// Vérification si la limite n'est pas dépassée pour l'utilisateur

				allowed, err := userLimiter.Allow(ctx, userKey) // On vérifie si l'utilisateur est autorisé à faire la requête en utilisant la clé Redis associée à l'utilisateur
				// Vérifie une éventuelle erreur interne

				if err != nil {
					http.Error(
						response,
						"Erreur interne du serveur",
						http.StatusInternalServerError,
					)
					return
				}

				// Si l'utilisateur a dépassé sa limite, on refuse la requête

				if !allowed {
					apierrors.WriteHTTP(
						response,
						apierrors.TooManyRequests(
							"Trop de tentatives",
						),
					)
					return
				}
				// Si l'utilisateur est autorisé, on transmet la requête au handler suivant

				next.ServeHTTP(response, request)
				return
			}

			// 2. Limitation par adresse IP.

			if ipLimiter == nil {
				next.ServeHTTP(response, request)
				return
			}

			// Récupération de l'adresse IP du client à partir de l'adresse distante de la requête HTTP

			ip, _, err := net.SplitHostPort(request.RemoteAddr)
			// On utilise SplitHostPort pour séparer l'adresse IP et le port de l'adresse distante de la requête HTTP

			if err != nil {
				apierrors.WriteHTTP(
					response,
					apierrors.InvalidRequest(
						"Adresse IP invalide",
					),
				)
				return
			}
			// Construction de la clé Redis associée à l'adresse IP
			ipKey := "ip:" + ip

			// Vérification si la limite n'est pas dépassée pour l'adresse IP
			allowed, err := ipLimiter.Allow(ctx, ipKey)
			// Vérifie une éventuelle erreur interne
			if err != nil {
				http.Error(
					response,
					"Erreur interne du serveur",
					http.StatusInternalServerError,
				)
				return
			}

			// Si l'adresse IP a dépassé sa limite, on refuse la requête

			if !allowed {
				apierrors.WriteHTTP(
					response,
					apierrors.TooManyRequests(
						"Trop de tentatives",
					),
				)
				return
			}

			// Si l'adresse IP est autorisée, on transmet la requête au handler suivant
			next.ServeHTTP(response, request)
		})
	}
}
