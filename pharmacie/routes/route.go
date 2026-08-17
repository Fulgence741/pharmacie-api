package routes

import (
	"net/http"
	"pharmacie-api/middleware"
	"pharmacie-api/pharmacie/handlers"
)

func GestionRoutesPharmacie(
	ipLimiter *middleware.RateLimiter,
	userLimiter *middleware.RateLimiter,
) {

	// Routes protégées par le middleware d'authentification
	//================================================================

	// Administrateur et pharmacien
	http.Handle(
		"POST /pharmacie",
		middleware.Logger(
			middleware.Auth(
				middleware.RequireRole("admin", "pharmacien")(
					http.HandlerFunc(
						handlers.AjouterPharmacie)),
			),
		),
	)

	// Route accessible à tous
	http.Handle(
		"POST /pharmacie/list",

		middleware.Logger(
			middleware.RateLimit(ipLimiter, nil)(
				http.HandlerFunc(
					handlers.ListerPharmacie)),
		),
	)

	// Route accessible à tous
	http.Handle(
		"POST /pharmacie/{id_pharmacie}/get",
		middleware.Logger(
			middleware.RateLimit(ipLimiter, nil)(
				http.HandlerFunc(
					handlers.AfficherPharmacie)),
		),
	)

	// Administrateur et pharmacien
	http.Handle(
		"POST /pharmacie/{id_pharmacie}/put",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin", "pharmacien")(
						http.HandlerFunc(
							handlers.ModifierPharmacie)),
				),
			),
		),
	)

	// Uniquement pour l'administrateur
	http.Handle(
		"POST /pharmacie/{id_pharmacie}/delete",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin")(
						http.HandlerFunc(
							handlers.SupprimerPharmacie)),
				),
			),
		),
	)
	//===============================================================

	// Routes publiques

}
