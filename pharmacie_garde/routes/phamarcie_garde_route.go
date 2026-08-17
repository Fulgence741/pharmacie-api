package routes

import (
	"net/http"
	"pharmacie-api/middleware"
	"pharmacie-api/pharmacie_garde/handlers"
)

func GestionRoutesPharmacieGarde(
	ipLimiter *middleware.RateLimiter,
	userLimiter *middleware.RateLimiter,
) {

	// Routes protégées par le middleware d'authentification
	//=========================================================

	//Administrateur uniquement
	http.Handle(
		"POST /pharmacie-garde",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin")(
						http.HandlerFunc(
							handlers.Ajouter)),
				),
			),
		),
	)

	//Accessible à tous
	http.Handle(
		"POST /pharmacie-garde/list",

		middleware.Logger(
			middleware.RateLimit(ipLimiter, nil)(
				http.HandlerFunc(
					handlers.Lister)),
		),
	)

	// Amdinistrateur uniquement
	http.Handle(
		"POST /pharmacie-garde/{id}/delete",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin")(
						http.HandlerFunc(
							handlers.Supprimer)),
				),
			),
		),
	)
	//=========================================================

	// Routes publiques

}
