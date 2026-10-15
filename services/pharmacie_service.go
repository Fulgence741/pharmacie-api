package services

import (
	"context"
	"errors"
	"pharmacie-api/models"
	"pharmacie-api/repositories"
	"pharmacie-api/utils"
)

func AjouterPharmacieServices(ctx context.Context, pharmacie models.Pharmacie) error {

	// Règles de validation métier

	if pharmacie.Nom == "" {
		return errors.New("Le champ nom est obligatoire")
	}

	if pharmacie.Adresse == "" {
		return errors.New("Le champ adresse est obligatoire")
	}
	if pharmacie.Telephone == "" {
		return errors.New("Le chaamp telephone est obligatoire")
	}
	if pharmacie.Email == "" {
		return errors.New("le champ email est obligatoire")
	}
	if pharmacie.Ville == "" {
		return errors.New("Le champ ville est obligatoire")
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
		return models.Pharmacie{}, errors.New("Identifiant invalide")
	}
	pharmacie, err := repositories.AfficherPharmacieDB(ctx, id_pharmacie)
	if err != nil {
		return models.Pharmacie{}, err
	}

	return pharmacie, nil
}

func ModifierPharmacieService(ctx context.Context, id_pharmacie int, pharmacie models.Pharmacie) error {

	if id_pharmacie <= 0 {
		return errors.New("Identifiant invalide")
	}
	err := repositories.ModifierPharmacieDB(ctx, id_pharmacie, pharmacie)
	if err != nil {
		return err
	}

	return nil
}

func SupprimerPharmacieServices(ctx context.Context, id_pharmacie int) error {
	if id_pharmacie <= 0 {
		return errors.New("Identifiant invalide")
	}
	err := repositories.SupprimerPharmacieDB(ctx, id_pharmacie)
	if err != nil {
		return err
	}
	return nil
}
