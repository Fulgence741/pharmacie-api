package auth

import (
	"os"
	"pharmacie-api/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// SecretKey retourne la clé secrète utilisée pour signer et vérifier les tokens JWT. Elle est lue depuis la variable d'environnement JWT_SECRET.
func SecretKey() []byte {
	// Récupération de la clé secrète depuis la variable d'environnement JWT_SECRET
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// Si la variable d'environnement n'est pas définie, on panique pour signaler une erreur critique
		panic("JWT_SECRET manquant")
	}
	// Retourne la clé secrète sous forme de slice d'octets
	return []byte(secret)
}

// GenererJWT génère un token JWT pour un utilisateur donné. Le token contient l'ID, l'email, le rôle de l'utilisateur et une date d'expiration.

func GenererJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims( // Création d'un nouveau token JWT avec les claims spécifiés
		jwt.SigningMethodHS256, // Méthode de signature HMAC SHA256
		jwt.MapClaims{ // Claims du token (informations contenues dans le token)
			"id":    user.ID_USER,
			"email": user.Email,
			"role":  user.Role,
			"exp":   time.Now().Add(1 * time.Hour).Unix(), // Date d'expiration du token (1 heure à partir de maintenant)
		},
	)

	// Signature du token avec la clé secrète
	tokenString, err := token.SignedString(SecretKey())

	// Vérification des erreurs lors de la signature du token
	if err != nil {
		return "", err
	}

	// Retourne le token JWT signé sous forme de chaîne de caractères
	return tokenString, nil
}
