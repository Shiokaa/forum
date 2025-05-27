package models

// Modele de user identique à la structure de la base de données
type Users struct {
	User_Id    int
	Role_id    int
	Name       string
	Email      string
	Password   string
	Created_at string
	Updated_at string
}
