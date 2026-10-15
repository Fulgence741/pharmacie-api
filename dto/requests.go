package dto

// CreateUserRequest décrit les informations acceptées lors d'une inscription.
type CreateUserRequest struct {
	Nom      string `json:"nom"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Fonction string `json:"fonction"`
}

// LoginRequest contient les identifiants envoyés pour ouvrir une session.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UpdateUserRoleRequest limite la modification au rôle public de l'utilisateur.
type UpdateUserRoleRequest struct {
	Role string `json:"role"`
}

// PharmacieRequest contient les champs modifiables d'une pharmacie.
type PharmacieRequest struct {
	Nom       string `json:"nom"`
	Adresse   string `json:"adresse"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
	Ville     string `json:"ville"`
}

// GardeRequest contient les champs modifiables d'une garde.
type GardeRequest struct {
	NomGarde   string `json:"nom_garde"`
	DateGarde  string `json:"date_garde"`
	HeureDebut string `json:"heure_debut"`
	HeureFin   string `json:"heure_fin"`
}

// PharmacieGardeRequest porte les identifiants des deux ressources à relier.
type PharmacieGardeRequest struct {
	PharmacieID int `json:"id_pharmacie"`
	GardeID     int `json:"id_garde"`
}

// TokenResponse est la réponse renvoyée après une authentification réussie.
type TokenResponse struct {
	Token string `json:"token"`
}
