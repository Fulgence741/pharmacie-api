package handlers

import (
	"net/http"
	"pharmacie-api/dto"
	apierrors "pharmacie-api/errors"
	"pharmacie-api/models"
	"pharmacie-api/services"
	"pharmacie-api/utils"
)

// @Summary ajout d'une garde
// @Description Permet d'ajouter une garde
// @Tags Garde
// @Success 201 {string} string "Garde ajoutée avec succès"
// @Failure 400 {string} string ""
// @Router /garde [post]

// AjouterGarde crée une garde à partir des informations reçues en JSON.
func AjouterGarde(reponse http.ResponseWriter, requete *http.Request) {
	var donnees dto.GardeRequest
	if erreur := decodeJSON(requete, &donnees); erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}

	garde := models.Garde{
		Nom_Garde:  donnees.NomGarde,
		DateGarde:  donnees.DateGarde,
		HeureDebut: donnees.HeureDebut,
		HeureFin:   donnees.HeureFin,
	}
	if erreur := services.AjouterGardeService(requete.Context(), garde); erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	writeJSON(reponse, http.StatusCreated, "Garde ajoutée avec succès")
}

// @Summary Liste de gardes
// @Description Permet de renvoyer la liste de toutes les gardes disponibles en base de données
// @Tags Garde
// @Success 200 {array} models.Garde
// @Failure 400 {string} string ""
// @Router /garde/list [post]

// ListerGardes renvoie les gardes correspondant aux filtres et à la pagination demandés.
func ListerGardes(reponse http.ResponseWriter, requete *http.Request) {
	filtre := models.GardeFilter{Nom_garde: requete.URL.Query().Get("nom_garde")}
	gardes, erreur := services.ListerGardeService(requete.Context(), utils.GetPagination(requete), filtre)
	if erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	writeJSON(reponse, http.StatusOK, gardeResponses(gardes))
}

// @Summary Affichage de garde
// @Description Permet de renvoyer une garde spécifique
// @Tags Garde
// @Success 200 {object} models.Garde
// @Failure 400 {string} string "Id non valide"
// @Router /garde/{id_garde}/get [post]

// AfficherGarde renvoie la garde correspondant à l'identifiant du chemin.
func AfficherGarde(reponse http.ResponseWriter, requete *http.Request) {
	identifiant, erreur := pathID(requete, "id_garde")
	if erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}

	garde, erreur := services.ObtenirGardeService(requete.Context(), identifiant)
	if erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	writeJSON(reponse, http.StatusOK, gardeResponse(garde))
}

// @Summary Modification de garde
// @Description Permet de modifier une garde spécifique
// @Tags Garde
// @Success 200 {string} string "Garde modifiée avec succès"
// @Failure 400 {string} string "Id invalide"
// @Router /garde/{id_garde}/put [post]

// ModifierGarde met à jour la garde désignée par l'identifiant du chemin.
func ModifierGarde(reponse http.ResponseWriter, requete *http.Request) {
	identifiant, erreur := pathID(requete, "id_garde")
	if erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}

	var donnees dto.GardeRequest
	if erreur := decodeJSON(requete, &donnees); erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	garde := models.Garde{
		Nom_Garde:  donnees.NomGarde,
		DateGarde:  donnees.DateGarde,
		HeureDebut: donnees.HeureDebut,
		HeureFin:   donnees.HeureFin,
	}
	if erreur := services.ModifierGardeService(requete.Context(), identifiant, garde); erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	writeJSON(reponse, http.StatusOK, "Garde modifiée avec succès")
}

// @Summary Suppression de garde
// @Description Permet de supprimer une garde spécifique
// @Tags Garde
// @Success 200 {string} string "Garde supprimée avec succès"
// @Failure 400 {string} string "Id non valide"
// @Router /garde/{id_garde}/delete [post]

// SupprimerGarde supprime la garde désignée par l'identifiant du chemin.
func SupprimerGarde(reponse http.ResponseWriter, requete *http.Request) {
	identifiant, erreur := pathID(requete, "id_garde")
	if erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	if erreur := services.SupprimerGardeService(requete.Context(), identifiant); erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	writeJSON(reponse, http.StatusOK, "Garde supprimée avec succès")
}
