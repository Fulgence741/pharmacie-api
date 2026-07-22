package models

type Garde struct {
	ID_GARDE   int    `json:"id_garde"`
	Nom_Garde  string `json:"nom_garde"`
	DateGarde  string `json:"date_garde"`
	HeureDebut string `json:"heure_debut"`
	HeureFin   string `json:"heure_fin"`
}
