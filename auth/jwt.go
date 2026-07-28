package auth

import (
	"pharmacie-api/users/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("ma-cle-secret")

func GenererJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":     user.ID_USER,
			"email:": user.Email,
			"role":   user.Role,
			"exp":    time.Now().Add(1 * time.Hour).Unix(),
		},
	)

	tokenString, err := token.SignedString(SecretKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
