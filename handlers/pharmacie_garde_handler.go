package handlers

import (
	"net/http"
	"pharmacie-api/dto"
	apierrors "pharmacie-api/errors"
	"pharmacie-api/models"
	"pharmacie-api/services"
	"pharmacie-api/utils"
)

// @Summary Ajouter pharmacie_garde
// @Description Permet de lier une pharmacie à une garde par leurs identifiants
// @Tags Pharmacie-garde
// @Success 200 {string} string "Opération réussie"
// @Failure 400 {string} string "Données invalides"
// @Router /pharmacie-garde [post]
func AjouterPharmacieGarde(response http.ResponseWriter, request *http.Request) {
	var input dto.PharmacieGardeRequest
	if err := decodeJSON(request, &input); err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	lien := models.PharmacieGarde{
		PharmacieID: input.PharmacieID,
		GardeID:     input.GardeID,
	}
	if err := services.AjouterPharmacieGardeService(request.Context(), lien); err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	writeJSON(response, http.StatusOK, "Oppération reussie !")
}

// @Summary Lister les pharmacies de garde
// @Description Renvoie les pharmacies de garde disponibles en base de données
// @Tags Pharmacie-garde
// @Success 200 {array} models.PharmacieGardeA
// @Failure 500 {string} string "Erreur interne"
// @Router /pharmacie-garde/list [post]
func ListerPharmacieGarde(response http.ResponseWriter, request *http.Request) {
	filter := models.PharmacieGardeFliter{Nom: request.URL.Query().Get("nom")}
	result, err := services.ListerPharmacieGardeService(request.Context(), utils.GetPagination(request), filter)
	if err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	writeJSON(response, http.StatusOK, pharmacieGardeResponses(result))
}

// @Summary Suppression de pharmacie-garde
// @Description Supprime le lien entre une pharmacie et une garde sans supprimer ces ressources
// @Tags Pharmacie-garde
// @Success 200 {string} string "Opération réussie"
// @Failure 400 {string} string "Id invalide"
// @Router /pharmacie-garde/{id}/delete [post]
func SupprimerPharmacieGarde(response http.ResponseWriter, request *http.Request) {
	id, err := pathID(request, "id")
	if err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	if err := services.SupprimerPharmacieGardeService(request.Context(), id); err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	writeJSON(response, http.StatusOK, "Opération reussie!")
}
