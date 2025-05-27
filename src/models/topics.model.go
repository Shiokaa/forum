package models

// Modele de topics identique à la structure de la base de données
type Topics struct {
	Topic_id   int
	Forum_id   int
	User_Id    int
	Title      string
	Status     bool
	Created_at string
	Updated_at string
}
