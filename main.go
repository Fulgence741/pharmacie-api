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
	// Assurez-vous que le fichier de log est fermé à la fin de l'exécution.
	defer func() {
		log.SetOutput(os.Stdout)                // Réinitialiser la sortie des logs vers le terminal
		if err := logFile.Close(); err != nil { // Fermer le fichier de log
			log.Printf("Erreur à la fermeture du fichier de log: %v", err)
		}
	}() // Fermer le fichier de log à la fin de l'exécution

	// 3. Initialiser la connexion à la base de données et à Redis.
	if err := database.ConnexionDB(cfg); err != nil {
		return err
	}
	// Assurez-vous que la connexion à la base de données est fermée à la fin de l'exécution.
	defer database.Close()
	log.Print("Connexion PostgreSQL établie")

	// Initialiser Redis
	if err := config.InitRedis(cfg.RedisAddr); err != nil {
		return err
	}
	// Assurez-vous que la connexion à Redis est fermée à la fin de l'exécution.
	defer func() {
		if err := config.CloseRedis(); err != nil {
			log.Printf("Erreur à la fermeture de Redis: %v", err)
		}
	}() // Fermer Redis à la fin de l'exécution
	log.Printf("Connexion Redis établie: adresse=%s", cfg.RedisAddr)

	// 4. Préparer les routes et lancer le serveur HTTP.
	ipLimiter := middleware.NewRateLimiter(5, time.Minute)
	userLimiter := middleware.NewRateLimiter(10, time.Minute)
	handler := routes.GestionRoutes(ipLimiter, userLimiter)

	// 5. Démarrer le serveur HTTP.
	return startHTTPServer(cfg.ServerPort, handler)
}
