package routes

import (
	"net/http"
	"pharmacie-api/middleware"
	"pharmacie-api/pharmacie/handlers"
)

func GestionRoutesPharmacie(limiter *middleware.RateLimiter) {

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
		"GET /pharmacie",
		middleware.RateLimit(limiter)(
			middleware.Logger(
				middleware.Auth(
					http.HandlerFunc(
						handlers.ListerPharmacie)),
			),
		),
	)

	// Route accessible à tous
	http.Handle(
		"GET /pharmacie/{id_pharmacie}",
		middleware.Logger(
			middleware.Auth(
				http.HandlerFunc(
					handlers.AfficherPharmacie)),
		),
	)

	// Administrateur et pharmacien
	http.Handle(
		"PUT /pharmacie/{id_pharmacie}",
		middleware.Logger(
			middleware.Auth(
				middleware.RequireRole("admin", "pharmacien")(
					http.HandlerFunc(
						handlers.ModifierPharmacie)),
			),
		),
	)

	// Uniquement pour l'administrateur
	http.Handle(
		"DELETE /pharmacie/{id_pharmacie}",
		middleware.Logger(
			middleware.Auth(
				middleware.RequireRole("admin")(
					http.HandlerFunc(
						handlers.SupprimerPharmacie)),
			),
		),
	)
	//===============================================================

	// Routes publiques

}
