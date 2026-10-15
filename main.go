package main

import (
	"fmt"
	"log"
	"net/http"
	"pharmacie-api/database"
	_ "pharmacie-api/docs"
	"pharmacie-api/middleware"
	"pharmacie-api/routes"
	"time"

	"github.com/joho/godotenv"
)

// @title Pharmacie API
// @version 1.0
// @description API de gestion des pharmacies, gardes et utilisateurs
// @host localhost:8080
// @BasePath /

func main() {

	// Initialisation de la base avant d'accepter les requêtes.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur du chargement du fichier .env :", err)
	}

	config := database.LoadConfig()
	err = database.ConnexionDB(config)

	if err != nil {
		log.Fatal("Erreur de la connexion à postgreSQL:", err)
	}

	// Les limites sont partagées par toutes les routes de l'application.
	ipLimiter := middleware.NewRateLimiter(5, time.Minute)
	userLimiter := middleware.NewRateLimiter(10, time.Minute)

	// Toutes les routes sont déclarées dans le routeur central.
	routes.GestionRoutes(ipLimiter, userLimiter)

	//Demarrer le serveur HTTP
	fmt.Println("Serveur demarré sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
