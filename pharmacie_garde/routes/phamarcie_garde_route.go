package routes

import (
	"net/http"
	"pharmacie-api/pharmacie_garde/handlers"
)

func GestionRoutesPharmacieGarde() {
	http.HandleFunc("POST /pharmacie-garde", handlers.Ajouter)
	http.HandleFunc("GET /pharmacie-garde", handlers.Lister)
	http.HandleFunc("DELETE /pharmacie-garde/{id}", handlers.Supprimer)
}
