package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"
)

type IPState struct {
	Requests int
	Start    time.Time
}

type RateLimiter struct {
	limit   int
	window  time.Duration
	clients map[string]IPState
	mu      sync.Mutex
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:   limit,
		window:  window,
		clients: make(map[string]IPState),
	}
}

func (request *RateLimiter) Allow(ip string) bool {
	request.mu.Lock()
	defer request.mu.Unlock()
	state, exists := request.clients[ip]
	if !exists {
		state = IPState{
			Requests: 1,
			Start:    time.Now(),
		}

		request.clients[ip] = state
		return true
	}
	now := time.Now()
	elapsed := now.Sub(state.Start)

	if elapsed >= request.window {
		state.Requests = 1
		state.Start = now
		request.clients[ip] = state

		return true
	}

	if state.Requests >= request.limit {
		return false
	}
	state.Requests++
	request.clients[ip] = state
	return true
}

func RateLimit(limiter *RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

			ip, _, err := net.SplitHostPort(request.RemoteAddr)

			if err != nil {
				http.Error(response, "Adresse IP invalide", http.StatusBadRequest)
				return
			}

			if !limiter.Allow(ip) {
				http.Error(response, "Too Many Requests", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(response, request)

		})
	}
}
