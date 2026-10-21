package errors

import (
	"net/http"
	"testing"
)

func TestInvalidRequest(t *testing.T) {
	err := InvalidRequest("Nom invalide")

	httpErr, ok := err.(*HTTPError)
	if !ok {
		t.Fatal("On attendait une HTTPError")
	}

	if httpErr.StatusCode != http.StatusBadRequest {
		t.Errorf(
			"Code HTTP incorrect: attendu %d, obtenu %d",
			http.StatusBadRequest,
			httpErr.StatusCode,
		)
	}

	if httpErr.Message != "Nom invalide" {
		t.Errorf(
			"Message incorrect: attendu %q, obtenu %q",
			"Nom invalide",
			httpErr.Message,
		)
	}
}
