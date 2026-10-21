// Package routes centralise l'enregistrement de tous les endpoints de l'API.
package routes

import (
	"net/http"

	"pharmacie-api/handlers"
	"pharmacie-api/middleware"

	httpSwagger "github.com/swaggo/http-swagger"
)

// GestionRoutes enregistre les routes métier et la documentation Swagger.
func GestionRoutes(
	ipLimiter *middleware.RateLimiter,
	userLimiter *middleware.RateLimiter,
) *http.ServeMux {
	mux := http.NewServeMux()
	/* Pharmacies : consultation publique, écriture réservée aux rôles autorisés.
	============================================================================*/

	mux.Handle(
		"POST /pharmacie",
		middleware.Logger(
			middleware.Auth(
				middleware.RequireRole("admin", "pharmacien")(
					http.HandlerFunc(handlers.AjouterPharmacie),
				),
			),
		),
	)
	mux.Handle(
		"POST /pharmacie/list",
		middleware.Logger(
			middleware.RateLimit(ipLimiter, nil)(
				http.HandlerFunc(handlers.ListerPharmacie),
			),
		),
	)
	mux.Handle(
		"POST /pharmacie/{id_pharmacie}/get",
		middleware.Logger(
			middleware.RateLimit(ipLimiter, nil)(
				http.HandlerFunc(handlers.AfficherPharmacie),
			),
		),
	)
	mux.Handle(
		"POST /pharmacie/{id_pharmacie}/put",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin", "pharmacien")(
						http.HandlerFunc(handlers.ModifierPharmacie),
					),
				),
			),
		),
	)
	mux.Handle(
		"POST /pharmacie/{id_pharmacie}/delete",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin")(
						http.HandlerFunc(handlers.SupprimerPharmacie),
					),
				),
			),
		),
	)

	/* Gardes : toutes les opérations nécessitent une authentification.
	=========================================================================*/
	mux.Handle(
		"POST /garde",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin", "pharmacien")(
						http.HandlerFunc(handlers.AjouterGarde),
					),
				),
			),
		),
	)
	mux.Handle(
		"POST /garde/list",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin", "pharmacien")(
						http.HandlerFunc(handlers.ListerGardes),
					),
				),
			),
		),
	)
	mux.Handle(
		"POST /garde/{id_garde}/get",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin", "pharmacien")(
						http.HandlerFunc(handlers.AfficherGarde),
					),
				),
			),
		),
	)
	mux.Handle(
		"POST /garde/{id_garde}/put",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin", "pharmacien")(
						http.HandlerFunc(handlers.ModifierGarde),
					),
				),
			),
		),
	)
	mux.Handle(
		"POST /garde/{id_garde}/delete",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin")(
						http.HandlerFunc(handlers.SupprimerGarde),
					),
				),
			),
		),
	)

	/* Pharmacies de garde : création/suppression admin, consultation publique.
	=====================================================================*/
	mux.Handle(
		"POST /pharmacie-garde",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin")(
						http.HandlerFunc(handlers.AjouterPharmacieGarde),
					),
				),
			),
		),
	)
	mux.Handle(
		"POST /pharmacie_garde/list",
		middleware.Logger(
			middleware.RateLimit(ipLimiter, nil)(
				http.HandlerFunc(handlers.ListerPharmacieGarde),
			),
		),
	)
	mux.Handle(
		"POST /pharmacie_garde/{id}/delete",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin")(
						http.HandlerFunc(handlers.SupprimerPharmacieGarde),
					),
				),
			),
		),
	)

	/* Utilisateurs : inscription et connexion publiques, administration protégée.
	==============================================================================*/
	mux.Handle(
		"POST /user",
		middleware.Logger(
			middleware.RateLimit(ipLimiter, nil)(
				http.HandlerFunc(handlers.AjouterUser),
			),
		),
	)
	mux.Handle(
		"POST /user/list",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin")(
						http.HandlerFunc(handlers.ListerUser),
					),
				),
			),
		),
	)
	mux.Handle(
		"POST /user/{id}/delete",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin")(
						http.HandlerFunc(handlers.SupprimerUser),
					),
				),
			),
		),
	)
	mux.Handle(
		"POST /user/{id}/role/patch",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin")(
						http.HandlerFunc(handlers.ModifierRole),
					),
				),
			),
		),
	)
	mux.Handle(
		"POST /user/{id}/get",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin", "pharmacien")(
						http.HandlerFunc(handlers.AfficherUser),
					),
				),
			),
		),
	)
	mux.Handle(
		"POST /login",
		middleware.Logger(
			middleware.RateLimit(ipLimiter, nil)(
				http.HandlerFunc(handlers.ConnexionUser),
			),
		),
	)

	// Documentation  de l'API.
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
	return mux
}
