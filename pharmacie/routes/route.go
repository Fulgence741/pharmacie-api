package routes

import (
	"net/http"
	"pharmacie-api/middleware"
	"pharmacie-api/pharmacie/handlers"
)

func GestionRoutesPharmacie() {

	// Routes protégées par le middleware
	//================================================================
	http.Handle(
		"POST /pharmacie",
		middleware.Logger(
			middleware.Auth(
				http.HandlerFunc(
					handlers.AjouterPharmacie)),
		),
	)
	http.Handle(
		"GET /pharmacie",
		middleware.Logger(
			middleware.Auth(
				http.HandlerFunc(
					handlers.ListerPharmacie)),
		),
	)

	http.Handle(
		"GET /pharmacie/{id_pharmacie}",
		middleware.Logger(
			middleware.Auth(
				http.HandlerFunc(
					handlers.AfficherPharmacie)),
		),
	)
	http.Handle(
		"PUT /pharmacie/{id_pharmacie}",
		middleware.Logger(
			middleware.Auth(
				http.HandlerFunc(
					handlers.ModifierPharmacie)),
		),
	)

	http.Handle(
		"DELETE /pharmacie/{id_pharmacie}",
		middleware.Logger(
			middleware.Auth(
				http.HandlerFunc(
					handlers.SupprimerPharmacie)),
		),
	)
	//===============================================================

	// Routes publiques

}
