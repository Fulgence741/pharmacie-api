package routes

import (
	"net/http"
	"pharmacie-api/garde/handlers"
)

func GestionRoutesGarde() {

	http.HandleFunc("POST /garde", handlers.AjouterGarde)
	http.HandleFunc("GET /garde", handlers.ListerGardes)
	http.HandleFunc("GET /garde/{id_garde}", handlers.AfficherGarde)
	http.HandleFunc("PUT /garde/{id_garde}", handlers.ModifierGarde)
	http.HandleFunc("DELETE /garde/{id_garde}", handlers.SupprimerGarde)

}
