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
		return items, fmt.Errorf(" échec de la requête SQL : %w", err)
	}
	defer rows.Close()

	// On parcoure chaque "row" pour pouvoir les envoyés dans notre structure de join
	for rows.Next() {
		var item models.Topics_Join_Users

		if err := rows.Scan(&item.Topics.Topic_id, &item.Topics.Title, &item.Topics.Created_at, &item.Users.Name); err != nil {
			log.Printf(" Erreur de scan topics : %v", err)
			continue
		}

		items = append(items, item)
	}

	return items, nil
}

func (r *TopicsRepositories) GetTopicWithId(id int) (models.Topics_Join_Users_Forums, error) {
	var item models.Topics_Join_Users_Forums

	// Query permettant de join l'utilisateur ayant écrit le topic ainsi que le forum du topic, aux détails du topic
	query := `
	SELECT t.topic_id, t.forum_id, t.user_id, t.title, t.status, t.created_at, t.updated_at, u.name, f.name
	FROM topics AS t
	JOIN users AS u ON u.user_id = t.user_id
	JOIN forums AS f ON f.forum_id = t.forum_id
	WHERE t.topic_id = ?
	`

	// Récupération de la query en une seul "row"
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

	// Query permettant d'effectuer un join sur les messsages pour récupérer l'user du message et le topic où se trouve le message
	query := `
	SELECT m.message_id, m.content, m.created_at, u.name, m.user_id
	FROM messages AS m
	JOIN users AS u ON m.user_id = u.user_id
	JOIN topics AS t ON m.topic_id = t.topic_id
	WHERE t.topic_id = ?
    `

	// Récupération de la query en "row"
	rows, err := r.db.Query(query, id)
	if err != nil {
		return items, fmt.Errorf(" échec de la requête SQL : %w", err)
	}
	defer rows.Close()

	// On parcoure chaque "row" pour pouvoir les envoyés dans notre structure de join
	for rows.Next() {
		var item models.Topics_Join_Messages

		if err := rows.Scan(&item.Messages.Message_id, &item.Messages.Content, &item.Messages.Created_at, &item.Users.Name, &item.Messages.User_id); err != nil {
			log.Printf(" Erreur de scan messages : %v", err)
			continue
		}

		items = append(items, item)
	}

	return items, nil

}

func (r *TopicsRepositories) GetPaginatedTopics(offset, limit int) ([]models.Topics_Join_Users, error) {
	var items []models.Topics_Join_Users
	query := `
    SELECT t.topic_id, t.title, t.created_at, u.name
    FROM topics AS t
    JOIN users AS u ON u.user_id = t.user_id
	ORDER BY t.created_at DESC
	LIMIT ? OFFSET ?
    `
	rows, err := r.db.Query(query, limit, offset)
	// ... (le reste de la fonction est similaire à GetTopicsWithCreators)
	if err != nil {
		return items, fmt.Errorf(" échec de la requête SQL : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Topics_Join_Users

		if err := rows.Scan(&item.Topics.Topic_id, &item.Topics.Title, &item.Topics.Created_at, &item.Users.Name); err != nil {
			log.Printf(" Erreur de scan topics : %v", err)
			continue
		}
		items = append(items, item)
	}
	return items, nil
}

// GetTotalTopicsCount compte le nombre total de topics.
func (r *TopicsRepositories) GetTotalTopicsCount() (int, error) {
	var count int
	query := "SELECT COUNT(topic_id) FROM topics"
	err := r.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("échec du comptage des topics : %w", err)
	}
	return count, nil
}

