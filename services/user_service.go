package services

import (
	"context"
	"database/sql"
	stderrors "errors"
	"fmt"
	"log"
	"pharmacie-api/cache"
	"pharmacie-api/errors"
	"pharmacie-api/models"
	"pharmacie-api/repositories"
	"pharmacie-api/utils"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Les services reçoivent le contexte du handler et le transmettent aux repositories
// pour que les requêtes SQL respectent l'annulation et le délai de la requête HTTP.
func AjouterUserService(ctx context.Context, user models.User) error {
	if user.Nom == "" {
		return errors.InvalidRequest("Le champ nom est obligatoire")
	}
	if user.Email == "" {
		return errors.InvalidRequest("Le champ email est obligatoire")
	}
	if !strings.Contains(user.Email, "@") {
		return errors.InvalidRequest("Email invalide")
	}
	if user.Password == "" {
		return errors.InvalidRequest("Le champ password est obligatoire")
	}
	if len(user.Password) < 8 {
		return errors.InvalidRequest("Le mot de passe doit contenir au moins 8 caractères")
	}
	if user.Fonction == "" {
		return errors.InvalidRequest("Le champ fonction est obligatoire")
	}

	motDePasseHache, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user.Password = string(motDePasseHache)
	user.Role = "user"
	return repositories.AjouterUserDB(ctx, user)
}

func ListerUserService(ctx context.Context,
	pagination utils.Pagination,
	filter models.UserFilter) ([]models.User, error) {

	// Gestion de cache

	cacheKey := fmt.Sprintf(
		"cache:user:nom:%q:limit:%d:offset:%d",
		strings.ToLower(strings.TrimSpace(filter.Nom)),
		pagination.Limit,
		pagination.Offset,
	)

	var user []models.User

	found, err := cache.GetJSON(ctx, cacheKey, &user)
	if err == nil && found {
		log.Print("Cache HIT ressource=utilisateur backend=Redis")
		return user, nil
	}

	if err != nil {
		log.Printf("Cache READ_ERROR ressource=utilisateur backend=Redis erreur=%v", err)
	}

	log.Print("Cache MISS ressource=utilisateur backend=PostgreSQL")
	user, err = repositories.ListerUserDB(ctx, pagination.Limit,
		pagination.Offset,
		filter)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func Connexion(ctx context.Context, email string, password string) (models.User, error) {
	user, err := repositories.CnnexionUserDB(ctx, email)

	if err != nil {
		if stderrors.Is(err, sql.ErrNoRows) {
			return models.User{}, errors.Unauthorized("Email ou mot de passe incorrect")
		}
		return models.User{}, err
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		if stderrors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return models.User{}, errors.Unauthorized("Email ou mot de passe incorrect")
		}
		return models.User{}, err
	}
	return user, nil
}

func SupprimerUserService(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.InvalidRequest("Identifiant invalide")
	}

	err := repositories.SupprimerUserDB(ctx, id)
	if err != nil {
		return err
	}

	return nil

}

func ModifierRoleServices(ctx context.Context, id int, role string) error {
	if id <= 0 {
		return errors.InvalidRequest("Identifiant invalide")
	}

	switch role {
	case "pharmacien", "user":
	default:
		return errors.InvalidRequest("Rôle invalide")
	}

	return repositories.ModifierRoleDB(ctx, id, role)
}

func AfficherUserServices(ctx context.Context, id int) (models.User, error) {

	if id <= 0 {
		return models.User{}, errors.InvalidRequest("Identifiant invalide")
	}

	user, err := repositories.AfficherUserDB(ctx, id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil

}
