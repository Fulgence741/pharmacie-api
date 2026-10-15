package services

import (
	"context"
	"errors"
	"pharmacie-api/models"
	"pharmacie-api/repositories"
	"pharmacie-api/utils"
)

func AjouterPharmacieGardeService(ctx context.Context, pharmacieGarde models.PharmacieGarde) error {

	if pharmacieGarde.PharmacieID <= 0 || pharmacieGarde.GardeID <= 0 {
		return errors.New("Identifiant invalide")
	}

	return repositories.AjouterPharmacieGardeDB(ctx, pharmacieGarde)
}

func ListerPharmacieGardeService(ctx context.Context,
	pagination utils.Pagination,
	filter models.PharmacieGardeFliter) ([]models.PharmacieGardeA, error) {

	pharmacieGarde, err := repositories.ListerPharmacieGardeDB(ctx,
		pagination.Limit,
		pagination.Offset,
		filter,
	)
	if err != nil {
		return nil, err
	}
	return pharmacieGarde, nil
}

func SupprimerPharmacieGardeService(ctx context.Context, id int) error {

	if id <= 0 {
		return errors.New("Id invalide")
	}
	err := repositories.SupprimerPharmacieGardeDB(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