// GetByCategoryID récupère les derniers topics pour un ID de catégorie donné.
func (r *TopicsRepositories) GetByCategoryID(categoryID int, limit int) ([]models.Topics_Join_Users, error) {
	var items []models.Topics_Join_Users
	query := `
    SELECT t.topic_id, t.title, t.created_at, u.name
    FROM topics AS t
    JOIN users AS u ON u.user_id = t.user_id
    JOIN forums AS f ON f.forum_id = t.forum_id
    WHERE f.category_id = ?
    ORDER BY t.created_at DESC
    LIMIT ?`
	rows, err := r.db.Query(query, categoryID, limit)
	// ... (la logique de boucle et de scan est la même que dans GetPaginatedTopics)
	if err != nil {
		return items, fmt.Errorf(" échec de la requête SQL : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Topics_Join_Users
		if err := rows.Scan(&item.Topics.Topic_id, &item.Topics.Title, &item.Topics.Created_at, &item.Users.Name); err != nil {
			log.Printf(" Erreur de scan topics : %v", err)
			continue
		}
		items = append(items, item)
	}
	return items, nil
}

// GetByForumIDPaginated récupère les topics paginés pour un ID de forum donné.
func (r *TopicsRepositories) GetByForumIDPaginated(forumID, offset, limit int) ([]models.Topics_Join_Users, error) {
	var items []models.Topics_Join_Users
	query := `
    SELECT t.topic_id, t.title, t.created_at, u.name
    FROM topics AS t
    JOIN users AS u ON u.user_id = t.user_id
    WHERE t.forum_id = ?
    ORDER BY t.created_at DESC
    LIMIT ? OFFSET ?`

	rows, err := r.db.Query(query, forumID, limit, offset)
	// ... (la logique de boucle et de scan est la même que dans les autres fonctions de récupération de topics)
	if err != nil {
		return items, fmt.Errorf("échec de la requête SQL : %w", err)
	}
	defer rows.Close()

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

// GetTotalTopicsCountByForumID compte le nombre total de topics dans un forum.
func (r *TopicsRepositories) GetTotalTopicsCountByForumID(forumID int) (int, error) {
	var count int
	query := "SELECT COUNT(topic_id) FROM topics WHERE forum_id = ?"
	err := r.db.QueryRow(query, forumID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("échec du comptage des topics pour le forum %d : %w", forumID, err)
	}
	return count, nil
}

// CreateTopic insère un nouveau topic dans la base de données.
func (r *TopicsRepositories) CreateTopic(topic models.Topics) (int, error) {
	query := "INSERT INTO `topics`(`forum_id`, `user_id`, `title`, `status`) VALUES (?,?,?,?);"

	// Le statut est `true` par défaut pour un nouveau topic.
	sqlResult, sqlErr := r.db.Exec(query, topic.Forum_id, topic.User_Id, topic.Title, true)
	if sqlErr != nil {
		return -1, fmt.Errorf("erreur lors de l'ajout du topic : %w", sqlErr)
	}

	id, idErr := sqlResult.LastInsertId()
	if idErr != nil {
		return -1, fmt.Errorf("erreur lors de la récupération de l'ID du topic : %w", idErr)
	}

	return int(id), nil
}

// DeleteTopic supprime un topic par son ID.
func (r *TopicsRepositories) DeleteTopic(id int) error {
	query := "DELETE FROM topics WHERE topic_id = ?"
	_, err := r.db.Exec(query, id)
	return err
}

// GetTopicsByUserID récupère tous les topics créés par un utilisateur spécifique.
func (r *TopicsRepositories) GetTopicsByUserID(userID int) ([]models.Topics_Join_Users, error) {
	var items []models.Topics_Join_Users
	query := `
    SELECT t.topic_id, t.title, t.created_at, u.name
    FROM topics AS t
    JOIN users AS u ON t.user_id = u.user_id
    WHERE t.user_id = ?
    ORDER BY t.created_at DESC
    `
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("échec de la requête pour les topics de l'utilisateur : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Topics_Join_Users
		if err := rows.Scan(&item.Topics.Topic_id, &item.Topics.Title, &item.Topics.Created_at, &item.Users.Name); err != nil {
			log.Printf("Erreur de scan pour un topic utilisateur : %v", err)
			continue
		}
		items = append(items, item)
	}
	return items, nil
}
