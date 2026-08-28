package services

import (
	"errors"
	"pharmacie-api/users/models"
	"pharmacie-api/users/repositories"
	"pharmacie-api/utils"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func AjouterUserService(user models.User) error {
	if user.Nom == "" {
		return errors.New("Le champ nom est obligatoire")
	}
	if user.Email == "" {
		return errors.New("Le champ Email est obligatoire")
	}
	if !strings.Contains(user.Email, "@") {
		return errors.New("Email invalide")
	}
	if user.Password == "" {
		return errors.New("Le champ password est obligatoire")
	}
	if len(user.Password) < 8 {
		return errors.New("Le mot de passe doit conteenir au moins 8 caractères")
	}
	if user.Fonction == "" {
		return errors.New("Le champ fonction est obligatoire")
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
	return repositories.AjouterUserDB(user)
}

func ListerUserService(
	pagination utils.Pagination,
	filter models.UserFilter) ([]models.User, error) {
	user, err := repositories.ListerUserDB(pagination.Limit,
		pagination.Offset,
		filter)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func Connexion(email string, password string) (models.User, error) {
	user, err := repositories.CnnexionUserDB(email)

	if err != nil {
		return models.User{}, errors.New("Email incorrect")
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return models.User{}, errors.New("Mot de passe incorrect")
	}
	return user, nil
}

func SupprimerUserService(id int) error {
	if id <= 0 {
		return errors.New("Id invalide")
	}

	err := repositories.SupprimerUserDB(id)
	if err != nil {
		return err
	}

	return nil

}

func ModifierRoleServices(id int, role string) error {
	if id <= 0 {
		return errors.New("Utilisateur introuvable")
	}

	switch role {
	case "pharmacien", "user":
	default:
		return errors.New("Rôle invalide")
	}

	return repositories.ModifierRoleDB(id, role)
}

func AfficherUserServices(id int) (models.User, error) {

	if id <= 0 {
		return models.User{}, errors.New("Identifiant invalide")
	}

	user, err := repositories.AfficherUserDB(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil

}

/*
// Fonction de modification de mot de passe à implémenter plus tard   !!
======================================================================
func ChangerPasswordServices(
	userID int,
	changement models.ChangerPassword,
) error {
	passwordHash, err := repositories.ObtenirPasswordDB(userID)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(passwordHash),
		[]byte(changement.AncienPassword),
	)
	if err != nil {
		return err
	}

	if len(changement.NouveauPassword) < 8 {
		return errors.New("Le nouveau mot de passe doit contenir au moins 8 caractères")
	}

	nouveauPasswordHash, err := bcrypt.GenerateFromPassword(
		[]byte(changement.NouveauPassword),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	return nil
} */
