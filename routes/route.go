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
) {
	/* Pharmacies : consultation publique, écriture réservée aux rôles autorisés.
	============================================================================*/

	http.Handle(
		"POST /pharmacie",
		middleware.Logger(
			middleware.Auth(
				middleware.RequireRole("admin", "pharmacien")(
					http.HandlerFunc(handlers.AjouterPharmacie),
				),
			),
		),
	)
	http.Handle(
		"POST /pharmacie/list",
		middleware.Logger(
			middleware.RateLimit(ipLimiter, nil)(
				http.HandlerFunc(handlers.ListerPharmacie),
			),
		),
	)
	http.Handle(
		"POST /pharmacie/{id_pharmacie}/get",
		middleware.Logger(
			middleware.RateLimit(ipLimiter, nil)(
				http.HandlerFunc(handlers.AfficherPharmacie),
			),
		),
	)
	http.Handle(
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
	http.Handle(
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
	http.Handle(
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
	http.Handle(
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
	http.Handle(
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
	http.Handle(
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
	http.Handle(
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
	http.Handle(
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
	http.Handle(
		"POST /pharmacie-garde/list",
		middleware.Logger(
			middleware.RateLimit(ipLimiter, nil)(
				http.HandlerFunc(handlers.ListerPharmacieGarde),
			),
		),
	)
	http.Handle(
		"POST /pharmacie-garde/{id}/delete",
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
	http.Handle(
		"POST /user",
		middleware.Logger(
			middleware.RateLimit(ipLimiter, nil)(
				http.HandlerFunc(handlers.AjouterUser),
			),
		),
	)
	http.Handle(
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
	http.Handle(
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
	http.Handle(
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
	http.Handle(
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
	http.Handle(
		"POST /login",
		middleware.Logger(
			middleware.RateLimit(ipLimiter, nil)(
				http.HandlerFunc(handlers.ConnexionUser),
			),
		),
	)

	// Documentation  de l'API.
	http.Handle("/swagger/", httpSwagger.WrapHandler)
}
