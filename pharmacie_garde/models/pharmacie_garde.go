package models

// Objet pour les colonnes disponibles en base de donnée
type PharmacieGarde struct {
	ID          int `json:"id"`
	PharmacieID int `json:"id_pharmacie"`
	GardeID     int `json:"id_garde"`
}

// Objet combinant les informations d'une pharmacie et d'une garde à renvoyer au client lors de la demande d'une pharmacie de garde

type PharmacieGardeA struct {
	Nom        string `json:"nom"`
	Adresse    string `json:"adresse"`
	NomGarde   string `json:"nom_garde"`
	DateGarde  string `json:"date_garde"`
	HeureDebut string `json:"heure_debut"`
	HeureFin   string `json:"heure_fin"`
}

// Objet pour filtrer les recherche de pharmacie de garde

type PharmacieGardeFliter struct {
	Nom string
}
