package routes

import (
	"net/http"
	"pharmacie-api/pharmacie/handlers"
)

func GestionRoutes() {
	http.HandleFunc("POST /pharmacie", handlers.AjouterPharmacie)
	http.HandleFunc("GET /pharmacie", handlers.ListerPharmacie)
	http.HandleFunc("GET /pharmacie/{id_pharmacie}", handlers.AfficherPharmacie)
	http.HandleFunc("PUT /pharmacie/{id_pharmacie}", handlers.ModifierPharmacie)
	http.HandleFunc("DELETE /pharmacie/{id_pharmacie}", handlers.SupprimerPharmacie)
}
