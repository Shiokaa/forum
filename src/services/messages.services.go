package services

import (
	"database/sql"
	"fmt"
	"forum/src/models"
	"forum/src/repositories"
)

// Structure permettant l'injection des repositories
type MessagesServices struct {
	MessageRepositories *repositories.MessagesRepositories
	RepliesRepositories *repositories.RepliesRepositories
}

// Fonction initialisant l'injection de la base de donnée dans le repositorie de user
func MessagesServicesInit(db *sql.DB) *MessagesServices {
	return &MessagesServices{MessageRepositories: repositories.MessageRepositoriesInit(db), RepliesRepositories: repositories.RepliesRepositoriesInit(db)}
}

// Fonction permettant de parcourir un message via l'id, ici nous récupérons aussi l'utilisateur et le titre du topic
func (s *MessagesServices) ReadMessagesId(idMessage int) (models.Topics_Join_Messages, error) {

	if idMessage < 1 {
		return models.Topics_Join_Messages{}, fmt.Errorf(" Erreur récupération du message - identifiant invalide : %d", idMessage)
	}

	messsage, errMessage := s.MessageRepositories.GetMessageById(idMessage)
	if errMessage != nil {
		return models.Topics_Join_Messages{}, errMessage
	}

	return messsage, nil
}

// Fonction permettant de parcourir un message via l'id, ici nous récupérons aussi l'utilisateur et le titre du topic
func (s *MessagesServices) ReadRepliesId(idReply int) ([]models.Replies_Join_User, error) {

	if idReply < 1 {
		return []models.Replies_Join_User{}, fmt.Errorf(" Erreur récupération du message - identifiant invalide : %d", idReply)
	}

	replies, errReplies := s.RepliesRepositories.GetReplies(idReply)
	if errReplies != nil {
		return []models.Replies_Join_User{}, errReplies
	}

	return replies, nil
}

func (s *MessagesServices) CreatedReply(reply models.Replies_Joins_User_Message) (int, error) {
	// Vérification de la précense des données
	if reply.Users.User_id < 0 || reply.Messages.Message_id < 0 || reply.Replies.Content == "" {
		return -1, fmt.Errorf(" Erreur ajout reponse - Données manquantes ou invalides")
	}

	// Envoie des données vers le repositorie
	userId, userErr := s.MessageRepositories.PostReplie(reply)
	if userErr != nil {
		return -1, userErr
	}

	return userId, nil
}

// DisplayRecent transmet l'appel pour récupérer les messages les plus récents.
func (s *MessagesServices) DisplayRecent(limit int) ([]models.Topics_Join_Messages, error) {
	if limit <= 0 {
		return nil, fmt.Errorf("la limite doit être un nombre positif")
	}

	items, err := s.MessageRepositories.GetRecentMessages(limit)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *MessagesServices) DeleteMessage(id int) error {
	return s.MessageRepositories.DeleteMessage(id)
}

// ReadReplyByID récupère les détails d'une réponse.
func (s *MessagesServices) ReadReplyByID(id int) (models.Replies, error) {
	return s.RepliesRepositories.GetReplyByID(id)
}

// DeleteReply demande la suppression d'une réponse.
func (s *MessagesServices) DeleteReply(id int) error {
	return s.RepliesRepositories.DeleteReply(id)
}
