package errors

import (
	"net/http"
	"testing"
)

// TestUnauthorized vérifie que la fonction Unauthorized renvoie une erreur HTTP avec le code 401 et le message attendu.
func TestInvalidRequest(t *testing.T) {
	err := InvalidRequest("Nom invalide")

	// Vérification que l'erreur est bien de type HTTPError
	httpErr, ok := err.(*HTTPError)
	if !ok {
		t.Fatal("On attendait une HTTPError")
	}
	// Vérification du code HTTP et du message de l'erreur

	if httpErr.StatusCode != http.StatusBadRequest {
		t.Errorf(
			"Code HTTP incorrect: attendu %d, obtenu %d",
			http.StatusBadRequest,
			httpErr.StatusCode,
		)
	}

	// Vérification du message de l'erreur

	if httpErr.Message != "Nom invalide" {
		t.Errorf(
			"Message incorrect: attendu %q, obtenu %q",
			"Nom invalide",
			httpErr.Message,
		)
	}
}
