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

// AjouterPharmacieGarde associe une pharmacie à une garde à partir de leurs identifiants.
func AjouterPharmacieGarde(reponse http.ResponseWriter, requete *http.Request) {
	var donnees dto.PharmacieGardeRequest
	if erreur := decodeJSON(requete, &donnees); erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	lien := models.PharmacieGarde{
		PharmacieID: donnees.PharmacieID,
		GardeID:     donnees.GardeID,
	}
	if erreur := services.AjouterPharmacieGardeService(requete.Context(), lien); erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	writeJSON(reponse, http.StatusOK, "Oppération reussie !")
}

// @Summary Lister les pharmacies de garde
// @Description Renvoie les pharmacies de garde disponibles en base de données
// @Tags Pharmacie-garde
// @Success 200 {array} models.PharmacieGardeA
// @Failure 500 {string} string "Erreur interne"
// @Router /pharmacie-garde/list [post]

// ListerPharmacieGarde renvoie les associations correspondant au filtre et à la pagination.
func ListerPharmacieGarde(reponse http.ResponseWriter, requete *http.Request) {
	filtre := models.PharmacieGardeFliter{Nom: requete.URL.Query().Get("nom")}
	resultat, erreur := services.ListerPharmacieGardeService(requete.Context(), utils.GetPagination(requete), filtre)
	if erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	writeJSON(reponse, http.StatusOK, pharmacieGardeResponses(resultat))
}

// @Summary Suppression de pharmacie-garde
// @Description Supprime le lien entre une pharmacie et une garde sans supprimer ces ressources
// @Tags Pharmacie-garde
// @Success 200 {string} string "Opération réussie"
// @Failure 400 {string} string "Id invalide"
// @Router /pharmacie-garde/{id}/delete [post]

// SupprimerPharmacieGarde supprime l'association, sans supprimer la pharmacie ni la garde.
func SupprimerPharmacieGarde(reponse http.ResponseWriter, requete *http.Request) {
	identifiant, erreur := pathID(requete, "id")
	if erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	if erreur := services.SupprimerPharmacieGardeService(requete.Context(), identifiant); erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	writeJSON(reponse, http.StatusOK, "Opération reussie!")
}
