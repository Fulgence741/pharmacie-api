package routes

import (
	"net/http"
	"pharmacie-api/garde/handlers"
	"pharmacie-api/middleware"
)

func GestionRoutesGarde() {

	// Routes Protgées par middleware d'authentification
	//=============================================================

	// Administrateur et pharmacien
	http.Handle(
		"POST /garde",
		middleware.Logger(
			middleware.Auth(
				middleware.RequireRole("admin", "pharmacien")(
					http.HandlerFunc(
						handlers.AjouterGarde)),
			),
		),
	)

	// AAdministrateur et pharmacien
	http.Handle(
		"GET /garde",
		middleware.Logger(
			middleware.Auth(
				middleware.RequireRole("admin", "pharmacien")(
					http.HandlerFunc(
						handlers.ListerGardes)),
			),
		),
	)

	// Administrateur et pharmacien
	http.Handle(
		"GET /garde/{id_garde}",
		middleware.Logger(
			middleware.Auth(
				middleware.RequireRole("admin", "pharmacien")(
					http.HandlerFunc(
						handlers.AfficherGarde)),
			),
		),
	)

	//Administrateur et pharmacien
	http.Handle(
		"PUT /garde/{id_garde}",
		middleware.Logger(
			middleware.Auth(
				middleware.RequireRole("admin", "pharmacien")(
					http.HandlerFunc(
						handlers.ModifierGarde)),
			),
		),
	)

	// Uniquement pour administrateur
	http.Handle(
		"DELETE /garde/{id_garde}",
		middleware.Logger(
			middleware.Auth(
				middleware.RequireRole("admin")(
					http.HandlerFunc(
						handlers.SupprimerGarde)),
			),
		),
	)
	//==============================================================

	// Routes publiques
}
