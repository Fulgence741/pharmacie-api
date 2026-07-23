package models

type PharmacieGarde struct {
	ID          int `json:"id"`
	PharmacieID int `json:"id_pharmacie"`
	GardeID     int `json:"id_garde"`
}

type PharmacieGardeA struct {
	Nom        string `json:"nom"`
	Adresse    string `json:"adresse"`
	NomGarde   string `json:"nom_garde"`
	DateGarde  string `json:"date_garde"`
	HeureDebut string `json:"heure_debut"`
	HeureFin   string `json:"heure_fin"`
}
