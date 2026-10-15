package errors

import "errors"

// Erreurs du sytème go
var (
	ErrInternalServer = errors.New("erreur interne du serveur")
	ErrNotFound       = errors.New("ressource introuvable")
)
