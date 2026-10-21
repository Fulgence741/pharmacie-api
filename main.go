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

	"pharmacie-api/config"
	"pharmacie-api/database"
	_ "pharmacie-api/docs"
	"pharmacie-api/logging"
	"pharmacie-api/middleware"
	"pharmacie-api/routes"
)

// @title Pharmacie API
// @version 1.0
// @description API de gestion des pharmacies, gardes et utilisateurs
// @host localhost:8080
// @BasePath /

func main() {
	if err := run(); err != nil {
		log.Fatalf("Erreur au démarrage ou à l'arrêt de l'API: %v", err)
	}
}

func run() (runErr error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	logFile, err := logging.InitLogger(cfg.LogFile)
	if err != nil {
		return fmt.Errorf("initialisation des logs: %w", err)
	}
	defer func() {
		log.SetOutput(os.Stdout)
		if err := logFile.Close(); err != nil {
			runErr = errors.Join(runErr, fmt.Errorf("fermeture du fichier de log: %w", err))
		}
	}()

	if err := database.ConnexionDB(cfg); err != nil {
		return err
	}
	defer database.Close()
	log.Print("Connexion PostgreSQL établie")

	if err := config.InitRedis(cfg.RedisAddr); err != nil {
		return err
	}
	defer func() {
		if err := config.CloseRedis(); err != nil {
			runErr = errors.Join(runErr, err)
		}
	}()
	log.Printf("Connexion Redis établie: adresse=%s", cfg.RedisAddr)

	ipLimiter := middleware.NewRateLimiter(5, time.Minute)
	userLimiter := middleware.NewRateLimiter(10, time.Minute)
	handler := routes.GestionRoutes(ipLimiter, userLimiter)

	port := strings.TrimPrefix(cfg.ServerPort, ":")
	server := &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("Serveur HTTP démarré: adresse=%s", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- fmt.Errorf("serveur HTTP: %w", err)
		} else {
			serverErrors <- nil
		}
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
