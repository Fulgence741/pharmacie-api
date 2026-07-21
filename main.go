package main

import (
	"fmt"
	"net/http"
	"pharmacie-api/database"
	"pharmacie-api/pharmacie/routes"
)

func main() {

	database.ConnexionDB()
	routes.GestionRoutes()

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Erreur du serveur", err)
	}

}
