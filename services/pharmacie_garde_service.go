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
	"time"
)

func AjouterPharmacieGardeService(ctx context.Context, pharmacieGarde models.PharmacieGarde) error {

	if pharmacieGarde.PharmacieID <= 0 || pharmacieGarde.GardeID <= 0 {
		return errors.InvalidRequest("Identifiant invalide")
	}

	return repositories.AjouterPharmacieGardeDB(ctx, pharmacieGarde)
}

func ListerPharmacieGardeService(
	ctx context.Context,
	pagination utils.Pagination,
	filter models.PharmacieGardeFliter,
) ([]models.PharmacieGardeA, error) {

	// 1. Construire une clé unique selon le filtre et la pagination.
	cacheKey := fmt.Sprintf(
		"cache:pharmacie_garde:nom:%q:limit:%d:offset:%d",
		strings.ToLower(strings.TrimSpace(filter.Nom)),
		pagination.Limit,
		pagination.Offset,
	)

	// 2. Préparer la variable qui recevra les données du cache.
	var pharmacieGarde []models.PharmacieGardeA

	// 3. Rechercher les données dans Redis.
	found, err := cache.GetJSON(ctx, cacheKey, &pharmacieGarde)

	// 4. Si les données existent, les retourner directement.
	if err == nil && found {
		log.Print("Cache HIT ressource=pharmacie-garde backend=Redis")
		return pharmacieGarde, nil
	}

	// 5. Signaler une éventuelle erreur Redis.
	if err != nil {
		log.Printf("Cache READ_ERROR ressource=pharmacie-garde backend=Redis erreur=%v", err)
	}

	// 6. Si le cache est absent, récupérer les données dans PostgreSQL.
	log.Print("Cache MISS ressource=pharmacie-garde backend=PostgreSQL")

	pharmacieGarde, err = repositories.ListerPharmacieGardeDB(
		ctx,
		pagination.Limit,
		pagination.Offset,
		filter,
	)
	if err != nil {
		return nil, err
	}

	// 7. Enregistrer les résultats dans Redis pendant cinq minutes.
	err = cache.SetJSON(
		ctx,
		cacheKey,
		pharmacieGarde,
		5*time.Minute,
	)

	// 8. Signaler une éventuelle erreur sans bloquer la réponse.
	if err != nil {
		log.Printf("Cache WRITE_ERROR ressource=pharmacie-garde backend=Redis erreur=%v", err)
	}

	// 9. Retourner les résultats.
	return pharmacieGarde, nil
}

func SupprimerPharmacieGardeService(ctx context.Context, id int) error {

	if id <= 0 {
		return errors.InvalidRequest("Identifiant invalide")
	}
	err := repositories.SupprimerPharmacieGardeDB(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
