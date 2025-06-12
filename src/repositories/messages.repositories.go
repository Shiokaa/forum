package repositories

import (
	"database/sql"
	"fmt"
	"forum/src/models"
)

// Structure permettant l'injection de la base de donnée
type MessagesRepositories struct {
	db *sql.DB
}

// Fonction pour initialiser le repositorie de user avec l'injection de la base de donnée
func MessageRepositoriesInit(db *sql.DB) *MessagesRepositories {
	return &MessagesRepositories{db: db}
}

func (r *MessagesRepositories) GetMessageById(id int) (models.Topics_Join_Messages, error) {
	var item models.Topics_Join_Messages

	// Query permettant de récupérer un message selon l'id avec le topic lié au message et l'utilisateur
	query := `
	SELECT m.message_id, m.content, m.created_at, t.title, u.name
	FROM messages AS m
	JOIN users AS u ON u.user_id = m.user_id
	JOIN topics AS t ON t.topic_id = m.topic_id
	WHERE m.message_id = ?
	`

	// Récupération de la query en une seul "row"
	sqlErr := r.db.QueryRow(query, id).Scan(&item.Messages.Message_id, &item.Messages.Content, &item.Messages.Created_at, &item.Topics.Title, &item.Users.Name)
	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return models.Topics_Join_Messages{}, nil
		}
		return models.Topics_Join_Messages{}, fmt.Errorf(" Erreur récupération item - Erreur : \n\t %s", sqlErr.Error())
	}

	return item, nil
}

func (r *MessagesRepositories) PostReplie(reply models.Replies_Joins_User_Message) (int, error) {
	query := "INSERT INTO `message_replies`(`user_id`, `reply_to_id`, `content`) VALUES (?,?,?);"

	// Utilisation de la query en remplaçant les valeurs par celles à injecter
	sqlResult, sqlErr := r.db.Exec(query,
		reply.Users.User_id,
		reply.Messages.Message_id,
		reply.Replies.Content,
	)
	if sqlErr != nil {
		return -1, fmt.Errorf(" Erreur ajout reponse - Erreur : \n\t %s", sqlErr.Error())
	}

	id, idErr := sqlResult.LastInsertId()
	if idErr != nil {
		return -1, fmt.Errorf(" Erreur ajout reponse - Erreur récupération identifiant : \n\t %s", idErr.Error())
	}

	return int(id), nil
}

func (r *MessagesRepositories) CreateMessage(message models.Messages) (int, error) {
	query := "INSERT INTO messages (topic_id, user_id, content) VALUES (?, ?, ?)"

	// Utilisation de la query en remplaçant les valeurs par celles à injecter
	sqlResult, sqlErr := r.db.Exec(query,
		message.Topic_id,
		message.User_id,
		message.Content,
	)
	if sqlErr != nil {
		return -1, fmt.Errorf(" erreur lors de l'ajout du message - Erreur : \n\t %s", sqlErr.Error())
	}

	id, idErr := sqlResult.LastInsertId()
	if idErr != nil {
		return -1, fmt.Errorf(" erreur lors de la récupération de l'identifiant - Erreur : \n\t %s", idErr.Error())
	}

	return int(id), nil
}
