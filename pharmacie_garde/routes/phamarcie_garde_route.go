package routes

import (
	"net/http"
	"pharmacie-api/middleware"
	"pharmacie-api/pharmacie_garde/handlers"
)

func GestionRoutesPharmacieGarde() {

	// Routes protégées par le middleware d'authentification
	//=========================================================

	//Administrateur uniquement
	http.Handle(
		"POST /pharmacie-garde",
		middleware.Logger(
			middleware.Auth(
				middleware.RequireRole("admin")(
					http.HandlerFunc(
						handlers.Ajouter)),
			),
		),
	)

	//Accessible à tous
	http.Handle(
		"GET /pharmacie-garde",
		middleware.Logger(
			middleware.Auth(
				http.HandlerFunc(
					handlers.Lister)),
		),
	)

	// Amdinistrateur uniquement
	http.Handle(
		"DELETE /pharmacie-garde/{id}",
		middleware.Logger(
			middleware.Auth(
				middleware.RequireRole("admin")(
					http.HandlerFunc(
						handlers.Supprimer)),
			),
		),
	)
	//=========================================================

	// Routes publiques

}
