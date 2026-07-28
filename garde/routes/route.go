package routes

import (
	"net/http"
	"pharmacie-api/garde/handlers"
	"pharmacie-api/middleware"
)

func GestionRoutesGarde() {

	http.Handle(
		"POST /garde",
		middleware.Auth(http.HandlerFunc(handlers.AjouterGarde)), // Protection des routes par le middleware
	)
	http.Handle(
		"GET /garde",
		middleware.Auth(http.HandlerFunc(handlers.ListerGardes)),
	)
	http.Handle(
		"GET /garde/{id_garde}",
		middleware.Auth(http.HandlerFunc(handlers.AfficherGarde)),
	)
	http.Handle(
		"PUT /garde/{id_garde}",
		middleware.Auth(http.HandlerFunc(handlers.ModifierGarde)),
	)
	http.Handle(
		"DELETE /garde/{id_garde}",
		middleware.Auth(http.HandlerFunc(handlers.SupprimerGarde)),
	)

}
