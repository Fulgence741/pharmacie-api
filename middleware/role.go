package middleware

import (
	"net/http"
	apierrors "pharmacie-api/errors"
)

/*
	RequireRole est un middleware qui vérifie si l'utilisateur a le rôle requis pour accéder à la ressource.
	Il prend en paramètre une liste de rôles autorisés et vérifie si le rôle de l'utilisateur

(stocké dans le contexte de la requête) correspond à l'un des rôles autorisés. Si l'utilisateur n'a pas le rôle requis,
il renvoie une réponse HTTP 403 Forbidden.
*/
func RequireRole(roleAutorise ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

			// Code exécuté avant le handler

			// Récupération du rôle de l'utilisateur depuis le contexte de la requête
			role, ok := request.Context().Value(userRoleKey).(string)
			if !ok {
				apierrors.WriteHTTP(response, apierrors.Unauthorized("Rôle absent"))
				// Si le rôle n'est pas présent dans le contexte, on renvoie une réponse HTTP 401 Unauthorized
				return
			}

			// Vérification si le rôle de l'utilisateur correspond à l'un des rôles autorisés
			autorise := false // Initialisation de la variable autorise à false

			// Parcours de la liste des rôles autorisés pour vérifier si le rôle de l'utilisateur correspond à l'un d'eux
			for _, roleAutorise := range roleAutorise {
				if role == roleAutorise {
					// Si le rôle de l'utilisateur correspond à l'un des rôles autorisés, on met la variable autorise à true et on sort de la boucle
					autorise = true
					break
				}
			}

			// Si le rôle de l'utilisateur n'est pas autorisé, on renvoie une réponse HTTP 403 Forbidden
			if !autorise {
				apierrors.WriteHTTP(response, apierrors.Forbidden("Accès interdit"))
				return
			}

			// Passe la main au prochain middleware ou handler

			next.ServeHTTP(response, request)
		})
	}
}
