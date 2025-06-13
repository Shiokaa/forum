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

func (s *MessagesServices) CreateMessage(message models.Messages) (int, error) {
	// Vérification de la présence des données
	if message.Topic_id < 1 || message.User_id < 1 || message.Content == "" {
		return -1, fmt.Errorf(" erreur lors de l'ajout du message - Données manquantes ou invalides")
	}

	// Envoi des données vers le repository
	messageId, err := s.MessageRepositories.CreateMessage(message)
	if err != nil {
		return -1, err
	}

	return messageId, nil
}

// GetTotalTopicsCount transmet l'appel pour compter le nombre total de topics.
func (s *TopicsServices) GetTotalTopicsCount() (int, error) {
	count, err := s.topicsRepositories.GetTotalTopicsCount()
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetPaginatedTopics transmet l'appel pour récupérer les topics de manière paginée.
func (s *TopicsServices) GetPaginatedTopics(offset, limit int) ([]models.Topics_Join_Users, error) {
	items, err := s.topicsRepositories.GetPaginatedTopics(offset, limit)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *TopicsServices) GetByCategoryID(categoryID int, limit int) ([]models.Topics_Join_Users, error) {
	return s.topicsRepositories.GetByCategoryID(categoryID, limit)
}

func (s *TopicsServices) GetByForumIDPaginated(forumID, offset, limit int) ([]models.Topics_Join_Users, error) {
	return s.topicsRepositories.GetByForumIDPaginated(forumID, offset, limit)
}

func (s *TopicsServices) GetTotalTopicsCountByForumID(forumID int) (int, error) {
	return s.topicsRepositories.GetTotalTopicsCountByForumID(forumID)
}

// CreateTopic valide les données et demande la création d'un nouveau topic.
func (s *TopicsServices) CreateTopic(topic models.Topics) (int, error) {
	if topic.Title == "" || topic.Forum_id < 1 || topic.User_Id < 1 {
		return -1, fmt.Errorf("données du topic invalides ou manquantes")
	}

	topicID, err := s.topicsRepositories.CreateTopic(topic)
	if err != nil {
		return -1, err
	}

	return topicID, nil
}
