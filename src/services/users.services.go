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
		return -1, fmt.Errorf(" Erreur ajout user - Données manquantes ou invalides")
	}

	// Envoie des données vers le repositorie
	userId, userErr := s.userRepositories.CreateUser(user)
	if userErr != nil {
		return -1, userErr
	}

	return userId, nil
}

func (s *UsersServices) Connect(email string, password string) (models.Users, error) {

	if email == "" || password == "" {
		return models.Users{}, fmt.Errorf(" Erreur connection - Données manquantes ou invalides")
	}

	user, err := s.userRepositories.ConnectUser(email, password)
	if err != nil {
		return models.Users{}, err
	}

	return user, nil
}

func (s *UsersServices) ReadId(idUser int) (models.Users, error) {
	if idUser < 1 {
		return models.Users{}, fmt.Errorf(" Erreur récupération du topic - identifiant invalide : %d", idUser)
	}

	topic, errTopic := s.userRepositories.GetUserById(idUser)
	if errTopic != nil {
		return models.Users{}, errTopic
	}

	return topic, nil
}

// GetAll récupère la liste de tous les utilisateurs.
func (s *UsersServices) GetAll() ([]models.Users, error) {
	return s.userRepositories.GetAllUsers()
}

// DeleteUser demande la suppression d'un utilisateur.
func (s *UsersServices) DeleteUser(id int) error {
	return s.userRepositories.DeleteUser(id)
}
