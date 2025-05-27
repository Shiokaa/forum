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

func (r *TopicsRepositories) GetTopicTitleWithCreatorName() ([]*models.Topics, []*models.Users, error) {
	var topicList []*models.Topics
	var userList []*models.Users

	query := `
	SELECT t.title, u.name 
	FROM topics AS t 
	JOIN users AS u 
	ON u.user_id = t.topic_id
	`

	sqlResult, sqlErr := r.db.Query(query)
	if sqlErr != nil {
		return topicList, userList, fmt.Errorf(" Erreur récupération produit - Erreur : \n\t %s", sqlErr.Error())
	}

	defer sqlResult.Close()

	for sqlResult.Next() {
		var topics models.Topics
		var users models.Users
		scanErr := sqlResult.Scan(&topics.Title, &users.Name)
		if scanErr != nil {
			log.Println(" Erreur lors du scan de la fonction : GetTopicWithCreatorName")
			continue
		}

		topicList = append(topicList, &topics)
		userList = append(userList, &users)
	}

	return topicList, userList, nil
}
