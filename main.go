package main

import (
	"fmt"
	"net/http"
	"pharmacie-api/database"
	gardeRoutes "pharmacie-api/garde/routes"
	pharmacieRoutes "pharmacie-api/pharmacie/routes"
	pharmacieGardeRoutes "pharmacie-api/pharmacie_garde/routes"
)

func main() {

	database.ConnexionDB()
	pharmacieRoutes.GestionRoutesPharmacie()
	gardeRoutes.GestionRoutesGarde()
	pharmacieGardeRoutes.GestionRoutesPharmacieGarde()
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Erreur du serveur", err)
	}

}
