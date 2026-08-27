package errors

import "errors"

// Erreurs du sytème go
var (
	ErrInternalServer = errors.New("erreur interne du serveur")
	ErrUnauthorized   = errors.New("accès non autorisé")
	ErrForbidden      = errors.New("accès interdit")
	ErrNotFound       = errors.New("ressource introuvable")
)
