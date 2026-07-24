package models

type User struct {
	ID_USER  int    `json:"id"`
	Nom      string `json:"nom"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Fonction string `json:"fonction"`
	Role     string `json:"role"`
}
