package dto

import "time"

// UserResponse expose les informations publiques d'un utilisateur, sans son mot de passe.
type UserResponse struct {
	ID       int    `json:"id"`
	Nom      string `json:"nom"`
	Email    string `json:"email"`
	Fonction string `json:"fonction"`
	Role     string `json:"role"`
}

// PharmacieResponse décrit les champs de pharmacie renvoyés par l'API.
type PharmacieResponse struct {
	IDPharmacie int    `json:"id_pharmacie"`
	Nom         string `json:"nom"`
	Adresse     string `json:"adresse"`
	Telephone   string `json:"telephone"`
	Email       string `json:"email"`
	Ville       string `json:"ville"`
	Status      string `json:"status"`
}

// GardeResponse décrit les champs de garde renvoyés par l'API.
type GardeResponse struct {
	IDGarde    int    `json:"id_garde"`
	NomGarde   string `json:"nom_garde"`
	DateGarde  string `json:"date_garde"`
	HeureDebut string `json:"heure_debut"`
	HeureFin   string `json:"heure_fin"`
}

// PharmacieGardeResponse représente une pharmacie associée à sa période de garde.
type PharmacieGardeResponse struct {
	Nom        string    `json:"nom"`
	Adresse    string    `json:"adresse"`
	NomGarde   string    `json:"nom_garde"`
	DateGarde  time.Time `json:"date_garde"`
	HeureDebut string    `json:"heure_debut"`
	HeureFin   string    `json:"heure_fin"`
}
