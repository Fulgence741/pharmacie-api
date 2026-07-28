package routes

import (
	"net/http"
	"pharmacie-api/middleware"
	"pharmacie-api/pharmacie_garde/handlers"
)

func GestionRoutesPharmacieGarde() {
	http.Handle(
		"POST /pharmacie-garde",
		middleware.Auth(http.HandlerFunc(handlers.Ajouter)),
	)

	http.Handle(
		"GET /pharmacie-garde",
		middleware.Auth(http.HandlerFunc(handlers.Lister)),
	)

	http.Handle(
		"DELETE /pharmacie-garde/{id}",
		middleware.Auth(http.HandlerFunc(handlers.Supprimer)),
	)

}
