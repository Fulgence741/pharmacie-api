package services

import (
	"context"
	"fmt"
	"log"
	"pharmacie-api/cache"
	"pharmacie-api/errors"
	"pharmacie-api/models"
	"pharmacie-api/repositories"
	"pharmacie-api/utils"
	"strings"
)

func AjouterGardeService(ctx context.Context, garde models.Garde) error {

	// Règles de validation
	if garde.DateGarde == "" {
		return errors.InvalidRequest("La date de la garde est obligatoire")
	}
	if garde.Nom_Garde == "" {
		return errors.InvalidRequest("Le nom de la garde est obligatoire")
	}
	if garde.HeureDebut == "" {
		return errors.InvalidRequest("L'heure de début est obligatoire")
	}
	if garde.HeureFin == "" {
		return errors.InvalidRequest("L'heure de fin est obligatoire")
	}
	return repositories.AjouterGardeDB(ctx, garde)
}

func ListerGardeService(ctx context.Context, pagination utils.Pagination,
	filter models.GardeFilter) ([]models.Garde, error) {

	// Gestion de cache
	cacheKey := fmt.Sprintf(
		"cache:garde:nom:%q:limit:%d:offset:%d",
		strings.ToLower(strings.TrimSpace(filter.Nom_garde)),
		pagination.Limit,
		pagination.Offset,
	)

	var garde []models.Garde

	found, err := cache.GetJSON(ctx, cacheKey, &garde)

	if err == nil && found {
		log.Print("Cache HIT ressource=garde backend=Redis")
		return garde, nil
	}

	if err != nil {
		log.Printf("Cache READ_ERROR ressource=garde backend=Redis erreur=%v", err)
	}
	log.Print("Cache MISS ressource=garde backend=PostgreSQL")

	// Appel du repository
	garde, err = repositories.ListerGardesDB(ctx,
		pagination.Limit,
		pagination.Offset,
		filter,
	)
	if err != nil {
		return nil, err
	}

	return garde, nil
}

func ObtenirGardeService(ctx context.Context, id_garde int) (models.Garde, error) {
	if id_garde <= 0 {
		return models.Garde{}, errors.InvalidRequest("Identifiant invalide")
	}

	garde, err := repositories.AfficherGardeDB(ctx, id_garde)
	if err != nil {
		return models.Garde{}, err
	}

	return garde, nil
}

func ModifierGardeService(ctx context.Context, id_garde int, garde models.Garde) error {

	if id_garde <= 0 {
		return errors.InvalidRequest("Identifiant invalide")
	}
	err := repositories.ModifierGardeDB(ctx, id_garde, garde)
	if err != nil {
		return err
	}
	return nil
}

func SupprimerGardeService(ctx context.Context, id_garde int) error {
	if id_garde <= 0 {
		return errors.InvalidRequest("Identifiant invalide")
	}

	err := repositories.SupprimerGardeDB(ctx, id_garde)
	if err != nil {
		return err
	}
	return nil
}
