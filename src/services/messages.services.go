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
