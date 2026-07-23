package repositories

import (
	"pharmacie-api/database"
	"pharmacie-api/garde/models"
)

func AjouterGardeDB(newGarde models.Garde) error {
	requete := `
					INSERT INTO gardes (nom_garde, date_garde, heure_debut, heure_fin)
					VALUES ($1, $2, $3, $4)
	`
	_, err := database.DB.Exec(requete,
		newGarde.Nom_Garde,
		newGarde.DateGarde,
		newGarde.HeureDebut,
		newGarde.HeureFin)
	return err
}

func ListerGardesDB() ([]models.Garde, error) {
	requete := `
						SELECT id_garde, nom_garde, date_garde, heure_debut, heure_fin FROM gardes
		`
	rows, err := database.DB.Query(requete)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var listeGardes []models.Garde
	for rows.Next() {
		var parcourirGarde models.Garde
		err := rows.Scan(
			&parcourirGarde.ID_GARDE,
			&parcourirGarde.Nom_Garde,
			&parcourirGarde.DateGarde,
			&parcourirGarde.HeureDebut,
			&parcourirGarde.HeureFin)

		if err != nil {
			return nil, err
		}
		listeGardes = append(listeGardes, parcourirGarde)
	}
	return listeGardes, nil

}

func AfficherGardeDB(id_garde int) (models.Garde, error) {
	var obtenirGarde models.Garde
	requete := `
						SELECT id_garde, nom_garde, date_garde, heure_debut, heure_fin FROM gardes WHERE id_garde = $1
		`
	err := database.DB.QueryRow(requete, id_garde).Scan(
		&obtenirGarde.ID_GARDE,
		&obtenirGarde.Nom_Garde,
		&obtenirGarde.DateGarde,
		&obtenirGarde.HeureDebut,
		&obtenirGarde.HeureFin,
	)
	return obtenirGarde, err
}

func ModifierGardeDB(id_garde int, putGarde models.Garde) error {
	requete := `
					UPDATE gardes 
					SET nom_garde = $1,
					date_garde = $2,
					heure_debut = $3,
					heure_fin = $4
					WHERE id_garde = $5

			`
	_, err := database.DB.Exec(requete,
		putGarde.Nom_Garde,
		putGarde.DateGarde,
		putGarde.HeureDebut,
		putGarde.HeureFin,
		id_garde)
	return err

}

func SupprimerGardeDB(id_garde int) error {
	requete := `
						DELETE FROM gardes WHERE id_garde= $1 
	`
	_, err := database.DB.Exec(requete, id_garde)
	return err
}
