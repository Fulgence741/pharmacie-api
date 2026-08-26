package services

import (
	"errors"
	"pharmacie-api/garde/models"
	"pharmacie-api/garde/repositories"
	"pharmacie-api/utils"
)

func AjouterGardeService(garde models.Garde) error {

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
	return repositories.AjouterGardeDB(garde)
}

func ListerGardeService(pagination utils.Pagination,
	filter models.GardeFilter) ([]models.Garde, error) {

	garde, err := repositories.ListerGardesDB(pagination.Limit,
		pagination.Offset,
		filter,
	)
	if err != nil {
		return nil, err
	}

	return garde, nil
}

func ObtenirGardeService(id_garde int) (models.Garde, error) {
	if id_garde <= 0 {
		return models.Garde{}, errors.New("Identifant invalide")
	}

	garde, err := repositories.AfficherGardeDB(id_garde)
	if err != nil {
		return models.Garde{}, err
	}

	return garde, nil
}

func ModifierGardeService(id_garde int, garde models.Garde) error {

	if id_garde <= 0 {
		return errors.New("Identifiant invalide")
	}
	err := repositories.ModifierGardeDB(id_garde, garde)
	if err != nil {
		return err
	}
	return nil
}

func SupprimerGardeService(id_garde int) error {
	if id_garde <= 0 {
		return errors.New("Identifiant invalide")
	}

	err := repositories.SupprimerGardeDB(id_garde)
	if err != nil {
		return err
	}
	return nil
}
