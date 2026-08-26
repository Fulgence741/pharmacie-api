package errors

import (
	"errors"
)

var (
	ErrDatabaseConnection = errors.New("Erreur de connection à la base de donnée")
	ErrDatabseQuery       = errors.New("Erreur lors de la requête")
)
