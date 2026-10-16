package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	apierrors "pharmacie-api/errors"
)

// decodeJSON lit le corps de la requête et convertit son contenu JSON dans la cible fournie.
func decodeJSON(requete *http.Request, cible any) error {
	if erreur := json.NewDecoder(requete.Body).Decode(cible); erreur != nil {
		return apierrors.InvalidRequest("Données invalides")
	}
	return nil
}

// pathID récupère un identifiant dans le chemin et vérifie qu'il s'agit d'un entier positif.
func pathID(requete *http.Request, nomParametre string) (int, error) {
	identifiant, erreur := strconv.Atoi(requete.PathValue(nomParametre))
	if erreur != nil || identifiant <= 0 {
		return 0, apierrors.InvalidRequest("Identifiant invalide")
	}
	return identifiant, nil
}

// writeJSON définit le statut HTTP et encode le contenu de la réponse au format JSON.
func writeJSON(reponse http.ResponseWriter, codeStatut int, contenu any) {
	reponse.Header().Set("Content-Type", "application/json")
	reponse.WriteHeader(codeStatut)
	if erreur := json.NewEncoder(reponse).Encode(contenu); erreur != nil {
		log.Printf("Impossible d'écrire la réponse JSON: %v", erreur)
	}
}
