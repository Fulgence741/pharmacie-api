package repositories

import (
	"context"
	"fmt"
	"pharmacie-api/database"
	apierrors "pharmacie-api/errors"
	"pharmacie-api/models"
)

// Le ctx reçu est celui de la requête HTTP : ne pas le remplacer par Background,
// afin de laisser PostgreSQL arrêter le travail si le client abandonne la requête.
func AjouterUserDB(ctx context.Context, newUser models.User) error {
	requete := `
				INSERT INTO users (nom, email, password, fonction, role)
				VALUES ($1, $2, $3, $4, $5)
	`
	_, err := database.DB.Exec(ctx, requete,
		newUser.Nom,
		newUser.Email,
		newUser.Password,
		newUser.Fonction,
		newUser.Role)
	return err
}

func ListerUserDB(ctx context.Context,
	limit int,
	offset int,
	filter models.UserFilter) ([]models.User, error) {
	requet := `
				SELECT id, nom, email, fonction, role
				FROM users
	`
	var args []interface{}
	paramIndex := 1
	if filter.Nom != "" {
		requet += fmt.Sprintf(" WHERE nom ILIKE $%d", paramIndex)
		args = append(args, "%"+filter.Nom+"%")
		paramIndex++
	}

	requet += fmt.Sprintf(
		"LIMIT $%d OFFSET $%d",
		paramIndex,
		paramIndex+1,
	)

	args = append(args, limit, offset)
	rows, err := database.DB.Query(ctx,
		requet, args...)
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

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return listeUser, nil
}

func CnnexionUserDB(ctx context.Context, email string) (models.User, error) {
	var user models.User

	requete := `
				SELECT id, nom, email, password, fonction, role 
				FROM users
				WHERE email = $1
	`

	err := database.DB.QueryRow(ctx, requete, email).Scan(
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

func SupprimerUserDB(ctx context.Context, id int) error {

	requete := `
				DELETE FROM users WHERE id = $1
	`
	_, err := database.DB.Exec(ctx,
		requete, id)
	return err
}

func ModifierRoleDB(ctx context.Context, id int, role string) error {
	requete := `
				UPDATE users
					SET role = $1
					WHERE id = $2
	`
	resultat, err := database.DB.Exec(ctx,
		requete, role, id)
	if err != nil {
		return err
	}

	nombre := resultat.RowsAffected()

	if nombre == 0 {
		return apierrors.NotFound("Utilisateur introuvable")
	}

	return nil
}

func AfficherUserDB(ctx context.Context, id int) (models.User, error) {
	var afficherU models.User
	requete := `
				SELECT id,
				nom,
				email,
				fonction, 
				role FROM users WHERE id = $1
	`

	err := database.DB.QueryRow(ctx,
		requete, id).Scan(
		&afficherU.ID_USER,
		&afficherU.Nom,
		&afficherU.Email,
		&afficherU.Fonction,
		&afficherU.Role,
	)

	return afficherU, err
}

func ObtenirPasswordDB(ctx context.Context, userID int) (string, error) {
	requete := `
					SELECT password 
					FROM users
					WHERE id = $1
	`
	var passwordHash string

	err := database.DB.QueryRow(ctx,
		requete, userID).Scan(&passwordHash)
	if err != nil {
		return "", err
	}

	return passwordHash, nil
}
