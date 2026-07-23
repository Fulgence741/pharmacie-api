package services

import (
	"pharmacie-api/pharmacie/models"
	"pharmacie-api/pharmacie/repositories"
)

func AjouterPharmacieServices(pharmacie models.Pharmacie) error {
	return repositories.AjouterPharmacieDB(pharmacie)
}

func ListerPharmacieServices() ([]models.Pharmacie, error) {
	pharmacie, err := repositories.ListerPharmacieDB()
	if err != nil {
		return nil, err
	}
	return pharmacie, nil
}

func AfficherPharmacieServices(id_pharmacie int) (models.Pharmacie, error) {
	pharmacie, err := repositories.AfficherPharmacieDB(id_pharmacie)
	if err != nil {
		return models.Pharmacie{}, err
	}

	return pharmacie, nil
}

func ModifierPharmacieService(id_pharmacie int, pharmacie models.Pharmacie) error {
	err := repositories.ModifierPharmacieDB(id_pharmacie, pharmacie)
	if err != nil {
		return err
	}

	return nil
}

func SupprimerPharmacieServices(id_pharmacie int) error {
	err := repositories.SupprimerPharmacieDB(id_pharmacie)
	if err != nil {
		return err
	}
	return nil
}
