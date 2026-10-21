package main

import (
	"log"
	"os"
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
	if err := startApplication(); err != nil {
		log.Fatalf("Erreur au démarrage ou à l'arrêt de l'API: %v", err)
	}
}

// startApplication prépare les services, puis démarre l'API.
func startApplication() error {
	// 1. Lire la configuration.
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// 2. Activer les logs dans le terminal et dans le fichier configuré.
	logFile, err := logging.InitLogger(cfg.LogFile)
	if err != nil {
		return err
	}
	defer func() {
		log.SetOutput(os.Stdout)
		if err := logFile.Close(); err != nil {
			log.Printf("Erreur à la fermeture du fichier de log: %v", err)
		}
	}()

	// 3. Connecter les services nécessaires à l'API.
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
			log.Printf("Erreur à la fermeture de Redis: %v", err)
		}
	}()
	log.Printf("Connexion Redis établie: adresse=%s", cfg.RedisAddr)

	// 4. Préparer les routes et lancer le serveur HTTP.
	ipLimiter := middleware.NewRateLimiter(5, time.Minute)
	userLimiter := middleware.NewRateLimiter(10, time.Minute)
	handler := routes.GestionRoutes(ipLimiter, userLimiter)

	return startHTTPServer(cfg.ServerPort, handler)
}
