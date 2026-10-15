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
func AjouterGarde(response http.ResponseWriter, request *http.Request) {
	var input dto.GardeRequest
	if err := decodeJSON(request, &input); err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}

	garde := models.Garde{
		Nom_Garde:  input.NomGarde,
		DateGarde:  input.DateGarde,
		HeureDebut: input.HeureDebut,
		HeureFin:   input.HeureFin,
	}
	if err := services.AjouterGardeService(request.Context(), garde); err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	writeJSON(response, http.StatusCreated, "Garde ajoutée avec succès")
}

// @Summary Liste de gardes
// @Description Permet de renvoyer la liste de toutes les gardes disponibles en base de données
// @Tags Garde
// @Success 200 {array} models.Garde
// @Failure 400 {string} string ""
// @Router /garde/list [post]
func ListerGardes(response http.ResponseWriter, request *http.Request) {
	filter := models.GardeFilter{Nom_garde: request.URL.Query().Get("nom_garde")}
	gardes, err := services.ListerGardeService(request.Context(), utils.GetPagination(request), filter)
	if err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	writeJSON(response, http.StatusOK, gardeResponses(gardes))
}

// @Summary Affichage de garde
// @Description Permet de renvoyer une garde spécifique
// @Tags Garde
// @Success 200 {object} models.Garde
// @Failure 400 {string} string "Id non valide"
// @Router /garde/{id_garde}/get [post]
func AfficherGarde(response http.ResponseWriter, request *http.Request) {
	id, err := pathID(request, "id_garde")
	if err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}

	garde, err := services.ObtenirGardeService(request.Context(), id)
	if err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	writeJSON(response, http.StatusOK, gardeResponse(garde))
}

// @Summary Modification de garde
// @Description Permet de modifier une garde spécifique
// @Tags Garde
// @Success 200 {string} string "Garde modifiée avec succès"
// @Failure 400 {string} string "Id invalide"
// @Router /garde/{id_garde}/put [post]
func ModifierGarde(response http.ResponseWriter, request *http.Request) {
	id, err := pathID(request, "id_garde")
	if err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}

	var input dto.GardeRequest
	if err := decodeJSON(request, &input); err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	garde := models.Garde{
		Nom_Garde:  input.NomGarde,
		DateGarde:  input.DateGarde,
		HeureDebut: input.HeureDebut,
		HeureFin:   input.HeureFin,
	}
	if err := services.ModifierGardeService(request.Context(), id, garde); err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	writeJSON(response, http.StatusOK, "Garde modifiée avec succès")
}

// @Summary Suppression de garde
// @Description Permet de supprimer une garde spécifique
// @Tags Garde
// @Success 200 {string} string "Garde supprimée avec succès"
// @Failure 400 {string} string "Id non valide"
// @Router /garde/{id_garde}/delete [post]
func SupprimerGarde(response http.ResponseWriter, request *http.Request) {
	id, err := pathID(request, "id_garde")
	if err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	if err := services.SupprimerGardeService(request.Context(), id); err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	writeJSON(response, http.StatusOK, "Garde supprimée avec succès")
}
