package routes

import (
	"net/http"
	"pharmacie-api/garde/handlers"
	"pharmacie-api/middleware"
)

func GestionRoutesGarde() {

	// Routes Protgées par middleware d'authentification
	//=============================================================
	http.Handle(
		"POST /garde",
		middleware.Logger(
			middleware.Auth(
				http.HandlerFunc(
					handlers.AjouterGarde)),
		),
	)
	http.Handle(
		"GET /garde",
		middleware.Logger(
			middleware.Auth(
				http.HandlerFunc(
					handlers.ListerGardes)),
		),
	)
	http.Handle(
		"GET /garde/{id_garde}",
		middleware.Logger(
			middleware.Auth(
				http.HandlerFunc(
					handlers.AfficherGarde)),
		),
	)
	http.Handle(
		"PUT /garde/{id_garde}",
		middleware.Logger(
			middleware.Auth(
				http.HandlerFunc(
					handlers.ModifierGarde)),
		),
	)
	http.Handle(
		"DELETE /garde/{id_garde}",
		middleware.Logger(
			middleware.Auth(
				http.HandlerFunc(
					handlers.SupprimerGarde)),
		),
	)
	//==============================================================

	// Routes publiques
}
