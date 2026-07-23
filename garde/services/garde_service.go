package services

import (
	"pharmacie-api/garde/models"
	"pharmacie-api/garde/repositories"
)

func AjouterGardeService(garde models.Garde) error {
	return repositories.AjouterGardeDB(garde)
}

func ListerGardeService() ([]models.Garde, error) {
	garde, err := repositories.ListerGardesDB()
	if err != nil {
		return nil, err
	}
	return garde, nil
}

func ObtenirGardeService(id_garde int) (models.Garde, error) {
	garde, err := repositories.AfficherGardeDB(id_garde)
	if err != nil {
		return models.Garde{}, err
	}

	return garde, nil
}

func ModifierGardeService(id_garde int, garde models.Garde) error {
	err := repositories.ModifierGardeDB(id_garde, garde)
	if err != nil {
		return err
	}
	return nil
}

func SupprimerGardeService(id_garde int) error {
	err := repositories.SupprimerGardeDB(id_garde)
	if err != nil {
		return err
	}
	return nil
}
