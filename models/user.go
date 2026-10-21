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
