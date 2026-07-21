package models

type Pharmacie struct {
	ID_PHARMACIE int    `json:"id_pharmacie"`
	Nom          string `json:"nom"`
	Adresse      string `json:"adresse"`
	Telephone    string `json:"telephone"`
	Email        string `json:"email"`
	Ville        string `json:"ville"`
}
