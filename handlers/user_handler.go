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

// Les handlers transmettent le contexte de la requête aux services afin que son
// annulation et ses échéances soient respectées jusqu'à la base de données.

// @Summary Inscription d'un utilisateur
// @Description Permet de s'inscrire en remplissant les champs
// @Tags User
// @Success 201 {string} string "Utilisateur créé"
// @Failure 400 {string} string "Données invalides"
// @Router /user [post]

// AjouterUser crée un utilisateur à partir des informations reçues en JSON.
func AjouterUser(reponse http.ResponseWriter, requete *http.Request) {
	var donnees dto.CreateUserRequest
	if erreur := decodeJSON(requete, &donnees); erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}

	utilisateur := models.User{
		Nom:      donnees.Nom,
		Email:    donnees.Email,
		Password: donnees.Password,
		Fonction: donnees.Fonction,
	}
	if erreur := services.AjouterUserService(requete.Context(), utilisateur); erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}

	writeJSON(reponse, http.StatusCreated, "Utilisateur créé avec succès")
}

// @Summary Afficher les utilisateurs
// @Description Permet d'afficher la liste de tous les utilisateurs
// @Tags User
// @Produce json
// @Router /user/list [post]
// ListerUser renvoie les utilisateurs correspondant aux filtres et à la pagination demandés.
func ListerUser(reponse http.ResponseWriter, requete *http.Request) {
	filtre := models.UserFilter{Nom: requete.URL.Query().Get("nom")}
	utilisateurs, erreur := services.ListerUserService(requete.Context(), utils.GetPagination(requete), filtre)
	if erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	writeJSON(reponse, http.StatusOK, userResponses(utilisateurs))
}

// @Summary Connexion utilisateur
// @Description Permet à un utilisateur de se connecter
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {string} string "Connexion réussie"
// @Failure 401 {string} string "Email ou mot de passe incorrect"
// @Router /login [post]
// ConnexionUser authentifie un utilisateur et renvoie son jeton d'accès.
func ConnexionUser(reponse http.ResponseWriter, requete *http.Request) {
	var donnees dto.LoginRequest
	if erreur := decodeJSON(requete, &donnees); erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}

	utilisateur, erreur := services.Connexion(requete.Context(), donnees.Email, donnees.Password)
	if erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}

	jeton, erreur := auth.GenererJWT(utilisateur)
	if erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	writeJSON(reponse, http.StatusOK, dto.TokenResponse{Token: jeton})
}

// @Summary Supprimer un utilisateur
// @Description Permet à l'administrateur de supprimer un utilisateur
// @Tags User
// @Success 200 {string} string "Utilisateur supprimé avec succès"
// @Failure 400 {string} string "Id invalide"
// @Router /user/{id}/delete [post]
// SupprimerUser supprime l'utilisateur désigné par l'identifiant du chemin.
func SupprimerUser(reponse http.ResponseWriter, requete *http.Request) {
	identifiant, erreur := pathID(requete, "id")
	if erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	if erreur := services.SupprimerUserService(requete.Context(), identifiant); erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	writeJSON(reponse, http.StatusOK, "Utilisateur supprimé avec succès")
}

// @Summary Modifier le rôle
// @Description Permet à l'admin de modifier le rôle d'un utilisateur
// @Tags User
// @Success 200 {string} string "Rôle modifié avec succès"
// @Failure 400 {string} string "Données invalides"
// @Router /user/{id}/role/patch [post]
// ModifierRole met à jour le rôle de l'utilisateur désigné par l'identifiant du chemin.
func ModifierRole(reponse http.ResponseWriter, requete *http.Request) {
	identifiant, erreur := pathID(requete, "id")
	if erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}

	var donnees dto.UpdateUserRoleRequest
	if erreur := decodeJSON(requete, &donnees); erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	if erreur := services.ModifierRoleServices(requete.Context(), identifiant, donnees.Role); erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	writeJSON(reponse, http.StatusOK, map[string]string{"message": "Rôle modifié avec succès"})
}

// @Summary Afficher un utilisateur
// @Description Permet à l'admin et le pharmacien d'afficher un utilisateur
// @Tags User
// @Success 200 {string} string "[utilisateur]"
// @Failure 400 {string} string "Id invalide"
// @Router /user/{id}/get [post]
// AfficherUser renvoie les informations de l'utilisateur désigné par l'identifiant du chemin.
func AfficherUser(reponse http.ResponseWriter, requete *http.Request) {
	identifiant, erreur := pathID(requete, "id")
	if erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}

	utilisateur, erreur := services.AfficherUserServices(requete.Context(), identifiant)
	if erreur != nil {
		apierrors.WriteHTTP(reponse, erreur)
		return
	}
	writeJSON(reponse, http.StatusOK, userResponse(utilisateur))
}
