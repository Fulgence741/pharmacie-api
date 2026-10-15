package main

import (
	"fmt"
	"net/http"
	"pharmacie-api/database"
	_ "pharmacie-api/docs"
	"pharmacie-api/middleware"
	"pharmacie-api/routes"
	"time"
)

// @title Pharmacie API
// @version 1.0
// @description API de gestion des pharmacies, gardes et utilisateurs
// @host localhost:8080
// @BasePath /

func main() {

	// Initialisation de la base avant d'accepter les requêtes.
	database.ConnexionDB()

	// Les limites sont partagées par toutes les routes de l'application.
	ipLimiter := middleware.NewRateLimiter(5, time.Minute)
	userLimiter := middleware.NewRateLimiter(10, time.Minute)

	// Toutes les routes API et Swagger sont déclarées dans le routeur central.
	routes.GestionRoutes(ipLimiter, userLimiter)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur du serveur", err)
		return
	}
}
