package handlers

import (
	"encoding/json"
	"net/http"
	"pharmacie-api/auth"
	"pharmacie-api/users/models"
	"pharmacie-api/users/services"
	"strconv"
)

// @Summary Inscription d'un utilisateur
// @Description permet de s'inscrire en remplissant les champs
// @Tags User
// @Succes 201 {string} string "Utilisateur créé"
// @Faillure 400 {string} string "Données invalides"
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

func ListerUser(response http.ResponseWriter, request *http.Request) {
	user, err := services.ListerUserService()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "applicatio/json")
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
// @Router /login [post]
// @Success 200 {string} string "Connexion réussie"
// @Failure 401 {string} string "Email ou mot de passe incorrect"
// @Accept json
// @Produce json
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
