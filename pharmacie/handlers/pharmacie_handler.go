package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"pharmacie-api/pharmacie/models"
	"pharmacie-api/pharmacie/services"
	"strconv"
)

// @Summary Ajouter une pharmacie
// @Description Permet d'ajouter une pharmacie
// @Tags Pharmacie
// @Success 201 {string} string "Pharmacie ajoutée avec succès"
// @Failure 400 {string} string ""
// @Router /pharmacie [post]
func AjouterPharmacie(response http.ResponseWriter, request *http.Request) {

	var newPharmacie models.Pharmacie
	err := json.NewDecoder(request.Body).Decode(&newPharmacie)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	err = services.AjouterPharmacieServices(newPharmacie)
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

// @Summary Lister pharmacie
// @Description Permet de renvoyer la liste de toutes pharmacies disponible en base de donnée
// @Tags Pharmacie
// @Success 201 {string} string ""
// @Failure 400 {string} string ""
// @Router /pharmacie [get]
func ListerPharmacie(response http.ResponseWriter, request *http.Request) {

	liste, err := services.ListerPharmacieServices()
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

// @Summary Afficher une pharmacie
// @Description Permet de renvoyer une pharmacie spécifique
// @Tags Pharmacie
// @Success 201 {string} string ""
// @Failure 400 {string} string "Id invalides"
// @Router /pharmacie/{id} [get]
func AfficherPharmacie(response http.ResponseWriter, request *http.Request) {
	chemin := request.PathValue("id_pharmacie")

	id_pharmacie, err := strconv.Atoi(chemin)

	if err != nil {
		http.Error(response, "ID invalide", http.StatusBadRequest)
		return
	}

	afficherPharmacie, err := services.AfficherPharmacieServices(id_pharmacie)

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

// @Summary Modification d'une pharmacie
// @Description Permet de modifier une pharmacie spécifique
// @Tags Pharmacie
// @Success 201 {string} string "Pharmacie modifiée avec succès"
// @Failure 400 {string} string "Id invalides"
// @Router /pharmacie [put]
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
	err = services.ModifierPharmacieService(id_pharmacie, putPharmacie)
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

// @Summary Suppression d'une pharmacie
// @Description Permet de supprimer une pharmacie spécifique
// @Tags Pharmacie
// @Success 201 {string} string "pharmacie supprimée avec succès"
// @Failure 400 {string} string "Id invalides"
// @Router /pharmacie/{id} [delete]
func SupprimerPharmacie(response http.ResponseWriter, request *http.Request) {
	chemin := request.PathValue("id_pharmacie")
	id_pharmacie, err := strconv.Atoi(chemin)
	if err != nil {
		http.Error(response, "ID invalide", http.StatusBadRequest)
		return
	}

	err = services.SupprimerPharmacieServices(id_pharmacie)
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
