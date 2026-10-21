package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// startHTTPServer écoute les requêtes et arrête le serveur proprement à l'interruption.
func startHTTPServer(port string, handler http.Handler) error {
	// Créez un serveur HTTP avec un délai d'attente pour la lecture des en-têtes.
	server := &http.Server{
		Addr:              ":" + strings.TrimPrefix(port, ":"), // Assurez-vous que le port commence par ":"
		Handler:           handler,                             // Le gestionnaire pour traiter les requêtes HTTP
		ReadHeaderTimeout: 5 * time.Second,                     // Délai d'attente pour la lecture des en-têtes HTTP
	}

	// Le serveur écoute en arrière-plan pour que cette fonction puisse aussi attendre un signal d'arrêt.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	// Créez un contexte qui sera annulé à la réception d'un signal d'interruption
	defer stop() // Assurez-vous que le contexte est annulé à la fin de la fonction

	// Démarrez le serveur HTTP dans une goroutine pour ne pas bloquer l'exécution.

	serverErrors := make(chan error, 1) // Canal pour recevoir les erreurs du serveur HTTP
	go func() {                         // Goroutine pour exécuter le serveur HTTP
		log.Printf("Serveur HTTP démarré: adresse=%s", server.Addr) // Log pour indiquer que le serveur HTTP a démarré
		err := server.ListenAndServe()                              // Écoute les requêtes HTTP et retourne une erreur si le serveur s'arrête
		if errors.Is(err, http.ErrServerClosed) {                   // Vérifie si l'erreur est due à la fermeture du serveur
			err = nil
		} // Si le serveur s'arrête normalement, l'erreur est nulle
		if err != nil {
			err = fmt.Errorf("serveur HTTP: %w", err) // Si une erreur survient, l'envelopper avec un message contextuel
		}
		// Envoyer l'erreur (ou nil) dans le canal pour que la fonction principale puisse la traiter
		serverErrors <- err
	}() // Fin de la goroutine

	// Attendez soit une erreur du serveur HTTP, soit un signal d'arrêt.

	select { // Sélectionnez entre les erreurs du serveur HTTP et le signal d'arrêt
	case err := <-serverErrors: // Si le serveur HTTP retourne une erreur, la traiter
		return err
	case <-ctx.Done(): // Si un signal d'arrêt est reçu, procéder à l'arrêt du serveur HTTP
		log.Print("Signal d'arrêt reçu, arrêt du serveur HTTP...") // Log pour indiquer que le serveur HTTP va s'arrêter

		// Créez un contexte avec un délai d'attente pour l'arrêt du serveur HTTP.
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel() // S'assurer que le contexte d'arrêt est annulé à la fin de la fonction

		// Arrêtez le serveur HTTP proprement, en attendant que les requêtes en cours se terminent.
		if err := server.Shutdown(shutdownCtx); err != nil { // Arrête le serveur HTTP proprement avec un délai d'attente
			return fmt.Errorf("arrêt du serveur HTTP: %w", err)
		}
		// Attendez que la goroutine du serveur HTTP se termine et récupérez l'erreur (le cas échéant).
		return <-serverErrors
	} // Fin de la sélection entre les erreurs du serveur HTTP et le signal d'arrêt
}
