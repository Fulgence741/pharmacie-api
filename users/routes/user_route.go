package routes

import (
	"net/http"
	"pharmacie-api/middleware"
	"pharmacie-api/users/handlers"
)

func GestionRoutesUsers(
	ipLimiter *middleware.RateLimiter,
	userLimiter *middleware.RateLimiter,
) {

	// Routes protégées par le middleware d'authentification
	//============================================================

	// Route accessible  à tous
	http.Handle(
		"POST /user",
		middleware.Logger(
			middleware.RateLimit(ipLimiter, nil)(
				http.HandlerFunc(
					handlers.AjouterUser)),
		),
	)

	// Administrateur uniquement
	http.Handle(
		"POST /user/list",
		middleware.Logger(

			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin")(
						http.HandlerFunc(
							handlers.ListerUser)),
				),
			),
		),
	)

	// Administrateur uniquement

	http.Handle(

		"POST /user/{id}/delete",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin")(
						http.HandlerFunc(
							handlers.SupprimerUser)),
				),
			),
		),
	)

	// Addministrateu uquement
	http.Handle(
		"POST /user/{id}/role/patch",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin")(
						http.HandlerFunc(
							handlers.ModifierRole)),
				),
			),
		),
	)

	// Administrateur uniquement
	http.Handle(
		"POST /user/{id}/get",
		middleware.Logger(
			middleware.Auth(
				middleware.RateLimit(ipLimiter, userLimiter)(
					middleware.RequireRole("admin", "pharmacien")(
						http.HandlerFunc(
							handlers.AfficherUser)),
				),
			),
		),
	)
	//============================================================

	// Routes publiques ========================================

	// Accessible à tous
	http.Handle(
		"POST /login",

		middleware.Logger(
			middleware.RateLimit(ipLimiter, nil)(
				http.HandlerFunc(
					handlers.ConnexionUser),
			),
		),
	)

	/*
			// Routes pour modification de mot de passe à implémenter plus tard
			================================================================
		   	http.Handle(
		       "POST /user/change-password",
		       middleware.Logger(
		           middleware.Auth(
		               middleware.RateLimit(
		                   middleware.RequireRole(
		                       "user","admin","pharmacien",
		                       http.HandlerFunc(
		   						handlers.ChangerPassword)),
		                   ),
		               ),
		           ),
		       ),

		   }

	*/

	//ici se termine toutes mes routes

}
