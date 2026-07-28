package handlers

import (
	"encoding/json"
	"net/http"
	"pharmacie-api/auth"
	"pharmacie-api/users/models"
	"pharmacie-api/users/services"
)

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
