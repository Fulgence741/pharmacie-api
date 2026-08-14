package middleware

import (
	"net/http"
)

func RequireRole(roleAutorise ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

			role, ok := request.Context().Value("role").(string)
			if !ok {
				http.Error(response, "Role Absent", http.StatusUnauthorized)
				return
			}

			autorise := false

			for _, roleAutorise := range roleAutorise {
				if role == roleAutorise {
					autorise = true
					break
				}
			}
			if !autorise {
				http.Error(response, "Accès interdit", http.StatusForbidden)
				return
			}

			next.ServeHTTP(response, request)
		})
	}
}
