package repositories

import (
	"pharmacie-api/database"
	"pharmacie-api/pharmacie/models"
)

func AjouterPharmacieDB(newPharmacie models.Pharmacie) error {

	requete := `
					INSERT INTO pharmacies (nom, adresse, telephone, email, ville)
					VALUES ($1, $2, $3, $4, $5)	
	`
	_, err := database.DB.Exec(requete,
		newPharmacie.Nom,
		newPharmacie.Adresse,
		newPharmacie.Telephone,
		newPharmacie.Email,
		newPharmacie.Ville)

	return err

}

func ListerPharmacieDB() ([]models.Pharmacie, error) {
	requete := `
						SELECT id_pharmacie, nom, adresse, telephone, email, ville FROM pharmacies
		`
	rows, err := database.DB.Query(requete)
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

	requete := `
			SELECT id_pharmacie, nom, adresse, telephone, email, ville FROM pharmacies WHERE id_pharmacie = $1
						`
	err := database.DB.QueryRow(requete, id_pharmacie).Scan(
		&afficherPharmacie.ID_PHARMACIE,
		&afficherPharmacie.Nom,
		&afficherPharmacie.Adresse,
		&afficherPharmacie.Telephone,
		&afficherPharmacie.Email,
		&afficherPharmacie.Ville,
	)

	return afficherPharmacie, err

}

func ModifierPharmacieDB(id int, putPharmacie models.Pharmacie) error {

	requete := `
					UPDATE pharmacies 
					SET nom = $1,
					adresse = $2,
					telephone = $3,
					email = $4,
					ville = $5 
					WHERE id = $6

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
	requete := `
				DELETE FROM pharmacies WHERE id = $1 
`

	_, err := database.DB.Exec(requete, id)
	return err

}
