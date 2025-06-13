package repositories

import (
	"database/sql"
	"fmt"
	"forum/src/models"
	"log"
)

// Structure permettant l'injection de la base de donnée
type RepliesRepositories struct {
	db *sql.DB
}

// Fonction pour initialiser le repositorie de user avec l'injection de la base de donnée
func RepliesRepositoriesInit(db *sql.DB) *RepliesRepositories {
	return &RepliesRepositories{db: db}
}

func (r *RepliesRepositories) GetReplies(id int) ([]models.Replies_Join_User, error) {
	var items []models.Replies_Join_User

	// Query permettant de récupérer les réponses à un message
	query := `
	SELECT mr.content, mr.created_at, u.name, mr.user_id, mr.reply_id
	FROM message_replies AS mr
	JOIN users AS u ON u.user_id = mr.user_id
	WHERE mr.reply_to_id = ?;
    `

	// Récupération de la query en "row"
	rows, err := r.db.Query(query, id)
	if err != nil {
		return items, fmt.Errorf(" échec de la requête SQL : %w", err)
	}
	defer rows.Close()

	// On parcoure chaque "row" pour pouvoir les envoyés dans notre structure de join
	for rows.Next() {
		var item models.Replies_Join_User

		if err := rows.Scan(&item.Replies.Content, &item.Replies.Created_at, &item.Users.Name, &item.Replies.User_id, &item.Replies.Reply_id); err != nil {
			log.Printf(" Erreur de scan topics : %v", err)
			continue
		}

		items = append(items, item)
	}

	return items, nil
}

// GetReplyByID récupère une seule réponse par son ID.
func (r *RepliesRepositories) GetReplyByID(id int) (models.Replies, error) {
	var reply models.Replies
	query := "SELECT reply_id, user_id, reply_to_id FROM message_replies WHERE reply_id = ?"
	err := r.db.QueryRow(query, id).Scan(&reply.Reply_id, &reply.User_id, &reply.Reply_to_id)
	if err != nil {
		return models.Replies{}, err
	}
	return reply, nil
}

// DeleteReply supprime une réponse par son ID.
func (r *RepliesRepositories) DeleteReply(id int) error {
	query := "DELETE FROM message_replies WHERE reply_id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
