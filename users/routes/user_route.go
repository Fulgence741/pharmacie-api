package routes

import (
	"net/http"
	"pharmacie-api/middleware"
	"pharmacie-api/users/handlers"
)

func GestionRoutesUsers() {

	// Routes protégées par le middleware d'authentification
	//============================================================

	// Accessible à tous
	http.Handle(
		"POST /user",
		middleware.Logger(

			http.HandlerFunc(
				handlers.AjouterUser)),
	)

	// Administrateur uniquement
	http.Handle(
		"GET /user",
		middleware.Logger(

			middleware.Auth(
				middleware.RequireRole("admin")(
					http.HandlerFunc(
						handlers.ListerUser)),
			),
		),
	)

	// Administrateur uniquement

	http.Handle(

		"DELETE /user/{id}",
		middleware.Logger(
			middleware.Auth(
				middleware.RequireRole("admin")(
					http.HandlerFunc(
						handlers.SupprimerUser)),
			),
		),
	)

	// Addministrateu uquement
	http.Handle(
		"PATCH /user/{id}/role",
		middleware.Logger(
			middleware.Auth(
				middleware.RequireRole("admin")(
					http.HandlerFunc(
						handlers.ModifierRole)),
			),
		),
	)
	//============================================================

	// Routes publiques ========================================

	// Accessible à tous
	http.Handle(
		"POST /login",
		middleware.Logger(
			http.HandlerFunc(
				handlers.ConnexionUser),
		),
	)

}

//ici se termine toutes mes routes
