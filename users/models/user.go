package models

// Objet pour les colonnes disponibles en base de données

type User struct {
	ID_USER  int    `json:"id"`
	Nom      string `json:"nom"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Fonction string `json:"fonction"`
	Role     string `json:"role"`
}

// Objet pour filtrer les utilisateur par nom

type UserFilter struct {
	Nom string
}

// Objet pour changer le mot de passe par un utilisateur connecté

type ChangerPassword struct {
	AncienPassword  string `json:"ancien_password"`
	NouveauPassword string `json:"nouveau_password"`
}
