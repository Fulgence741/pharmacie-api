package errors

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteHTTPUsesAPIErrorStatusAndMessage(t *testing.T) {
	response := httptest.NewRecorder()

	WriteHTTP(response, InvalidRequest("Email invalide"))

	if response.Code != http.StatusBadRequest {
		t.Fatalf("statut reçu = %d, attendu %d", response.Code, http.StatusBadRequest)
	}
	var body map[string]string
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("décodage de la réponse JSON : %v", err)
	}
	if body["error"] != "Email invalide" {
		t.Fatalf("message reçu = %q, attendu %q", body["error"], "Email invalide")
	}
}

func TestWriteHTTPMapsMissingRowsToNotFound(t *testing.T) {
	response := httptest.NewRecorder()

	WriteHTTP(response, sql.ErrNoRows)

	if response.Code != http.StatusNotFound {
		t.Fatalf("statut reçu = %d, attendu %d", response.Code, http.StatusNotFound)
	}
}

func TestWriteHTTPHidesInternalErrorDetails(t *testing.T) {
	response := httptest.NewRecorder()

	WriteHTTP(response, errTestInternal{})

	if response.Code != http.StatusInternalServerError {
		t.Fatalf("statut reçu = %d, attendu %d", response.Code, http.StatusInternalServerError)
	}
	var body map[string]string
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("décodage de la réponse JSON : %v", err)
	}
	if body["error"] != ErrInternalServer.Error() {
		t.Fatalf("message reçu = %q, attendu le message générique d'erreur", body["error"])
	}
}

type errTestInternal struct{}

func (errTestInternal) Error() string {
	return "détail interne qui ne doit pas être exposé"
}
