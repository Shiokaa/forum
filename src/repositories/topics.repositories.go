package repositories

import (
	"database/sql"
	"fmt"
	"forum/src/models"
	"log"
)

// Structure permettant l'injection de la base de donnée
type TopicsRepositories struct {
	db *sql.DB
}

// Fonction pour initialiser le repositorie de user avec l'injection de la base de donnée
func TopicsRepositoriesInit(db *sql.DB) *TopicsRepositories {
	return &TopicsRepositories{db: db}
}

// Fonction permettant d'initialiser un join entre les topics et les users
func (r *TopicsRepositories) GetTopicsWithCreators() ([]models.Topics_Join_Users, error) {
	var items []models.Topics_Join_Users

	// Query permettant d'effectuer un join entre users et topics
	query := `
    SELECT t.topic_id, t.title, t.created_at, u.name
    FROM topics AS t
    JOIN users AS u ON u.user_id = t.user_id
    `

	// Récupération de la query en "row"
	rows, err := r.db.Query(query)
	if err != nil {
		return items, fmt.Errorf("échec de la requête SQL : %w", err)
	}
	defer rows.Close()

	// On parcoure chaque "row" pour pouvoir les envoyés dans notre structure de join
	for rows.Next() {
		var item models.Topics_Join_Users

		if err := rows.Scan(&item.Topics.Topic_id, &item.Topics.Title, &item.Topics.Created_at, &item.Users.Name); err != nil {
			log.Printf("Erreur de scan topics : %v", err)
			continue
		}

		items = append(items, item)
	}

	return items, nil
}

func (r *TopicsRepositories) GetTopicWithId(id int) (models.Topics_Join_Users_Forums, error) {
	var item models.Topics_Join_Users_Forums

	query := `
	SELECT t.topic_id, t.forum_id, t.user_id, t.title, t.status, t.created_at, t.updated_at, u.name, f.name
	FROM topics AS t
	JOIN users AS u ON u.user_id = t.user_id
	JOIN forums AS f ON f.forum_id = t.forum_id
	WHERE t.topic_id = ?
	`

	sqlErr := r.db.QueryRow(query, id).Scan(&item.Topics.Topic_id, &item.Topics.Forum_id, &item.Topics.User_Id, &item.Topics.Title, &item.Topics.Status, &item.Topics.Created_at, &item.Topics.Updated_at, &item.Users.Name, &item.Forums.Name)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return models.Topics_Join_Users_Forums{}, nil
		}
		return models.Topics_Join_Users_Forums{}, fmt.Errorf(" Erreur récupération item - Erreur : \n\t %s", sqlErr.Error())
	}

	return item, nil
}

func (r *TopicsRepositories) GetTopicWithMessage(id int) ([]models.Topics_Join_Messages, error) {
	var items []models.Topics_Join_Messages

	// Query permettant d'effectuer un join entre users et topics
	query := `
	SELECT m.content, m.created_at, u.name
	FROM messages AS m
	JOIN topics AS t
    JOIN users as u ON m.user_id = u.user_id
	WHERE (t.topic_id = ?) = m.topic_id
    `

	// Récupération de la query en "row"
	rows, err := r.db.Query(query, id)
	if err != nil {
		return items, fmt.Errorf("échec de la requête SQL : %w", err)
	}
	defer rows.Close()

	// On parcoure chaque "row" pour pouvoir les envoyés dans notre structure de join
	for rows.Next() {
		var item models.Topics_Join_Messages

		if err := rows.Scan(&item.Messages.Content, &item.Messages.Created_at, &item.Users.Name); err != nil {
			log.Printf("Erreur de scan messages : %v", err)
			continue
		}

		items = append(items, item)
	}

	return items, nil

}
