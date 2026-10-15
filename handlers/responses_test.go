package handlers

import (
	"encoding/json"
	"pharmacie-api/models"
	"strings"
	"testing"
)

func TestUserResponseDoesNotExposePassword(t *testing.T) {
	response := userResponse(models.User{
		ID_USER:  7,
		Nom:      "Alice",
		Email:    "alice@example.com",
		Password: "secret-hash",
	})

	encoded, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("encodage du DTO utilisateur : %v", err)
	}
	if strings.Contains(string(encoded), "password") || strings.Contains(string(encoded), "secret-hash") {
		t.Fatalf("le DTO utilisateur expose le mot de passe : %s", encoded)
	}
}
