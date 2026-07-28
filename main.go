package main

import (
	"fmt"
	"net/http"
	"pharmacie-api/database"
	gardeRoutes "pharmacie-api/garde/routes"
	pharmacieRoutes "pharmacie-api/pharmacie/routes"
	pharmacieGardeRoutes "pharmacie-api/pharmacie_garde/routes"
	UserRoutes "pharmacie-api/users/routes"
)

func main() {

	// Connexion à la base de donnée (toujours en premier)
	database.ConnexionDB()
	// Charger toutes les routes
	pharmacieRoutes.GestionRoutesPharmacie()
	gardeRoutes.GestionRoutesGarde()
	pharmacieGardeRoutes.GestionRoutesPharmacieGarde()
	UserRoutes.GestionRoutesUsers()
	// Demarrage du serveur local (toujours en dernier)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Erreur du serveur", err)
	}

}
