package errors

import (
	"errors"
)

// les erreurs de la base de donnée
var (
	ErrDatabaseConnection = errors.New("Erreur de connection à la base de donnée")
	ErrDatabseQuery       = errors.New("Erreur lors de la requête")
)
