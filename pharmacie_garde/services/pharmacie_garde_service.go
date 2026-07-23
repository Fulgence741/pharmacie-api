package services

import (
	"pharmacie-api/pharmacie_garde/models"
	"pharmacie-api/pharmacie_garde/repositories"
)

func AjouterService(pharmacieGarde models.PharmacieGarde) error {
	return repositories.AjouterDB(pharmacieGarde)
}

func ListerService() ([]models.PharmacieGardeA, error) {
	pharmacieGarde, err := repositories.ListerByGardeDB()
	if err != nil {
		return nil, err
	}
	return pharmacieGarde, nil
}

func SupprimerService(id int) error {
	err := repositories.SupprimerDb(id)
	if err != nil {
		return err
	}
	return nil
}
