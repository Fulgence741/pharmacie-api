package middleware

import (
	"net/http"
)

func RequireRole(roleAutorise string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
			role := request.Context().Value("role")
			if role != roleAutorise {
				http.Error(response, "Accès interdit", http.StatusForbidden)
				return
			}
			next.ServeHTTP(response, request)
		})
	}
}
