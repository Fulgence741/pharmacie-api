package services

import (
	"errors"
	"pharmacie-api/pharmacie_garde/models"
	"pharmacie-api/pharmacie_garde/repositories"
	"pharmacie-api/utils"
)

func AjouterService(pharmacieGarde models.PharmacieGarde) error {

	if pharmacieGarde.PharmacieID <= 0 || pharmacieGarde.GardeID <= 0 {
		return errors.New("Identifiant invalide")
	}

	return repositories.AjouterDB(pharmacieGarde)
}

func ListerService(
	pagination utils.Pagination,
	filter models.PharmacieGardeFliter) ([]models.PharmacieGardeA, error) {

	pharmacieGarde, err := repositories.ListerByGardeDB(
		pagination.Limit,
		pagination.Offset,
		filter,
	)
	if err != nil {
		return nil, err
	}
	return pharmacieGarde, nil
}

func SupprimerService(id int) error {

	if id <= 0 {
		return errors.New("Id invalide")
	}
	err := repositories.SupprimerDb(id)
	if err != nil {
		return err
	}
	return nil
}
