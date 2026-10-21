package utils

import (
	"net/http"
	"strconv"
)

// Pagination est une structure qui contient les informations de pagination pour les requêtes HTTP.
type Pagination struct {
	Page   int // Le numéro de la page demandée
	Limit  int // Le nombre maximum d'éléments par page
	Offset int // Le décalage calculé à partir du numéro de page et de la limite (utilisé pour les requêtes SQL)
}

// GetPagination extrait les paramètres de pagination (page et limit) d'une requête HTTP et calcule l'offset correspondant.
func GetPagination(request *http.Request) Pagination {

	// Récupération des paramètres de pagination depuis la requête HTTP
	page, err := strconv.Atoi(request.URL.Query().Get("page"))

	// Si le paramètre "page" est absent ou invalide, on utilise la valeur par défaut (1)
	if err != nil || page < 1 {
		page = 1
	}

	// Récupération du paramètre "limit" depuis la requête HTTP et conversion en entier

	limit, err := strconv.Atoi(request.URL.Query().Get("limit"))

	// Si le paramètre "limit" est absent ou invalide, on utilise la valeur par défaut (10)
	if err != nil || limit < 1 {
		limit = 10
	}

	// Limitation de la valeur maximale de "limit" à 100 pour éviter des requêtes trop lourdes

	if limit > 100 {
		limit = 100
	}

	// Calcul de l'offset à partir du numéro de page et de la limite
	offset := (page - 1) * limit

	// Retourne une structure Pagination contenant les informations de pagination calculées
	return Pagination{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}

}
