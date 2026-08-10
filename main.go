package main

import (
	_ "pharmacie-api/docs"
	"pharmacie-api/swagger"

	"fmt"
	"net/http"
	"pharmacie-api/database"
	gardeRoutes "pharmacie-api/garde/routes"
	pharmacieRoutes "pharmacie-api/pharmacie/routes"
	pharmacieGardeRoutes "pharmacie-api/pharmacie_garde/routes"
	UserRoutes "pharmacie-api/users/routes"
)

// @title Pharmacie API
// @version 1.0
// @description API de gestion des pharmacies, gardes et utilisateurs
// @host localhost:8080
// @BasePath /

func main() {

	// Connexion à la base de donnée (toujours en premier)
	database.ConnexionDB()
	// Charger toutes les routes
	pharmacieRoutes.GestionRoutesPharmacie()
	gardeRoutes.GestionRoutesGarde()
	pharmacieGardeRoutes.GestionRoutesPharmacieGarde()
	UserRoutes.GestionRoutesUsers()
	swagger.SwaggerRoutes()
	// Demarrage du serveur local (toujours en dernier)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Erreur du serveur", err)
	}

}
