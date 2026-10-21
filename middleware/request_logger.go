package middleware

import (
	"log"
	"net/http"
	"time"
)

// statusRecorder enregistre le code de statut HTTP de la réponse.
type statusRecorder struct {
	http.ResponseWriter      // Le ResponseWriter d'origine pour écrire la réponse HTTP
	status              int  // Le code de statut HTTP de la réponse
	wroteHeader         bool // Indique si WriteHeader a été appelé
}

// WriteHeader conserve le premier statut réellement envoyé au client.
func (r *statusRecorder) WriteHeader(statusCode int) {
	if r.wroteHeader { // Si WriteHeader a déjà été appelé, on ne fait rien pour éviter d'écraser le code de statut précédent
		return
	}
	r.wroteHeader = true                     // Indique que WriteHeader a été appelé
	r.status = statusCode                    // Enregistre le code de statut HTTP de la réponse
	r.ResponseWriter.WriteHeader(statusCode) // Appelle la méthode WriteHeader du ResponseWriter d'origine pour envoyer le code de statut au client
}

// Write enregistre le statut par défaut avant d'écrire le corps de la réponse.
func (r *statusRecorder) Write(b []byte) (int, error) {
	if !r.wroteHeader { // Si WriteHeader n'a pas encore été appelé, on enregistre le statut par défaut (200 OK)

		r.status = http.StatusOK // Enregistre le code de statut HTTP par défaut (200 OK)

		r.ResponseWriter.WriteHeader(http.StatusOK) // Appelle la méthode WriteHeader du ResponseWriter d'origine pour envoyer le code de statut par défaut au client
		r.wroteHeader = true
	}
	return r.ResponseWriter.Write(b) // Appelle la méthode Write du ResponseWriter d'origine pour écrire le corps de la réponse
}

// Unwrap permet aux utilitaires HTTP de retrouver le ResponseWriter d'origine.
func (r *statusRecorder) Unwrap() http.ResponseWriter {
	return r.ResponseWriter
}

// Logger trace chaque requête entrante avec son statut et sa durée de traitement.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter,
		request *http.Request) {
		start := time.Now() // Enregistre le temps de début de la requête HTTP

		// Création d'un statusRecorder pour intercepter le code de statut HTTP de la réponse
		recorder := &statusRecorder{
			ResponseWriter: response,      // Le ResponseWriter d'origine pour écrire la réponse HTTP
			status:         http.StatusOK, // Le code de statut HTTP par défaut (200 OK)
		}

		// Passe la main au prochain middleware ou handler
		next.ServeHTTP(recorder, request)

		// Code exécuté après le handler
		duration := time.Since(start) // Calcule la durée de traitement de la requête HTTP

		// Trace la requête HTTP avec son statut et sa durée de traitement
		log.Printf(
			"Requête HTTP méthode=%s chemin=%s statut=%d durée=%s", // Log format: HTTP request method, path, status code, and duration
			request.Method,   // Log the HTTP request method (GET, POST, etc.)
			request.URL.Path, // Log the HTTP request path (URL)
			recorder.status,  // Log the HTTP response status code (200, 404, etc.)
			duration,         // Log the duration of the HTTP request processing
		)
	})
}
