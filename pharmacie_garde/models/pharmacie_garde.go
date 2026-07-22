package models

type PharmacieGarde struct {
	ID          int `json:"id"`
	PharmacieID int `json:"id_pharmacie"`
	GardeID     int `json:"id_garde"`
}
