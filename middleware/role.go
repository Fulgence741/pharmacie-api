package middleware

import (
	"net/http"
	apierrors "pharmacie-api/errors"
)

func RequireRole(roleAutorise ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

			role, ok := request.Context().Value(userRoleKey).(string)
			if !ok {
				apierrors.WriteHTTP(response, apierrors.Unauthorized("Rôle absent"))
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
				apierrors.WriteHTTP(response, apierrors.Forbidden("Accès interdit"))
				return
			}

			next.ServeHTTP(response, request)
		})
	}
}
