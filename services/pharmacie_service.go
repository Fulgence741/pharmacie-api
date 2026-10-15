package services

import (
	"context"
	"pharmacie-api/errors"
	"pharmacie-api/models"
	"pharmacie-api/repositories"
	"pharmacie-api/utils"
)

func AjouterPharmacieServices(ctx context.Context, pharmacie models.Pharmacie) error {

	// Règles de validation métier

	if pharmacie.Nom == "" {
		return errors.InvalidRequest("Le champ nom est obligatoire")
	}

	if pharmacie.Adresse == "" {
		return errors.InvalidRequest("Le champ adresse est obligatoire")
	}
	if pharmacie.Telephone == "" {
		return errors.InvalidRequest("Le champ téléphone est obligatoire")
	}
	if pharmacie.Email == "" {
		return errors.InvalidRequest("Le champ email est obligatoire")
	}
	if pharmacie.Ville == "" {
		return errors.InvalidRequest("Le champ ville est obligatoire")
	}

	return repositories.AjouterPharmacieDB(ctx, pharmacie)
}

func ListerPharmacieServices(ctx context.Context,
	pagination utils.Pagination,
	filter models.PharmacieFilter) ([]models.Pharmacie, error) {

	pharmacie, err := repositories.ListerPharmacieDB(ctx,
		pagination.Limit,
		pagination.Offset,
		filter,
	)
	if err != nil {
		return nil, err
	}

	return pharmacie, nil
}

func AfficherPharmacieServices(ctx context.Context, id_pharmacie int) (models.Pharmacie, error) {

	if id_pharmacie <= 0 {
		return models.Pharmacie{}, errors.InvalidRequest("Identifiant invalide")
	}
	pharmacie, err := repositories.AfficherPharmacieDB(ctx, id_pharmacie)
	if err != nil {
		return models.Pharmacie{}, err
	}

	return pharmacie, nil
}

func ModifierPharmacieService(ctx context.Context, id_pharmacie int, pharmacie models.Pharmacie) error {

	if id_pharmacie <= 0 {
		return errors.InvalidRequest("Identifiant invalide")
	}
	err := repositories.ModifierPharmacieDB(ctx, id_pharmacie, pharmacie)
	if err != nil {
		return err
	}

	return nil
}

func SupprimerPharmacieServices(ctx context.Context, id_pharmacie int) error {
	if id_pharmacie <= 0 {
		return errors.InvalidRequest("Identifiant invalide")
	}
	err := repositories.SupprimerPharmacieDB(ctx, id_pharmacie)
	if err != nil {
		return err
	}
	return nil
}
