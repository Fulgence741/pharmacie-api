package services

import (
"context"
	"errors"
	"pharmacie-api/models"
	"pharmacie-api/repositories"
	"pharmacie-api/utils"
)

func AjouterGardeService(ctx context.Context, garde models.Garde) error {

	// Règles de validation
	if garde.DateGarde == "" {
		return errors.New("La date de la garde est obligatoire")
	}
	if garde.Nom_Garde == "" {
		return errors.New("Le nom de la garde est obligatoire")
	}
	if garde.HeureDebut == "" {
		return errors.New("L'Heure de debut est obligatoire")
	}
	if garde.HeureFin == "" {
		return errors.New("L'Heure de fin est obligatoire")
	}
	return repositories.AjouterGardeDB(ctx, garde)
}

func ListerGardeService(ctx context.Context, pagination utils.Pagination,
	filter models.GardeFilter) ([]models.Garde, error) {

	garde, err := repositories.ListerGardesDB( ctx,
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
		return models.Garde{}, errors.New("Identifant invalide")
	}

	garde, err := repositories.AfficherGardeDB(ctx, id_garde)
	if err != nil {
		return models.Garde{}, err
	}

	return garde, nil
}

func ModifierGardeService(ctx context.Context, id_garde int, garde models.Garde) error {

	if id_garde <= 0 {
		return errors.New("Identifiant invalide")
	}
	err := repositories.ModifierGardeDB(ctx, id_garde, garde)
	if err != nil {
		return err
	}
	return nil
}

func SupprimerGardeService(ctx context.Context, id_garde int) error {
	if id_garde <= 0 {
		return errors.New("Identifiant invalide")
	}

	err := repositories.SupprimerGardeDB(ctx, id_garde)
	if err != nil {
		return err
	}
	return nil
}
