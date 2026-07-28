package routes

import (
	"net/http"
	"pharmacie-api/middleware"
	"pharmacie-api/users/handlers"
)

func GestionRoutesUsers() {

	http.Handle(
		"POST /user",
		middleware.Auth(http.HandlerFunc(handlers.AjouterUser)),
	)
	http.Handle(
		"GET /user",
		middleware.Auth(http.HandlerFunc(handlers.ListerUser)),
	)
	http.HandleFunc("POST /login", handlers.ConnexionUser)
}
