package routes

import (
	"net/http"
	"pharmacie-api/middleware"
	"pharmacie-api/users/handlers"
)

func GestionRoutesUsers() {

	// Routes protégées par le middleware
	//============================================================
	/*http.Handle(
		"POST /user",
		middleware.Auth(
			http.HandlerFunc(
				handlers.AjouterUser)),
	) */

	http.Handle(
		"GET /user",
		middleware.Logger(

			middleware.Auth(
				http.HandlerFunc(
					handlers.ListerUser)),
		),
	)

	http.Handle(

		"DELETE /user/{id}",
		middleware.Logger(
			middleware.Auth(
				http.HandlerFunc(
					handlers.SupprimerUser)),
		),
	)
	//============================================================

	// Routes publiques ========================================

	http.Handle(
		"POST /user",
		middleware.Logger(
			http.HandlerFunc(
				handlers.AjouterUser),
		),
	)

	http.Handle(
		"POST /login",
		middleware.Logger(
			http.HandlerFunc(
				handlers.ConnexionUser),
		),
	)

}
