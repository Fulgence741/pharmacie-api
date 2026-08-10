package repositories

import (
	"errors"
	"pharmacie-api/database"
	"pharmacie-api/users/models"
)

func AjouterUserDB(newUser models.User) error {
	requete := `
				INSERT INTO users (nom, email, password, fonction, role)
				VALUES ($1, $2, $3, $4, $5)
	`
	_, err := database.DB.Exec(requete,
		newUser.Nom,
		newUser.Email,
		newUser.Password,
		newUser.Fonction,
		newUser.Role)
	return err
}

func ListerUserDB() ([]models.User, error) {
	requet := `
				SELECT id, nom, email, fonction, role
				FROM users
	`
	rows, err := database.DB.Query(requet)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var listeUser []models.User
	for rows.Next() {
		var parcourirUser models.User
		err := rows.Scan(
			&parcourirUser.ID_USER,
			&parcourirUser.Nom,
			&parcourirUser.Email,
			&parcourirUser.Fonction,
			&parcourirUser.Role)

		if err != nil {
			return nil, err
		}

		listeUser = append(listeUser, parcourirUser)
	}

	return listeUser, nil
}

func CnnexionUserDB(email string) (models.User, error) {
	var user models.User

	requete := `
				SELECT id, nom, email, password, fonction, role 
				FROM users
				WHERE email = $1
	`

	err := database.DB.QueryRow(requete, email).Scan(
		&user.ID_USER,
		&user.Nom,
		&user.Email,
		&user.Password,
		&user.Fonction,
		&user.Role)

	if err != nil {
		return models.User{}, err
	}

	return user, nil

}

func SupprimerUserDB(id int) error {

	requete := `
				DELETE FROM users WHERE id = $1
	`
	_, err := database.DB.Exec(requete, id)
	return err
}

func ModifierRoleDB(id int, role string) error {
	requete := `
				UPDATE users
					SET role = $1
					WHERE id = $2
	`
	resultat, err := database.DB.Exec(requete, role, id)
	if err != nil {
		return err
	}

	nombre, err := resultat.RowsAffected()
	if err != nil {
		return err
	}

	if nombre == 0 {
		return errors.New("Utilisateur introuvable")
	}

	return nil
}
