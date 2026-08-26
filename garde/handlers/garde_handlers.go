package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"pharmacie-api/garde/models"
	"pharmacie-api/garde/services"
	"pharmacie-api/utils"
	"strconv"
)

// @Summary ajout d'une garde
// @Description Permet d'ajouter une garde
// @Tags Garde
// @Success 201 {string} string "Garde ajoutée aavec succès"
// @Failure 400 {string} string ""
// @Router /garde [post]
func AjouterGarde(response http.ResponseWriter, request *http.Request) {

	var newGarde models.Garde
	err := json.NewDecoder(request.Body).Decode(&newGarde)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.AjouterGardeService(newGarde)
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

// @Summary Liste de gardes
// @Description Permet de renvvoyer la liste toutes les gardes disponibles en base de données
// @Tags Garde
// @Success 201 {string} string ""
// @Failure 400 {string} string ""
// @Router /garde/list [post]
func ListerGardes(response http.ResponseWriter, request *http.Request) {
	nom_garde := request.URL.Query().Get("nom_garde")

	pagination := utils.GetPagination(request)
	filter := models.GardeFilter{
		Nom_garde: nom_garde,
	}
	liste, err := services.ListerGardeService(pagination, filter)
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

// @Summary Affichage de garde
// @Description Permet de renvoyer une garde spécifique
// @Tags Garde
// @Success 201 {string} string ""
// @Failure 401 {string} string "Id non valide"
// @Router /garde/{id_garde}/get [post]
func AfficherGarde(response http.ResponseWriter, request *http.Request) {

	chemin := request.PathValue("id_garde")
	id_garde, err := strconv.Atoi(chemin)
	if err != nil {
		http.Error(response, "ID non valide", http.StatusBadRequest)
		return
	}

	obtenirGarde, err := services.ObtenirGardeService(id_garde)
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

// @Summary Modification de garde
// @Description Permet de de modifier une garde spécifique
// @Tags Garde
// @Success 201 {string} string "Garde modifiée avec succès"
// @Failure 400 {string} string "Id invalides"
// @Router /garde/{id_garde}/put [post]
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

	err = services.ModifierGardeService(id_garde, putGarde)
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

// @Summary Supression de garde
// @Description Permet de supprimer une garde spécifique
// @Tags Garde
// @Success 201 {string} string "Garde supprimé avec succès"
// @Failure 400 {string} string "Id non valide"
// @Router /garde/{id_garde}/delete [post]
func SupprimerGarde(response http.ResponseWriter, request *http.Request) {
	chemin := request.PathValue("id_garde")
	id_garde, err := strconv.Atoi(chemin)
	if err != nil {
		http.Error(response, "ID non valide", http.StatusBadRequest)
		return
	}

	err = services.SupprimerGardeService(id_garde)
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
