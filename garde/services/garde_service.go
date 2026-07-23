package services

import (
	"errors"
	"pharmacie-api/garde/models"
	"pharmacie-api/garde/repositories"
)

func AjouterGardeService(garde models.Garde) error {

	if garde.DateGarde == "" {
		return errors.New("Le champ Date est obligatoires")
	}
	if garde.Nom_Garde == "" {
		return errors.New("Le champ Nom est obligatoire")
	}
	if garde.HeureDebut == "" {
		return errors.New("Le champ Heure de debut est obligatoire")
	}
	if garde.HeureFin == "" {
		return errors.New("Le champ Heure de fin est obligatoire")
	}
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
