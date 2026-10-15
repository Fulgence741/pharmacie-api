package services

import (
	"errors"
	"pharmacie-api/models"
	"pharmacie-api/repositories"
	"pharmacie-api/utils"
)

func AjouterPharmacieGardeService(pharmacieGarde models.PharmacieGarde) error {

	if pharmacieGarde.PharmacieID <= 0 || pharmacieGarde.GardeID <= 0 {
		return errors.New("Identifiant invalide")
	}

	return repositories.AjouterPharmacieGardeDB(pharmacieGarde)
}

func ListerPharmacieGardeService(
	pagination utils.Pagination,
	filter models.PharmacieGardeFliter) ([]models.PharmacieGardeA, error) {

	pharmacieGarde, err := repositories.ListerPharmacieGardeDB(
		pagination.Limit,
		pagination.Offset,
		filter,
	)
	if err != nil {
		return nil, err
	}
	return pharmacieGarde, nil
}

func SupprimerPharmacieGardeService(id int) error {

	if id <= 0 {
		return errors.New("Id invalide")
	}
	err := repositories.SupprimerPharmacieGardeDB(id)
	if err != nil {
		return err
	}
	return nil
}
