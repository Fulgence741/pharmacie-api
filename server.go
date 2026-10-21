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
	server := &http.Server{
		Addr:              ":" + strings.TrimPrefix(port, ":"),
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Le serveur écoute en arrière-plan pour que cette fonction puisse aussi attendre un signal d'arrêt.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("Serveur HTTP démarré: adresse=%s", server.Addr)
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			err = nil
		}
		if err != nil {
			err = fmt.Errorf("serveur HTTP: %w", err)
		}
		serverErrors <- err
	}()

	select {
	case err := <-serverErrors:
		return err
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("arrêt du serveur HTTP: %w", err)
		}
		return <-serverErrors
	}
}
