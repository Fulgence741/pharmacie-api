package handlers

import (
	"net/http"
	"pharmacie-api/models"
	"pharmacie-api/services"

	"pharmacie-api/dto"
	apierrors "pharmacie-api/errors"
	"pharmacie-api/utils"
)

// @Summary Ajouter une pharmacie
// @Description Permet d'ajouter une pharmacie
// @Tags Pharmacie
// @Success 201 {string} string "Pharmacie ajoutée avec succès"
// @Failure 400 {string} string ""
// @Router /pharmacie [post]

// AjouterPharmacie crée une pharmacie à partir des informations reçues en JSON.
func AjouterPharmacie(reponse http.ResponseWriter, requete *http.Request) {
	var donnees dto.PharmacieRequest
	if erreur := decodeJSON(requete, &donnees); erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	pharmacie := pharmacieFromRequest(donnees)
	if erreur := services.AjouterPharmacieServices(requete.Context(), pharmacie); erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	writeJSON(reponse, http.StatusCreated, "Pharmacie ajouté avec succès")
}

// @Summary Lister pharmacie
// @Description Permet de renvoyer la liste de toutes les pharmacies disponibles en base de données
// @Tags Pharmacie
// @Success 200 {array} models.Pharmacie
// @Failure 400 {string} string ""
// @Router /pharmacie/list [post]

// ListerPharmacie renvoie les pharmacies selon les critères de recherche et la pagination.
func ListerPharmacie(reponse http.ResponseWriter, requete *http.Request) {
	parametres := requete.URL.Query()
	filtre := models.PharmacieFilter{
		Nom:    parametres.Get("nom"),
		Ville:  parametres.Get("ville"),
		Status: parametres.Get("status"),
	}
	pharmacies, erreur := services.ListerPharmacieServices(requete.Context(), utils.GetPagination(requete), filtre)
	if erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	writeJSON(reponse, http.StatusOK, pharmacieResponses(pharmacies))
}

// @Summary Afficher une pharmacie
// @Description Permet de renvoyer une pharmacie spécifique
// @Tags Pharmacie
// @Success 200 {object} models.Pharmacie
// @Failure 400 {string} string "Id invalide"
// @Router /pharmacie/{id_pharmacie}/get [post]

// AfficherPharmacie renvoie la pharmacie correspondant à l'identifiant du chemin.
func AfficherPharmacie(reponse http.ResponseWriter, requete *http.Request) {
	identifiant, erreur := pathID(requete, "id_pharmacie")
	if erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	pharmacie, erreur := services.AfficherPharmacieServices(requete.Context(), identifiant)
	if erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	writeJSON(reponse, http.StatusOK, pharmacieResponse(pharmacie))
}

// @Summary Modification d'une pharmacie
// @Description Permet de modifier une pharmacie spécifique
// @Tags Pharmacie
// @Success 200 {string} string "Pharmacie modifiée avec succès"
// @Failure 400 {string} string "Id invalide"
// @Router /pharmacie/{id_pharmacie}/put [post]

// ModifierPharmacie met à jour la pharmacie désignée par l'identifiant du chemin.
func ModifierPharmacie(reponse http.ResponseWriter, requete *http.Request) {
	identifiant, erreur := pathID(requete, "id_pharmacie")
	if erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}

	var donnees dto.PharmacieRequest
	if erreur := decodeJSON(requete, &donnees); erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	if erreur := services.ModifierPharmacieService(requete.Context(), identifiant, pharmacieFromRequest(donnees)); erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	writeJSON(reponse, http.StatusOK, "Pharmacie modifiée avec succès")
}

// @Summary Suppression d'une pharmacie
// @Description Permet de supprimer une pharmacie spécifique
// @Tags Pharmacie
// @Success 200 {string} string "Pharmacie supprimée avec succès"
// @Failure 400 {string} string "Id invalide"
// @Router /pharmacie/{id_pharmacie}/delete [post]

// SupprimerPharmacie supprime la pharmacie désignée par l'identifiant du chemin.
func SupprimerPharmacie(reponse http.ResponseWriter, requete *http.Request) {
	identifiant, erreur := pathID(requete, "id_pharmacie")
	if erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	if erreur := services.SupprimerPharmacieServices(requete.Context(), identifiant); erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	writeJSON(reponse, http.StatusOK, "Pharmacie supprimée avec succès")
}

// pharmacieFromRequest convertit les données reçues en modèle de pharmacie.
func pharmacieFromRequest(donnees dto.PharmacieRequest) models.Pharmacie {
	return models.Pharmacie{
		Nom:       donnees.Nom,
		Adresse:   donnees.Adresse,
		Telephone: donnees.Telephone,
		Email:     donnees.Email,
		Ville:     donnees.Ville,
	}
}
