package main

import (
	_ "pharmacie-api/docs"
	"pharmacie-api/swagger"

	"fmt"
	"net/http"
	"pharmacie-api/database"
	gardeRoutes "pharmacie-api/garde/routes"
	"pharmacie-api/middleware"
	pharmacieRoutes "pharmacie-api/pharmacie/routes"
	pharmacieGardeRoutes "pharmacie-api/pharmacie_garde/routes"
	UserRoutes "pharmacie-api/users/routes"
	"time"
)

// @title Pharmacie API
// @version 1.0
// @description API de gestion des pharmacies, gardes et utilisateurs
// @host localhost:8080
// @BasePath /

func main() {

	// Connexion à la base de donnée (toujours en premier)
	database.ConnexionDB()

	// Mon middleware rater limite
	ipLimiter := middleware.NewRateLimiter(5, time.Minute)    // Limite par IP
	userLimiter := middleware.NewRateLimiter(10, time.Minute) // Limite par ID
	// Charger toutes les routes
	pharmacieRoutes.GestionRoutesPharmacie(ipLimiter, userLimiter)           // Routes pharmacie
	gardeRoutes.GestionRoutesGarde(ipLimiter, userLimiter)                   // Routes garde
	pharmacieGardeRoutes.GestionRoutesPharmacieGarde(ipLimiter, userLimiter) // Routes pharmacie-garde
	UserRoutes.GestionRoutesUsers(ipLimiter, userLimiter)                    // Routes User

	// Middleware autour du routeur
	swagger.SwaggerRoutes()
	// Demarrage du serveur local (toujours en dernier)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Erreur du serveur", err)
	}

}
