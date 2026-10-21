package errors

import (
	"database/sql"
	"encoding/json"
	stderrors "errors"
	"log"
	"net/http"
)

// HTTPError décrit une erreur que l'API peut transmettre au client.
type HTTPError struct {
	StatusCode int
	Message    string
}

func (err *HTTPError) Error() string {
	return err.Message
}

// InvalidRequest crée une erreur de validation destinée à être renvoyée en 400.
func InvalidRequest(message string) error {
	return &HTTPError{StatusCode: http.StatusBadRequest, Message: message}
}

// NotFound crée une erreur destinée à être renvoyée en 404.
func NotFound(message string) error {
	return &HTTPError{StatusCode: http.StatusNotFound, Message: message}
}

// Unauthorized crée une erreur d'authentification destinée à être renvoyée en 401.
func Unauthorized(message string) error {
	return &HTTPError{StatusCode: http.StatusUnauthorized, Message: message}
}

// Forbidden crée une erreur d'autorisation destinée à être renvoyée en 403.
func Forbidden(message string) error {
	return &HTTPError{StatusCode: http.StatusForbidden, Message: message}
}

// TooManyRequests crée une erreur de limitation destinée à être renvoyée en 429.
func TooManyRequests(message string) error {
	return &HTTPError{StatusCode: http.StatusTooManyRequests, Message: message}
}

// WriteHTTP écrit toutes les erreurs de l'API avec le même format JSON.
func WriteHTTP(response http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	// Définition des valeurs par défaut pour le code HTTP et le message d'erreur
	statusCode := http.StatusInternalServerError
	message := ErrInternalServer.Error()

	// Vérification du type d'erreur pour déterminer le code HTTP et le message appropriés
	var httpErr *HTTPError
	if stderrors.As(err, &httpErr) { // Vérification si l'erreur est de type HTTPError
		statusCode = httpErr.StatusCode // Si c'est le cas, on utilise le code HTTP et le message de l'erreur HTTPError
		message = httpErr.Message       // On utilise le message de l'erreur HTTPError
	} else if stderrors.Is(err, sql.ErrNoRows) { // Vérification si l'erreur est une erreur de type sql.ErrNoRows (pas de résultats trouvés)
		statusCode = http.StatusNotFound // Si c'est le cas, on renvoie un code HTTP 404 Not Found
		message = ErrNotFound.Error()    // On utilise le message de l'erreur ErrNotFound
	} else {
		log.Printf("Erreur interne API: erreur=%v", err)
	}

	// Écriture de la réponse HTTP avec le code et le message d'erreur appropriés
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	if encodeErr := json.NewEncoder(response).Encode(map[string]string{
		"error": message,
	}); encodeErr != nil {
		log.Printf("Échec écriture réponse JSON d'erreur: erreur=%v", encodeErr)
	}
}
