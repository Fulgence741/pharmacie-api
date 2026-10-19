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

// Objet qui rassemble toutes les informations et outils nécessaires pour appliquer une limitation de requêtes.
type RateLimiter struct {
	limit  int           // le nombre maximum de requête
	window time.Duration // la durée pendant laquelle on compte les requêtes
}

// Construction d'un rate limit
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:  limit,
		window: window,
	}
}

// Vérifier si une requete est autorisée

func (limiter *RateLimiter) Allow(
	ctx context.Context, // Context de la requete
	key string, // Clé qui identifie le client (ip ou id)
) (bool, error) {

	// permet d'incrémenter le compteur et de definir son expiration

	const scrypt = `
local current = redis.call("INCR", KEYS[1])
if current == 1 then 
				redis.call("PEXPIRE",
				 KEYS[1], ARGV[1])
				end 
				return current
`

	// KEYS = clé à incrémenter et ARVG = durée avant expiration

	rediskey := "rate_limit:" + key // la clé redis devient rate_limit: ip : + key

	windowMilliseconds := limiter.window.Milliseconds() // On convertis la durée de la fenêtre en miliseconde

	count, err := config.RedisClient.Eval(
		ctx,
		scrypt,
		[]string{rediskey},
		windowMilliseconds,
	).Int64()

	// Vérifier si redis a rencontré une erreur

	if err != nil {
		return false, fmt.Errorf(
			"erreur redis pendant le rate limitting : %w",
			err,
		)
	}
	// Vérifier si la limite est dépassée
	if count > int64(limiter.limit) {
		return false, nil
	}

	// Autoriser la requête
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

			ctx := request.Context()

			// 1. Limitation par utilisateur authentifié.

			userID, ok := ctx.Value(userIDKey).(int)

			if ok && userLimiter != nil {

				// On construit une clé unique pour cet ID
				userKey := fmt.Sprintf("user:%d", userID)

				// On vérifie si la limite n'est pas dépassée pour continuer les requêtes

				allowed, err := userLimiter.Allow(ctx, userKey)
				// vérifie une erreur interne du serveur

				if err != nil {
					http.Error(
						response,
						"Erreur interne du serveur",
						http.StatusInternalServerError,
					)
					return
				}

				// Refuser la requête si limite dépassée

				if !allowed {
					apierrors.WriteHTTP(
						response,
						apierrors.TooManyRequests(
							"Trop de tentatives",
						),
					)
					return
				}

				// Passer la requête au handler suivant

				next.ServeHTTP(response, request)
				return
			}

			// 2. Limitation par adresse IP.

			if ipLimiter == nil {
				next.ServeHTTP(response, request)
				return
			}

			ip, _, err := net.SplitHostPort(request.RemoteAddr)

			if err != nil {
				apierrors.WriteHTTP(
					response,
					apierrors.InvalidRequest(
						"Adresse IP invalide",
					),
				)
				return
			}
			// On constuit la clé redis associée à cette adresse ip
			ipKey := "ip:" + ip

			// On vérifie si la limitee n'est pas depassée
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

			// Si l'ip a depassé sa limite on refuse la requete

			if !allowed {
				apierrors.WriteHTTP(
					response,
					apierrors.TooManyRequests(
						"Trop de tentatives",
					),
				)
				return
			}

			// On transmet la requête au handler suivant
			next.ServeHTTP(response, request)
		})
	}
}
