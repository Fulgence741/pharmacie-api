package handlers

import (
	"pharmacie-api/dto"
	"pharmacie-api/models"
)

func userResponse(user models.User) dto.UserResponse {
	return dto.UserResponse{
		ID:       user.ID_USER,
		Nom:      user.Nom,
		Email:    user.Email,
		Fonction: user.Fonction,
		Role:     user.Role,
	}
}

func userResponses(users []models.User) []dto.UserResponse {
	result := make([]dto.UserResponse, 0, len(users))
	for _, user := range users {
		result = append(result, userResponse(user))
	}
	return result
}

func pharmacieResponse(pharmacie models.Pharmacie) dto.PharmacieResponse {
	return dto.PharmacieResponse{
		IDPharmacie: pharmacie.ID_PHARMACIE,
		Nom:         pharmacie.Nom,
		Adresse:     pharmacie.Adresse,
		Telephone:   pharmacie.Telephone,
		Email:       pharmacie.Email,
		Ville:       pharmacie.Ville,
		Status:      pharmacie.Status,
	}
}

func pharmacieResponses(pharmacies []models.Pharmacie) []dto.PharmacieResponse {
	result := make([]dto.PharmacieResponse, 0, len(pharmacies))
	for _, pharmacie := range pharmacies {
		result = append(result, pharmacieResponse(pharmacie))
	}
	return result
}

func gardeResponse(garde models.Garde) dto.GardeResponse {
	return dto.GardeResponse{
		IDGarde:    garde.ID_GARDE,
		NomGarde:   garde.Nom_Garde,
		DateGarde:  garde.DateGarde,
		HeureDebut: garde.HeureDebut,
		HeureFin:   garde.HeureFin,
	}
}

func gardeResponses(gardes []models.Garde) []dto.GardeResponse {
	result := make([]dto.GardeResponse, 0, len(gardes))
	for _, garde := range gardes {
		result = append(result, gardeResponse(garde))
	}
	return result
}

func pharmacieGardeResponses(pharmacieGardes []models.PharmacieGardeA) []dto.PharmacieGardeResponse {
	result := make([]dto.PharmacieGardeResponse, 0, len(pharmacieGardes))
	for _, pharmacieGarde := range pharmacieGardes {
		result = append(result, dto.PharmacieGardeResponse{
			Nom:        pharmacieGarde.Nom,
			Adresse:    pharmacieGarde.Adresse,
			NomGarde:   pharmacieGarde.NomGarde,
			DateGarde:  pharmacieGarde.DateGarde,
			HeureDebut: pharmacieGarde.HeureDebut,
			HeureFin:   pharmacieGarde.HeureFin,
		})
	}
	return result
}
