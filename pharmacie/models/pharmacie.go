package models

// Objet pour les colonnes se trouvant en base de donée
type Pharmacie struct {
	ID_PHARMACIE int    `json:"id_pharmacie"`
	Nom          string `json:"nom"`
	Adresse      string `json:"adresse"`
	Telephone    string `json:"telephone"`
	Email        string `json:"email"`
	Ville        string `json:"ville"`
	Status       string `json:"status"`
}

// Objet pour filtrer une recherche par Nom , Ville et status d'une pharmacie

type PharmacieFilter struct {
	Nom    string
	Ville  string
	Status string
}
