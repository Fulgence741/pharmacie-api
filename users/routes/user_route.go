package routes

import (
	"net/http"
	"pharmacie-api/users/handlers"
)

func GestionRoutesUsers() {
	http.HandleFunc("POST /user", handlers.AjouterUser)
	http.HandleFunc("GET /user", handlers.ListerUser)
	http.HandleFunc("POST /login", handlers.ConnexionUser)
}
