package repositories

import (
	"fmt"
	"pharmacie-api/database"
	"pharmacie-api/pharmacie/models"
)

func AjouterPharmacieDB(newPharmacie models.Pharmacie) error {

	// Requête pour pour ajouter une nouvelle pharmacie
	requete := `
					INSERT INTO pharmacies (nom,
					 adresse, 
					 telephone,
					  email,
					   ville)
					VALUES ($1,
					 $2,
					  $3,
					   $4,
					    $5)	
	`
	_, err := database.DB.Exec(requete,
		newPharmacie.Nom,
		newPharmacie.Adresse,
		newPharmacie.Telephone,
		newPharmacie.Email,
		newPharmacie.Ville)

	return err

}

func ListerPharmacieDB(
	limit int,
	offset int,
	filter models.PharmacieFilter) ([]models.Pharmacie, error) {

	// Requête pour lister toutes les pharmacies disponibles en base de donnée
	requete := `
						SELECT id_pharmacie,
						 nom, 
						 adresse,
						  telephone, 
						  email,
						   ville ,
						   status FROM pharmacies
						   
		`
	var args []interface{}
	paramIndex := 1
	hasWhere := false

	// Filtrer par nom
	if filter.Nom != "" {
		requete += fmt.Sprintf(" WHERE nom ILIKE $%d", paramIndex)
		args = append(args, "%"+filter.Nom+"%")
		paramIndex++
		hasWhere = true
	}

	// Filtrer par ville
	if filter.Ville != "" {
		if hasWhere {
			requete += fmt.Sprintf(" AND ville ILIKE $%d",
				paramIndex)
		} else {
			requete += fmt.Sprintf(" WHERE ville ILIKE $%d",
				paramIndex)
		}

		args = append(args, "%"+filter.Ville+"%")
		paramIndex++
		hasWhere = true
	}

	// Filtrer par status
	if filter.Status != "" {
		if hasWhere {
			requete += fmt.Sprintf(
				" AND status = $%d",
				paramIndex,
			)
		} else {
			requete += fmt.Sprintf(" WHERE status = $%d",
				paramIndex)
		}

		args = append(args, filter.Status)
		paramIndex++
		hasWhere = true
	}
	requete += fmt.Sprintf(
		" LIMIT  $%d OFFSET $%d",
		paramIndex,
		paramIndex+1,
	)
	args = append(args, limit, offset)

	rows, err := database.DB.Query(requete, args...,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var listePharmacie []models.Pharmacie
	for rows.Next() {
		var parcourirListe models.Pharmacie
		err := rows.Scan(
			&parcourirListe.ID_PHARMACIE,
			&parcourirListe.Nom,
			&parcourirListe.Adresse,
			&parcourirListe.Telephone,
			&parcourirListe.Email,
			&parcourirListe.Ville,
			&parcourirListe.Status,
		)

		if err != nil {
			return nil, err
		}

		listePharmacie = append(listePharmacie, parcourirListe)
	}

	return listePharmacie, nil

}

func AfficherPharmacieDB(id_pharmacie int) (models.Pharmacie, error) {
	var afficherPharmacie models.Pharmacie

	// Requête pour afficher une pharmacie par id
	requete := `
			SELECT id_pharmacie,
			 nom,
			  adresse,
			   telephone,
			    email,
				 ville,
				 status FROM pharmacies WHERE id_pharmacie = $1
						`
	err := database.DB.QueryRow(requete, id_pharmacie).Scan(
		&afficherPharmacie.ID_PHARMACIE,
		&afficherPharmacie.Nom,
		&afficherPharmacie.Adresse,
		&afficherPharmacie.Telephone,
		&afficherPharmacie.Email,
		&afficherPharmacie.Ville,
		&afficherPharmacie.Status,
	)

	return afficherPharmacie, err

}

func ModifierPharmacieDB(id int, putPharmacie models.Pharmacie) error {
	// Requête pour modifier une pharmacie
	requete := `
					UPDATE pharmacies 
					SET nom = $1,
					adresse = $2,
					telephone = $3,
					email = $4,
					ville = $5
					WHERE id_pharmacie = $6

	`
	_, err := database.DB.Exec(requete,
		putPharmacie.Nom,
		putPharmacie.Adresse,
		putPharmacie.Telephone,
		putPharmacie.Email,
		putPharmacie.Ville,
		id)
	return err
}

func SupprimerPharmacieDB(id int) error {

	// Reuqête pour supprimer une pharmacie
	requete := `
				DELETE FROM pharmacies WHERE id_pharmacie = $1 
`

	_, err := database.DB.Exec(requete, id)
	return err

}
