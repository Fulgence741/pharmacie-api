package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"
)

type RateLimitClient struct {
	Count     int
	startedAt time.Time
}

var ipClients = make(map[string]*RateLimitClient)

var userClients = make(map[int]*RateLimitClient)

var rateLimitMutex sync.Mutex

const (
	publicLimit        = 10
	authenticatedLimit = 100
	rateLimitWindow    = time.Minute
)

func autoriserClient(client *RateLimitClient, limite int) bool {

	if time.Since(client.startedAt) >= rateLimitWindow {
		client.Count = 1
		client.startedAt = time.Now()

		return true
	}

	if client.Count >= limite {
		return false
	}

	client.Count++

	return true
}

func recupererIP(request *http.Request) string {
	host, _, err := net.SplitHostPort(request.RemoteAddr)
	if err != nil {
		return request.RemoteAddr
	}
	return host
}

func recupererID(request *http.Request) (int, bool) {
	id, ok := request.Context().Value("id").(int)
	if !ok {
		return 0, false
	}

	return id, true
}

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		id, authentifie := recupererID(request)

		if authentifie {
			rateLimitMutex.Lock()
			client, existe := userClients[id]

			if !existe {
				userClients[id] = &RateLimitClient{
					Count:     1,
					startedAt: time.Now(),
				}
				rateLimitMutex.Unlock()
				next.ServeHTTP(response, request)
				return
			}
			if autoriserClient(client, authenticatedLimit) {
				rateLimitMutex.Unlock()
				http.Error(response, "Trop de requêtes", http.StatusTooManyRequests)
				return
			}
			rateLimitMutex.Unlock()
			next.ServeHTTP(response, request)
			return
		}
	})
}
