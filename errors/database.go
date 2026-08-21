package errors

import (
	"errors"
)

var (
	ErrDatabaseConnection = errors.New("Erreur de connection à a")
	ErrDatabseQuery       = errors.New("Erreur lors de la requête")
)
