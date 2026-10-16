package handlers

import (
	"pharmacie-api/dto"
	"pharmacie-api/models"
)

func userResponse(utilisateur models.User) dto.UserResponse {
	return dto.UserResponse{
		ID:       utilisateur.ID_USER,
		Nom:      utilisateur.Nom,
		Email:    utilisateur.Email,
		Fonction: utilisateur.Fonction,
		Role:     utilisateur.Role,
	}
}

func userResponses(utilisateurs []models.User) []dto.UserResponse {
	resultat := make([]dto.UserResponse, 0, len(utilisateurs))
	for _, utilisateur := range utilisateurs {
		resultat = append(resultat, userResponse(utilisateur))
	}
	return resultat
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
	resultat := make([]dto.PharmacieResponse, 0, len(pharmacies))
	for _, pharmacie := range pharmacies {
		resultat = append(resultat, pharmacieResponse(pharmacie))
	}
	return resultat
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
	resultat := make([]dto.GardeResponse, 0, len(gardes))
	for _, garde := range gardes {
		resultat = append(resultat, gardeResponse(garde))
	}
	return resultat
}

func pharmacieGardeResponses(pharmaciesDeGarde []models.PharmacieGardeA) []dto.PharmacieGardeResponse {
	resultat := make([]dto.PharmacieGardeResponse, 0, len(pharmaciesDeGarde))
	for _, pharmacieDeGarde := range pharmaciesDeGarde {
		resultat = append(resultat, dto.PharmacieGardeResponse{
			Nom:        pharmacieDeGarde.Nom,
			Adresse:    pharmacieDeGarde.Adresse,
			NomGarde:   pharmacieDeGarde.NomGarde,
			DateGarde:  pharmacieDeGarde.DateGarde,
			HeureDebut: pharmacieDeGarde.HeureDebut,
			HeureFin:   pharmacieDeGarde.HeureFin,
		})
	}
	return resultat
}
