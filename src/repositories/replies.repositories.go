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
	SELECT mr.content, mr.created_at, u.name
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

		if err := rows.Scan(&item.Replies.Content, &item.Replies.Created_at, &item.Users.Name); err != nil {
			log.Printf(" Erreur de scan topics : %v", err)
			continue
		}

		items = append(items, item)
	}

	return items, nil
}
