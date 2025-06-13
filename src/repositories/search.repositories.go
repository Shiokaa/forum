package repositories

import (
	"database/sql"
	"fmt"
	"forum/src/models"
	"log"
)

type SearchRepositories struct {
	db *sql.DB
}

func SearchRepositoriesInit(db *sql.DB) *SearchRepositories {
	return &SearchRepositories{db: db}
}

// GlobalSearch effectue une recherche sur les catégories, forums et topics.
func (r *SearchRepositories) GlobalSearch(query string) ([]models.SearchResult, error) {
	var results []models.SearchResult
	searchQuery := "%" + query + "%"

	// La requête UNION combine les résultats des trois requêtes SELECT.
	sqlQuery := `
	(SELECT 'Catégorie' as type, name as title, description, CONCAT('/categorie?id=', category_id) as url FROM categories WHERE name LIKE ? OR description LIKE ?)
	UNION
	(SELECT 'Forum' as type, name as title, description, CONCAT('/forum?id=', forum_id) as url FROM forums WHERE name LIKE ? OR description LIKE ?)
	UNION
	(SELECT 'Topic' as type, title as title, '' as description, CONCAT('/topic?id=', topic_id) as url FROM topics WHERE title LIKE ?)
	`

	rows, err := r.db.Query(sqlQuery, searchQuery, searchQuery, searchQuery, searchQuery, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("échec de la requête de recherche globale : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.SearchResult
		// Le champ 'description' peut être NULL, nous devons utiliser sql.NullString pour le gérer.
		var description sql.NullString
		if err := rows.Scan(&item.Type, &item.Title, &description, &item.URL); err != nil {
			log.Printf("Erreur de scan lors de la recherche globale : %v", err)
			continue
		}
		if description.Valid {
			item.Description = description.String
		}
		results = append(results, item)
	}
	return results, nil
}
