package repositories

import (
	"database/sql"
	"fmt"
	"forum/src/models"
	"log"
)

type TopicsRepositories struct {
	db *sql.DB
}

func TopicsRepositoriesInit(db *sql.DB) *TopicsRepositories {
	return &TopicsRepositories{db: db}
}

func (r *TopicsRepositories) GetTopicsWithCreators() ([]string, []string, error) {
	var users []string
	var topics []string

	query := `
    SELECT t.title, t.created_at, u.name
    FROM topics AS t
    JOIN users AS u ON u.user_id = t.user_id
    `

	rows, err := r.db.Query(query)
	if err != nil {
		return []string{}, []string{}, fmt.Errorf("échec de la requête SQL : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var topic models.Topics
		var user models.Users
		if err := rows.Scan(&topic.Title, &topic.Created_at, &user.Name); err != nil {
			log.Printf("Erreur de scan : %v", err)
			continue
		}

		users = append(users, user.Name)
		topics = append(topics, topic.Title, topic.Created_at)
	}

	return topics, users, nil
}
