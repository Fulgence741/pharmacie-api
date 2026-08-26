package repositories

import (
	"pharmacie-api/database"
	"pharmacie-api/pharmacie_garde/models"
)

func AjouterDB(new models.PharmacieGarde) error {

	// Requête pour ajouter une nouvelle garde en base de donnée
	requete := `
							INSERT INTO pharmacie_garde (id_pharmacie, id_garde)
							VALUES ($1, $2)
			`
	_, err := database.DB.Exec(requete,
		new.PharmacieID,
		new.GardeID)
	return err

}
func ListerByGardeDB(limit int, offset int) ([]models.PharmacieGardeA, error) {
	// Requête pour lister toutes les pharmacies de gardes disponibles en base de donnée
	requete := `
					SELECT 
					p.nom,
					p.adresse,
					g.nom_garde,
					g.date_garde,
					g.heure_debut,
					g.heure_fin
					FROM pharmacie_garde pg 
					JOIN pharmacies p 
					ON pg.id_pharmacie = p.id_pharmacie
					JOIN gardes	g
					ON pg.id_garde = g.id_garde
					LIMIT $1 OFFSET $2
	`
	rows, err := database.DB.Query(requete, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pharmacieGardes := make([]models.PharmacieGardeA, 0)
	for rows.Next() {
		var pharmacieGarde models.PharmacieGardeA
		err := rows.Scan(
			&pharmacieGarde.Nom,
			&pharmacieGarde.Adresse,
			&pharmacieGarde.NomGarde,
			&pharmacieGarde.DateGarde,
			&pharmacieGarde.HeureDebut,
			&pharmacieGarde.HeureFin)

		if err != nil {
			return nil, err
		}

		pharmacieGardes = append(pharmacieGardes, pharmacieGarde)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pharmacieGardes, nil
}

func SupprimerDb(id int) error {
	requete := `
				DELETE FROM pharmacie_garde 
				WHERE id = $1
	`
	_, err := database.DB.Exec(requete, id)
	return err

}
