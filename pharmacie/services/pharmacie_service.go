package services

import (
	"errors"
	"pharmacie-api/pharmacie/models"
	"pharmacie-api/pharmacie/repositories"
)

func AjouterPharmacieServices(pharmacie models.Pharmacie) error {

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
