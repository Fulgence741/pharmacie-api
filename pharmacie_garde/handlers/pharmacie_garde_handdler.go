package handlers

import (
	"encoding/json"
	"net/http"
	"pharmacie-api/pharmacie_garde/models"
	"pharmacie-api/pharmacie_garde/services"
	"strconv"
)

// @Summary Ajouter pharmacie_garde
// @Description Permet de faire un ajouter de pharmacie et de garde en faisant un lien par Id
// @Tags Pharmacie-garde
// @Success 201 {string} string "Oppération réussie"
// @Failure 401 {string} string "Erreur lors de l'ajout"
// @Router /pharmacie-garde [post]
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

// @Summary Lister les pharmacies de garde
// @Description Permet de renvoyer la liste de toutes les pharmacies de gardes disponibles en base de donnée
// @Tags Pharmacie-garde
// @Success 201 {string} string ""
// @Failure 400 {string} string ""
// @Router /pharmacie-garde [get]
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

// @Summary suppression de pharmacie-garde
// @Description Permet de supprimer la liaison qui est entre une pharmacie et une garde sans toutefois supprimer la pharmacie ni la garde
// @Tags Pharmacie-garde
// @Success 201 {string} string "Oppération réussie"
// @Failure 400 {string} string "Id invalide"
// @Router /pharmacie-garde/{id} [delete]
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
