package services

import (
	"errors"
	"pharmacie-api/users/models"
	"pharmacie-api/users/repositories"

	"golang.org/x/crypto/bcrypt"
)

func AjouterUserService(user models.User) error {
	if user.Nom == "" {
		return errors.New("Le champ nom est obligatoire")
	}
	if user.Email == "" {
		return errors.New("Le champ Email est obligatoire")
	}
	if user.Password == "" {
		return errors.New("Le champ password est obligatoire")
	}
	if user.Fonction == "" {
		return errors.New("Le champ fonction est obligatoire")
	}
	if user.Role == "" {
		return errors.New("Le champ role est obligaatoire")
	}

	motDePasseHache, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user.Password = string(motDePasseHache)
	return repositories.AjouterUserDB(user)
}

func ListerUserService() ([]models.User, error) {
	user, err := repositories.ListerUserDB()
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
