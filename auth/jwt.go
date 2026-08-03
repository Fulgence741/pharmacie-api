package auth

import (
	"os"
	"pharmacie-api/users/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SecretKey() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET manquant")
	}
	return []byte(secret)
}

func GenererJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    user.ID_USER,
			"email": user.Email,
			"role":  user.Role,
			"exp":   time.Now().Add(1 * time.Hour).Unix(),
		},
	)

	tokenString, err := token.SignedString(SecretKey())

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
