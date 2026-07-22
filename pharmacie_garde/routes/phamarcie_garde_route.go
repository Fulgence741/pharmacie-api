package routes

import (
	"net/http"
	"pharmacie-api/pharmacie_garde/handlers"
)

func GestionRoutesPharmacieGarde() {
	http.HandleFunc("POST /pharmacie-garde", handlers.Ajouter)
	http.HandleFunc("GET /pharmacie/{id}/garde", handlers.ListerByGarde)
	http.HandleFunc("GET /garde/{id}/pharmacie", handlers.ListerByPharmacie)
	http.HandleFunc("DELETE /pharmacie-garde", handlers.Supprimer)
}
