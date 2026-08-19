package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"pharmacie-api/auth"
	"pharmacie-api/users/models"
	"pharmacie-api/users/services"
	"strconv"
)

// @Summary Inscription d'un utilisateur
// @Description Permet de s'inscrire en remplissant les champs
// @Tags User
// @Success 201 {string} string "Utilisateur créé"
// @Failure 400 {string} string "Données invalides"
// @Router /user [post]
func AjouterUser(response http.ResponseWriter, request *http.Request) {
	var newUser models.User
	err := json.NewDecoder(request.Body).Decode(&newUser)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.AjouterUserService(newUser)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(response).Encode("Utilisateur créé avec succès")
	if err != nil {
		http.Error(response, "Erreur du JSON", http.StatusInternalServerError)
		return
	}
}

// @Summary Afficher les utilisateurs
// @Description Permet d'afficher la liste de tous les utilisateurs
// @Tags User
// @Produce json
// @Router /user/list [post]
func ListerUser(response http.ResponseWriter, request *http.Request) {
	user, err := services.ListerUserService()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode(user)
	if err != nil {
		http.Error(response, "Erreur du JSON", http.StatusInternalServerError)
		return
	}
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
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	userDB, err := services.Connexion(user.Email, user.Password)
	if err != nil {
		http.Error(response, err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := auth.GenererJWT(userDB)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode(map[string]string{
		"token": token,
	})
	if err != nil {
		http.Error(response, "Erreur du JSON", http.StatusInternalServerError)
		return
	}

}

// @Summary Supprimer un utilisateur
// @Description Permet à l'administrateur de supprimer un utilisateur
// @Tags User
// @Success 200 {string} string "Utilisateur supprimé avec succès"
// @Failure 400 {string} string "Id invalide"
// @Router /user/{id} [post]
func SupprimerUser(response http.ResponseWriter, request *http.Request) {
	chemin := request.PathValue("id")
	id, err := strconv.Atoi(chemin)
	if err != nil {
		http.Error(response, "Id invalide", http.StatusBadRequest)
		return
	}

	err = services.SupprimerUserService(id)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode("Utilisateur supprimé avec succès")
	if err != nil {
		http.Error(response, "Erreur du json", http.StatusInternalServerError)
		return
	}

}

// @Summary Modifier le rôle
// @Description Permet à l'admin de modifier le rôle d'un utilisateur
// @Tags User
// @Success 200 {string} string "Rôle modifié avec succès"
// @Failure 400 {string} string "Données invalides"
// @Router /user/{id}/role/patch [post]
func ModifierRole(response http.ResponseWriter, request *http.Request) {
	chemin := request.PathValue("id")

	id, err := strconv.Atoi(chemin)
	if err != nil {
		http.Error(response, "Id invalide", http.StatusBadRequest)
		return
	}

	var donnee struct {
		Role string `json:"role"`
	}

	err = json.NewDecoder(request.Body).Decode(&donnee)
	if err != nil {
		http.Error(response, "Données invalides", http.StatusBadRequest)
		return
	}

	err = services.ModifierRoleServices(id, donnee.Role)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(map[string]string{
		"message": "Rôle modifié avec succès",
	})
}

// @Summary Afficher user
// @Description Permet à l'admin et le pharmacien d'afficher un utilisateur
// @Tags User
// @Success 200 {string} string "[user]"
// @Failure 400 {string} string "Id invalide"
// @Router /user/{id}/get [post]
func AfficherUser(response http.ResponseWriter, request *http.Request) {
	chemin := request.PathValue("id")
	id, err := strconv.Atoi(chemin)
	if err != nil {
		http.Error(response, "Id invalide", http.StatusBadRequest)
		return
	}

	afficherUser, err := services.AfficherUserServices(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(response, "Utilisateur non trouvé", http.StatusNotFound)
			return
		}

		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	err = json.NewEncoder(response).Encode(afficherUser)
	if err != nil {
		http.Error(response, "Erreur du json", http.StatusInternalServerError)
		return
	}

}
