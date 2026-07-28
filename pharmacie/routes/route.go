package routes

import (
	"net/http"
	"pharmacie-api/middleware"
	"pharmacie-api/pharmacie/handlers"
)

func GestionRoutesPharmacie() {
	http.Handle(
		"POST /pharmacie",
		middleware.Auth(http.HandlerFunc(handlers.AjouterPharmacie)),
	)
	http.Handle(
		"GET /pharmacie",
		middleware.Auth(http.HandlerFunc(handlers.ListerPharmacie)),
	)

	http.Handle(
		"GET /pharmacie/{id_pharmacie}",
		middleware.Auth(http.HandlerFunc(handlers.AfficherPharmacie)),
	)
	http.Handle(
		"PUT /pharmacie/{id_pharmacie}",
		middleware.Auth(http.HandlerFunc(handlers.ModifierPharmacie)),
	)

	http.Handle(
		"DELETE /pharmacie/{id_pharmacie}",
		middleware.Auth(http.HandlerFunc(handlers.SupprimerPharmacie)),
	)

}
