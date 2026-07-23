package handlers

import (
	"encoding/json"
	"net/http"
	"pharmacie-api/pharmacie_garde/models"
	"pharmacie-api/pharmacie_garde/services"
	"strconv"
)

func Ajouter(response http.ResponseWriter, request *http.Request) {
	var pharmacieDeGarde models.PharmacieGarde
	err := json.NewDecoder(request.Body).Decode(&pharmacieDeGarde)
	if err != nil {
		http.Error(response, "Données invalides", http.StatusBadRequest)
		return
	}

	err = services.AjouterService(pharmacieDeGarde)
	if err != nil {

		http.Error(response, "Erreur lors de l'ajout", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode("Oppération reussie !")
	if err != nil {
		http.Error(response, "Erreur du json", http.StatusInternalServerError)
		return
	}
}

func Lister(response http.ResponseWriter, request *http.Request) {
	pharmacieGarde, err := services.ListerService()

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode(pharmacieGarde)
	if err != nil {
		http.Error(response, "Erreur du JSON", http.StatusInternalServerError)
		return
	}

}

func Supprimer(response http.ResponseWriter, request *http.Request) {
	chemin := request.PathValue("id")
	id, err := strconv.Atoi(chemin)
	if err != nil {
		http.Error(response, "ID invalide", http.StatusBadRequest)
		return
	}

	err = services.SupprimerService(id)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode("Opération reussie!")
	if err != nil {
		http.Error(response, "Erreur du JSON", http.StatusInternalServerError)
		return
	}

}
