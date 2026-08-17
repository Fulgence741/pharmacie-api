package middleware

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

type ClientState struct {
	Count int
	Start time.Time
}

type RateLimiter struct {
	limit   int
	window  time.Duration
	clients map[string]ClientState
	mu      sync.Mutex
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:   limit,
		window:  window,
		clients: make(map[string]ClientState),
	}
}

func (limiter *RateLimiter) Allow(key string) bool {
	limiter.mu.Lock()
	defer limiter.mu.Unlock()

	now := time.Now()

	state, exists := limiter.clients[key]

	// Première requête pour cette IP / cet utilisateur , en ce moment il est inconu
	if !exists {
		limiter.clients[key] = ClientState{
			Count: 1,
			Start: now,
		}
		return true
	}

	// La fenêtre est terminée
	if now.Sub(state.Start) >= limiter.window {
		limiter.clients[key] = ClientState{
			Count: 1,
			Start: now,
		}
		return true
	}

	// Limite atteinte
	if state.Count >= limiter.limit {
		return false
	}

	// Incrémentation du compteur
	state.Count++
	limiter.clients[key] = state

	return true
}

func (limiter *RateLimiter) Cleanup() {
	limiter.mu.Lock()
	defer limiter.mu.Unlock()

	now := time.Now()

	for key, state := range limiter.clients {
		if now.Sub(state.Start) >= limiter.window {
			delete(limiter.clients, key)
		}
	}
}

// Le middleware pour implémenter les deux, Limite par IP et limite par ID
func RateLimit(
	ipLimiter *RateLimiter,
	userLimiter *RateLimiter,
) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

			// ========================================
			// 1. Rate limit par IP
			// ========================================

			ip, _, err := net.SplitHostPort(request.RemoteAddr)

			if err != nil {
				http.Error(
					response,
					"Adresse IP invalide",
					http.StatusBadRequest,
				)
				return
			}

			if !ipLimiter.Allow("ip:" + ip) {
				http.Error(
					response,
					"Too Many Requests",
					http.StatusTooManyRequests,
				)
				return
			}

			// ========================================
			// 2. Rate limit par User ID
			// ========================================
			if userLimiter != nil {

				userID, ok := request.Context().Value("id").(int)

				if !ok {
					http.Error(
						response,
						"Utilisateur non identifié",
						http.StatusUnauthorized,
					)
					return
				}

				userKey := fmt.Sprintf("user:%d", userID)

				if !userLimiter.Allow(userKey) {
					http.Error(
						response,
						"Too Many Requests",
						http.StatusTooManyRequests,
					)
					return
				}
			}

			// On passe la main au handler

			next.ServeHTTP(response, request)
		})
	}
}
