package repositories

import (
	"database/sql"
	"fmt"
	"forum/src/models"
	"log"
)

type ForumsRepositories struct {
	db *sql.DB
}

func ForumsRepositoriesInit(db *sql.DB) *ForumsRepositories {
	return &ForumsRepositories{db: db}
}

// GetAllForums récupère tous les forums de la base de données.
func (r *ForumsRepositories) GetAllForums() ([]models.Forums, error) {
	var forums []models.Forums
	query := `SELECT forum_id, category_id, name, description FROM forums ORDER BY name ASC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("échec de la requête SQL pour les forums : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var forum models.Forums
		if err := rows.Scan(&forum.Forum_id, &forum.Categorie_id, &forum.Name, &forum.Description); err != nil {
			log.Printf("Erreur de scan pour un forum : %v", err)
			continue
		}
		forums = append(forums, forum)
	}
	return forums, nil
}

// GetByCategoryID récupère tous les forums pour un ID de catégorie donné.
func (r *ForumsRepositories) GetByCategoryID(categoryID int) ([]models.Forums, error) {
	var forums []models.Forums
	query := `SELECT forum_id, category_id, name, description FROM forums WHERE category_id = ? ORDER BY name ASC`
	rows, err := r.db.Query(query, categoryID)
	// ... (la logique de boucle et de scan est la même que dans GetAllForums)
	if err != nil {
		return nil, fmt.Errorf("échec de la requête SQL pour les forums : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var forum models.Forums
		if err := rows.Scan(&forum.Forum_id, &forum.Categorie_id, &forum.Name, &forum.Description); err != nil {
			log.Printf("Erreur de scan pour un forum : %v", err)
			continue
		}
		forums = append(forums, forum)
	}
	return forums, nil
}
