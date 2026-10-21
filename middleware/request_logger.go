package middleware

import (
	"log"
	"net/http"
	"time"
)

// statusRecorder est un wrapper pour http.ResponseWriter qui enregistre le code de statut HTTP de la réponse
type statusRecorder struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

// WriteHeader conserve le premier statut réellement envoyé au client.
func (r *statusRecorder) WriteHeader(statusCode int) {
	if r.wroteHeader {
		return
	}
	r.wroteHeader = true
	r.status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

// Write enregistre le statut par défaut avant d'écrire le corps de la réponse.
func (r *statusRecorder) Write(b []byte) (int, error) {
	if !r.wroteHeader {
		r.status = http.StatusOK
		r.wroteHeader = true
	}
	return r.ResponseWriter.Write(b)
}

// Unwrap permet aux utilitaires HTTP de retrouver le ResponseWriter d'origine.
func (r *statusRecorder) Unwrap() http.ResponseWriter {
	return r.ResponseWriter
}

// Logger trace chaque requête entrante avec son statut et sa durée de traitement.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter,
		request *http.Request) {
		start := time.Now()

		recorder := &statusRecorder{
			ResponseWriter: response,
			status:         http.StatusOK,
		}

		next.ServeHTTP(recorder, request)

		duration := time.Since(start)
		log.Printf(
			"Requête HTTP méthode=%s chemin=%s statut=%d durée=%s",
			request.Method,
			request.URL.Path,
			recorder.status,
			duration,
		)
	})
}
