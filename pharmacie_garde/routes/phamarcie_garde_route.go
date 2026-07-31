package routes

import (
	"net/http"
	"pharmacie-api/middleware"
	"pharmacie-api/pharmacie_garde/handlers"
)

func GestionRoutesPharmacieGarde() {

	// Routes protégées par le middleware
	//=========================================================
	http.Handle(
		"POST /pharmacie-garde",
		middleware.Logger(
			middleware.Auth(
				http.HandlerFunc(
					handlers.Ajouter)),
		),
	)

	http.Handle(
		"GET /pharmacie-garde",
		middleware.Logger(
			middleware.Auth(
				http.HandlerFunc(
					handlers.Lister)),
		),
	)

	http.Handle(
		"DELETE /pharmacie-garde/{id}",
		middleware.Logger(
			middleware.Auth(
				http.HandlerFunc(
					handlers.Supprimer)),
		),
	)
	//=========================================================

	// Routes publiques

}
