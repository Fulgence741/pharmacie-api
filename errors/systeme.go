package errors

import "errors"

var (
	ErrInternalServer = errors.New("erreur interne du serveur")
	ErrInvalidInput   = errors.New("données invalides")
	ErrUnauthorized   = errors.New("accès non autorisé")
	ErrForbidden      = errors.New("accès interdit")
	ErrNotFound       = errors.New("ressource introuvable")
)
