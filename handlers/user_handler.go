package handlers

import (
	"net/http"
	"pharmacie-api/auth"
	"pharmacie-api/dto"
	apierrors "pharmacie-api/errors"
	"pharmacie-api/models"
	"pharmacie-api/services"
	"pharmacie-api/utils"
)

// Les handlers transmettent request.Context() aux services pour que l'annulation
// de la requête et ses échéances suivent tout le chemin jusqu'à la base de données.

// @Summary Inscription d'un utilisateur
// @Description Permet de s'inscrire en remplissant les champs
// @Tags User
// @Success 201 {string} string "Utilisateur créé"
// @Failure 400 {string} string "Données invalides"
// @Router /user [post]
func AjouterUser(response http.ResponseWriter, request *http.Request) {
	var input dto.CreateUserRequest
	if err := decodeJSON(request, &input); err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}

	user := models.User{
		Nom:      input.Nom,
		Email:    input.Email,
		Password: input.Password,
		Fonction: input.Fonction,
	}
	if err := services.AjouterUserService(request.Context(), user); err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}

	writeJSON(response, http.StatusCreated, "Utilisateur créé avec succès")
}

// @Summary Afficher les utilisateurs
// @Description Permet d'afficher la liste de tous les utilisateurs
// @Tags User
// @Produce json
// @Router /user/list [post]
func ListerUser(response http.ResponseWriter, request *http.Request) {
	filter := models.UserFilter{Nom: request.URL.Query().Get("nom")}
	users, err := services.ListerUserService(request.Context(), utils.GetPagination(request), filter)
	if err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	writeJSON(response, http.StatusOK, userResponses(users))
}

// @Summary Connexion utilisateur
// @Description Permet à un utilisateur de se connecter
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {string} string "Connexion réussie"
// @Failure 401 {string} string "Email ou mot de passe incorrect"
// @Router /login [post]
func ConnexionUser(response http.ResponseWriter, request *http.Request) {
	var input dto.LoginRequest
	if err := decodeJSON(request, &input); err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}

	user, err := services.Connexion(request.Context(), input.Email, input.Password)
	if err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}

	token, err := auth.GenererJWT(user)
	if err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	writeJSON(response, http.StatusOK, dto.TokenResponse{Token: token})
}

// @Summary Supprimer un utilisateur
// @Description Permet à l'administrateur de supprimer un utilisateur
// @Tags User
// @Success 200 {string} string "Utilisateur supprimé avec succès"
// @Failure 400 {string} string "Id invalide"
// @Router /user/{id}/delete [post]
func SupprimerUser(response http.ResponseWriter, request *http.Request) {
	id, err := pathID(request, "id")
	if err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	if err := services.SupprimerUserService(request.Context(), id); err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	writeJSON(response, http.StatusOK, "Utilisateur supprimé avec succès")
}

// @Summary Modifier le rôle
// @Description Permet à l'admin de modifier le rôle d'un utilisateur
// @Tags User
// @Success 200 {string} string "Rôle modifié avec succès"
// @Failure 400 {string} string "Données invalides"
// @Router /user/{id}/role/patch [post]
func ModifierRole(response http.ResponseWriter, request *http.Request) {
	id, err := pathID(request, "id")
	if err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}

	var input dto.UpdateUserRoleRequest
	if err := decodeJSON(request, &input); err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	if err := services.ModifierRoleServices(request.Context(), id, input.Role); err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	writeJSON(response, http.StatusOK, map[string]string{"message": "Rôle modifié avec succès"})
}

// @Summary Afficher user
// @Description Permet à l'admin et le pharmacien d'afficher un utilisateur
// @Tags User
// @Success 200 {string} string "[user]"
// @Failure 400 {string} string "Id invalide"
// @Router /user/{id}/get [post]
func AfficherUser(response http.ResponseWriter, request *http.Request) {
	id, err := pathID(request, "id")
	if err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}

	user, err := services.AfficherUserServices(request.Context(), id)
	if err != nil {
		apierrors.WriteHTTP(response, err)
		return
	}
	writeJSON(response, http.StatusOK, userResponse(user))
}
