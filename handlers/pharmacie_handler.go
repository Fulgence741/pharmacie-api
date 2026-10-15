package handlers

import (
	"net/http"
	"pharmacie-api/dto"
	apierrors "pharmacie-api/errors"
	"pharmacie-api/models"
	"pharmacie-api/services"
	"pharmacie-api/utils"
)

// @Summary Ajouter une pharmacie
// @Description Permet d'ajouter une pharmacie
// @Tags Pharmacie
// @Success 201 {string} string "Pharmacie ajoutée avec succès"
// @Failure 400 {string} string ""
// @Router /pharmacie [post]
func AjouterPharmacie(response http.ResponseWriter, request *http.Request) {
	var input dto.PharmacieRequest
	if err := decodeJSON(request, &input); err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	pharmacie := pharmacieFromRequest(input)
	if err := services.AjouterPharmacieServices(request.Context(), pharmacie); err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	writeJSON(response, http.StatusCreated, "Pharmacie ajouté avec succès")
}

// @Summary Lister pharmacie
// @Description Permet de renvoyer la liste de toutes les pharmacies disponibles en base de données
// @Tags Pharmacie
// @Success 200 {array} models.Pharmacie
// @Failure 400 {string} string ""
// @Router /pharmacie/list [post]
func ListerPharmacie(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	filter := models.PharmacieFilter{
		Nom:    query.Get("nom"),
		Ville:  query.Get("ville"),
		Status: query.Get("status"),
	}
	pharmacies, err := services.ListerPharmacieServices(request.Context(), utils.GetPagination(request), filter)
	if err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	writeJSON(response, http.StatusOK, pharmacieResponses(pharmacies))
}

// @Summary Afficher une pharmacie
// @Description Permet de renvoyer une pharmacie spécifique
// @Tags Pharmacie
// @Success 200 {object} models.Pharmacie
// @Failure 400 {string} string "Id invalide"
// @Router /pharmacie/{id_pharmacie}/get [post]
func AfficherPharmacie(response http.ResponseWriter, request *http.Request) {
	id, err := pathID(request, "id_pharmacie")
	if err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	pharmacie, err := services.AfficherPharmacieServices(request.Context(), id)
	if err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	writeJSON(response, http.StatusOK, pharmacieResponse(pharmacie))
}

// @Summary Modification d'une pharmacie
// @Description Permet de modifier une pharmacie spécifique
// @Tags Pharmacie
// @Success 200 {string} string "Pharmacie modifiée avec succès"
// @Failure 400 {string} string "Id invalide"
// @Router /pharmacie/{id_pharmacie}/put [post]
func ModifierPharmacie(response http.ResponseWriter, request *http.Request) {
	id, err := pathID(request, "id_pharmacie")
	if err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}

	var input dto.PharmacieRequest
	if err := decodeJSON(request, &input); err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	if err := services.ModifierPharmacieService(request.Context(), id, pharmacieFromRequest(input)); err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	writeJSON(response, http.StatusOK, "Pharmacie modifiée avec succès")
}

// @Summary Suppression d'une pharmacie
// @Description Permet de supprimer une pharmacie spécifique
// @Tags Pharmacie
// @Success 200 {string} string "Pharmacie supprimée avec succès"
// @Failure 400 {string} string "Id invalide"
// @Router /pharmacie/{id_pharmacie}/delete [post]
func SupprimerPharmacie(response http.ResponseWriter, request *http.Request) {
	id, err := pathID(request, "id_pharmacie")
	if err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	if err := services.SupprimerPharmacieServices(request.Context(), id); err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	writeJSON(response, http.StatusOK, "Pharmacie supprimée avec succès")
}

func pharmacieFromRequest(input dto.PharmacieRequest) models.Pharmacie {
	return models.Pharmacie{
		Nom:       input.Nom,
		Adresse:   input.Adresse,
		Telephone: input.Telephone,
		Email:     input.Email,
		Ville:     input.Ville,
	}
}
