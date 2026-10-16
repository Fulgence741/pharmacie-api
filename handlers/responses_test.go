package handlers

import (
	"encoding/json"
	"pharmacie-api/models"
	"strings"
	"testing"
)

// NeExposePasMotDePasse vérifie que la réponse JSON
// d'un utilisateur ne contient ni son mot de passe ni le hashing.
func NeExposePasMotDePasse(t *testing.T) {
	reponse := userResponse(models.User{
		ID_USER:  7,
		Nom:      "Alice",
		Email:    "alice@example.com",
		Password: "secret-hash",
	})

	jsonEncode, erreur := json.Marshal(reponse)
	if erreur != nil {
		t.Fatalf("encodage du DTO utilisateur : %v", erreur)
	}
	if strings.Contains(string(jsonEncode), "password") || strings.Contains(string(jsonEncode), "secret-hash") {
		t.Fatalf("le DTO utilisateur expose le mot de passe : %s", jsonEncode)
	}
}
