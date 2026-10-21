package middleware

import (
	"log"
	"net/http"
	"time"
)

// statusRecorder est un wrapper pour http.ResponseWriter qui enregistre le code de statut HTTP de la réponse
type statusRecorder struct {
	http.ResponseWriter
	status int
}

// WriteHeader enregistre le code de statut et appelle la méthode WriteHeader du ResponseWriter sous-jacent
func (r *statusRecorder) WriteHeader(statusCode int) {
	r.status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

// Write enregistre le code de statut par défaut si WriteHeader n'a pas été appelé, puis appelle la méthode Write du ResponseWriter sous-jacent
func (r *statusRecorder) Write(b []byte) (int, error) {
	// Si WriteHeader n'a pas été appelé, on suppose que le code de statut est 200 OK
	if r.status == 0 {
		r.status = http.StatusOK
	}
	// Appelle la méthode Write du ResponseWriter sous-jacent
	return r.ResponseWriter.Write(b)
}

// Logger est un middleware qui enregistre les requêtes HTTP entrantes et leurs codes de statut
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter,
		request *http.Request) {
		start := time.Now()

		// Crée un statusRecorder pour capturer le code de statut de la réponse
		recorder := &statusRecorder{ResponseWriter: response,
			status: 0}

		//Passe la main au prochain middleware ou handler
		next.ServeHTTP(recorder, request)

		// Calcule la durée de traitement de la requête
		duration := time.Since(start)

		// Enregistre les informations de la requête et du code de statut
		log.Printf("%s %s %d %s", request.Method, request.URL.Path, recorder.status, duration)
	})
}
