package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"pharmacie-api/garde/models"
	"pharmacie-api/garde/repositories"
	"strconv"
)

func AjouterGarde(response http.ResponseWriter, request *http.Request) {

	var newGarde models.Garde
	err := json.NewDecoder(request.Body).Decode(&newGarde)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	err = repositories.AjouterGardeDB(newGarde)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "appliction/json")
	response.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(response).Encode("Garde ajouté avec succès")
	if err != nil {
		http.Error(response, "Erreur du json", http.StatusInternalServerError)
		return
	}

}

func ListerGardes(response http.ResponseWriter, request *http.Request) {
	liste, err := repositories.ListerGardesDB()
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

func AfficherGarde(response http.ResponseWriter, request *http.Request) {

	chemin := request.PathValue("id_garde")
	id_garde, err := strconv.Atoi(chemin)
	if err != nil {
		http.Error(response, "ID non valide", http.StatusBadRequest)
		return
	}

	obtenirGarde, err := repositories.AfficherGarde(id_garde)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(response, "Garde non trouvée", http.StatusNotFound)
			return
		}
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "appliaction/json")
	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode(obtenirGarde)
	if err != nil {
		http.Error(response, "Erreur du JSON", http.StatusInternalServerError)
		return
	}

}

func ModifierGarde(response http.ResponseWriter, request *http.Request) {

	chemin := request.PathValue("id_garde")
	id_garde, err := strconv.Atoi(chemin)
	if err != nil {
		http.Error(response, "ID invalide", http.StatusBadRequest)
		return
	}
	var putGarde models.Garde
	err = json.NewDecoder(request.Body).Decode(&putGarde)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	err = repositories.ModifierGarde(id_garde, putGarde)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Cotent-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode("Garde modifié avec succès")
	if err != nil {
		http.Error(response, "Erreur du json", http.StatusInternalServerError)
		return
	}

}

func SupprimerGarde(response http.ResponseWriter, request *http.Request) {
	chemin := request.PathValue("id_garde")
	id_garde, err := strconv.Atoi(chemin)
	if err != nil {
		http.Error(response, "ID non valide", http.StatusBadRequest)
		return
	}

	err = repositories.SupprimerGarde(id_garde)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode("Garde supprimé avec succès")
	if err != nil {
		http.Error(response, "Erreur du JSON", http.StatusInternalServerError)
		return

	}
}
