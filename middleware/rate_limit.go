package middleware

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

// Objet qui représente représente l'état de n'importe quelle clé.

type ClientState struct {
	Count int       // Stock le nombre de requete effectué
	Start time.Time // stock le moment ou la fenêtre actuelle de limitation commence
}

// Objet qui rassemble toutes les informations et outils nécessaires pour appliquer une limitation de requêtes.

type RateLimiter struct {
	limit   int                    // le nombre maximum de requête
	window  time.Duration          // la durée pendant laquelle on compte les requêtes
	clients map[string]ClientState //l'état associé à une clé qui identifie le client.
	mutex   sync.Mutex             // Vérrou pour ne pas que deux exécutions du concurents du code modifient la map en même temps
}

// Construction d'un rate limit
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:   limit,
		window:  window,
		clients: make(map[string]ClientState),
	}
}

func (limiter *RateLimiter) Allow(key string) bool {
	limiter.mutex.Lock()         // vérouille l'opération pendant l'execution par une clé
	defer limiter.mutex.Unlock() // defer exécute le unclock à la fin de la fonction

	now := time.Now() // Récupération de l'heure actuelle et la met dans la variable now

	state, exists := limiter.clients[key] // Donne-moi l'état associé à cette clé et dis-moi si cette clé existe.

	if !exists { // Si la clé n'existe pas

		// Dans ma map, à cette clé, enregistre ce ClientState
		limiter.clients[key] = ClientState{
			Count: 1,
			Start: now,
		}
		return true // Oui
	}

	// La fenêtre est terminée
	if now.Sub(state.Start) >= limiter.window { // vérifier entre les deux temps  , et voir s'il est supérieur ou égale à la limite autorisée
		limiter.clients[key] = ClientState{
			Count: 1,
			Start: now, // si c'est vrai alors on réinitialise et on commence une nouvelle fenêtre de requette
		}
		return true
	}

	// Limite atteinte
	if state.Count >= limiter.limit { //  on vérifie la limite
		return false // la requête est refusée
	}

	// Incrémentation du compteur
	state.Count++                // sinon on incrémente
	limiter.clients[key] = state // on remet l'état qui a été modifié dans la map

	return true // la requête est autorisée
}

// Fonction qui sert à supprimer les clients qui ont terminés leur fenêtre
func (limiter *RateLimiter) Cleanup() {
	limiter.mutex.Lock()         // Verouille
	defer limiter.mutex.Unlock() // Deverouille à la fin de la fonction

	now := time.Now() // prends le temps de maintenant

	for key, state := range limiter.clients { // entre dans la map et parcours là
		if now.Sub(state.Start) >= limiter.window { // si la durée écoulée depuis le début de sa fenêtre est supérieure ou égale à la durée de la fenêtre
			delete(limiter.clients, key) // supprime cette clé
		}
	}
}

// Le middleware pour implémenter les deux, Limite par IP et limite par ID
func RateLimit(
	ipLimiter *RateLimiter, // controle les ip
	userLimiter *RateLimiter, // controle les utilisateurs authentifiés
) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

			// ========================================
			// 2. Rate limit par user ID
			// ========================================
			userID, ok := request.Context().Value("id").(int)

			if ok && userLimiter != nil {

				// L'utilisateur est authentifié.
				// On utilise donc son User ID comme clé.

				userKey := fmt.Sprintf("user:%d", userID)

				if !userLimiter.Allow(userKey) {
					http.Error(
						response,
						"Too Many Requests",
						http.StatusTooManyRequests,
					)
					return
				}

				// La limite utilisateur est respectée.
				next.ServeHTTP(response, request)
				return
			}

			// ========================================
			// 2. Rate limit par IP
			// ========================================
			ip, _, err := net.SplitHostPort(request.RemoteAddr)

			if err != nil {
				http.Error(
					response,
					"Adresse IP invalide",
					http.StatusBadRequest,
				)
				return
			}

			// On utilise l'IP comme clé.

			if !ipLimiter.Allow("ip:" + ip) {
				http.Error(
					response,
					"Too Many Requests",
					http.StatusTooManyRequests,
				)
				return
			}

			// Si tout se passe bien on passe la main au handler pour qu'il puisse être exécuté
			next.ServeHTTP(response, request)
		})
	}
}
