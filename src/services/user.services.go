package services

import (
	"database/sql"
	"fmt"
	"forum/src/models"
	"forum/src/repositories"
)

// Structure permettant l'injection des repositories
type UserServices struct {
	userRepositories *repositories.UserRepositories
}

// Fonction initialisant l'injection de la base de donnée dans le repositorie de user
func UserServicesInit(db *sql.DB) *UserServices {
	return &UserServices{userRepositories: repositories.UserRepositoriesInit(db)}
}

// Fonction permettant la création d'un utilisateur
func (s *UserServices) Create(user models.User) (int, error) {
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
