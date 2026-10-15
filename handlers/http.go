package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	apierrors "pharmacie-api/errors"
)

func decodeJSON(request *http.Request, destination any) error {
	if err := json.NewDecoder(request.Body).Decode(destination); err != nil {
		return apierrors.InvalidRequest("Données invalides")
	}
	return nil
}

func pathID(request *http.Request, name string) (int, error) {
	id, err := strconv.Atoi(request.PathValue(name))
	if err != nil || id <= 0 {
		return 0, apierrors.InvalidRequest("Identifiant invalide")
	}
	return id, nil
}

func writeJSON(response http.ResponseWriter, statusCode int, value any) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	if err := json.NewEncoder(response).Encode(value); err != nil {
		log.Printf("Impossible d'écrire la réponse JSON: %v", err)
	}
}
