package services

import (
	"database/sql"
	"fmt"
	"forum/src/models"
	"forum/src/repositories"
)

// Structure permettant l'injection des repositories
type TopicsServices struct {
	topicsRepositories *repositories.TopicsRepositories
}

// Fonction initialisant l'injection de la base de donnée dans le repositorie de user
func TopicsServicesInit(db *sql.DB) *TopicsServices {
	return &TopicsServices{topicsRepositories: repositories.TopicsRepositoriesInit(db)}
}

// Fonction permettant de récupérer les valeurs et de les renvoyer
func (s *TopicsServices) Display() ([]models.Topics_Join_Users, error) {

	// Récupération des données avec gestion d'erreur
	items, err := s.topicsRepositories.GetTopicsWithCreators()
	if err != nil {
		return items, err
	}

	return items, nil
}

// Fonction permettant de parcourir un topic via l'id, ici nous récupérons aussi l'utilisateur et le forum du topic
func (s *TopicsServices) ReadId(idTopic int) (models.Topics_Join_Users_Forums, error) {

	if idTopic < 1 {
		return models.Topics_Join_Users_Forums{}, fmt.Errorf(" Erreur récupération du topic - identifiant invalide : %d", idTopic)
	}

	topic, errTopic := s.topicsRepositories.GetTopicWithId(idTopic)
	if errTopic != nil {
		return models.Topics_Join_Users_Forums{}, errTopic
	}

	return topic, nil
}

// FOnction permettant de récupérer les messages lié à un certain topic en passsant par l'id
func (s *TopicsServices) ReadMessages(idTopic int) ([]models.Topics_Join_Messages, error) {
	// Récupération des données avec gestion d'erreur
	items, err := s.topicsRepositories.GetTopicWithMessage(idTopic)
	if err != nil {
		return items, err
	}

	return items, nil
}
