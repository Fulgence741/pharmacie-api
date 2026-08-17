package routes

import (
	"net/http"
	"pharmacie-api/garde/handlers"
	"pharmacie-api/middleware"
)

func GestionRoutesGarde(
	ipLimiter *middleware.RateLimiter,
	userLimiter *middleware.RateLimiter,
) {

	// Routes Protgées par middleware d'authentification
	//=============================================================

	// Administrateur et pharmacien
	http.Handle(
		"POST /garde",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin", "pharmacien")(
						http.HandlerFunc(
							handlers.AjouterGarde)),
				),
			),
		),
	)

	// AAdministrateur et pharmacien
	http.Handle(
		"POST /garde/list",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin", "pharmacien")(
						http.HandlerFunc(
							handlers.ListerGardes)),
				),
			),
		),
	)

	// Administrateur et pharmacien
	http.Handle(
		"POST /garde/{id_garde}/get",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin", "pharmacien")(
						http.HandlerFunc(
							handlers.AfficherGarde)),
				),
			),
		),
	)

	//Administrateur et pharmacien
	http.Handle(
		"POST /garde/{id_garde}/put",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin", "pharmacien")(
						http.HandlerFunc(
							handlers.ModifierGarde)),
				),
			),
		),
	)

	// Uniquement pour administrateur
	http.Handle(
		"POST /garde/{id_garde}/delete",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin")(
						http.HandlerFunc(
							handlers.SupprimerGarde)),
				),
			),
		),
	)
	//==============================================================

	// Routes publiques
}
