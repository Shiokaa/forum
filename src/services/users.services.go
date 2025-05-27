package services

import (
	"database/sql"
	"fmt"
	"forum/src/models"
	"forum/src/repositories"
)

// Structure permettant l'injection des repositories
type UsersServices struct {
	userRepositories *repositories.UsersRepositories
}

// Fonction initialisant l'injection de la base de donnée dans le repositorie de user
func UsersServicesInit(db *sql.DB) *UsersServices {
	return &UsersServices{userRepositories: repositories.UsersRepositoriesInit(db)}
}

// Fonction permettant la création d'un utilisateur
func (s *UsersServices) Create(user models.Users) (int, error) {
	// Vérification de la précense des données
	if user.Name == "" || user.Email == "" || user.Password == "" || user.Role_id < 0 {
		return -1, fmt.Errorf(" Erreur ajout produit - Données manquantes ou invalides")
	}

	// Envoie des données vers le repositorie
	userId, userErr := s.userRepositories.CreateUser(user)
	if userErr != nil {
		return -1, userErr
	}

	return userId, nil
}
