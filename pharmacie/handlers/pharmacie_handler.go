package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"pharmacie-api/pharmacie/models"
	"pharmacie-api/pharmacie/repositories"
	"strconv"
)

func AjouterPharmacie(response http.ResponseWriter, request *http.Request) {
	var newPharmacie models.Pharmacie
	err := json.NewDecoder(request.Body).Decode(&newPharmacie)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	err = repositories.AjouterPharmacieDB(newPharmacie)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(response).Encode("Pharmacie ajouté avec succès")
	if err != nil {
		http.Error(response, "Erreur du json", http.StatusInternalServerError)
		return
	}

}
func ListerPharmacie(response http.ResponseWriter, request *http.Request) {

	liste, err := repositories.ListerPharmacieDB()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode(liste)
	if err != nil {
		http.Error(response, "Erreur du json", http.StatusInternalServerError)
		return
	}
}
func AfficherPharmacie(response http.ResponseWriter, request *http.Request) {
	chemin := request.PathValue("id_pharmacie")

	id_pharmacie, err := strconv.Atoi(chemin)

	if err != nil {
		http.Error(response, "ID invalide", http.StatusBadRequest)
		return
	}

	afficherPharmacie, err := repositories.AfficherPharmacieDB(id_pharmacie)

	if err != nil {

		if err == sql.ErrNoRows {
			http.Error(response, "Pharmacie non trouvé", http.StatusNotFound)
			return
		}
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode(afficherPharmacie)

	if err != nil {
		http.Error(response, "Error du JSON", http.StatusInternalServerError)
		return
	}

}
func ModifierPharmacie(response http.ResponseWriter, request *http.Request) {

	chemin := request.PathValue("id_pharmacie")

	id_pharmacie, err := strconv.Atoi(chemin)

	if err != nil {
		http.Error(response, "ID invalide", http.StatusBadRequest)
		return
	}

	var putPharmacie models.Pharmacie

	err = json.NewDecoder(request.Body).Decode(&putPharmacie)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	err = repositories.ModifierPharmacieDB(id_pharmacie, putPharmacie)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode("Livre modifié avec succès")
	if err != nil {
		http.Error(response, "Erreur du JSON", http.StatusInternalServerError)
		return
	}

}
func SupprimerPharmacie(response http.ResponseWriter, request *http.Request) {
	chemin := request.PathValue("id_pharmacie")
	id_pharmacie, err := strconv.Atoi(chemin)
	if err != nil {
		http.Error(response, "ID invalide", http.StatusBadRequest)
		return
	}

	err = repositories.SupprimerPharmacieDB(id_pharmacie)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "appliction/json")
	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode("Pharmacie supprimé avec succès")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

}
