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
}

// Fonction initialisant l'injection de la base de donnée dans le repositorie de user
func MessagesServicesInit(db *sql.DB) *MessagesServices {
	return &MessagesServices{MessageRepositories: repositories.MessageRepositoriesInit(db)}
}

// Fonction permettant de parcourir un message via l'id, ici nous récupérons aussi l'utilisateur et le titre du topic
func (s *MessagesServices) ReadId(idMessage int) (models.Topics_Join_Messages, error) {

	if idMessage < 1 {
		return models.Topics_Join_Messages{}, fmt.Errorf(" Erreur récupération du message - identifiant invalide : %d", idMessage)
	}

	messsage, errMessage := s.MessageRepositories.GetMessageById(idMessage)
	if errMessage != nil {
		return models.Topics_Join_Messages{}, errMessage
	}

	return messsage, nil
}
