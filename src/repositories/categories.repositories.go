package repositories

import (
	"database/sql"
	"fmt"
	"forum/src/models"
	"log"
)

type CategoriesRepositories struct {
	db *sql.DB
}

func CategoriesRepositoriesInit(db *sql.DB) *CategoriesRepositories {
	return &CategoriesRepositories{db: db}
}

// GetAllCategories récupère toutes les catégories de la base de données.
func (r *CategoriesRepositories) GetAllCategories() ([]models.Categories, error) {
	var categories []models.Categories
	query := `SELECT category_id, name, description FROM categories ORDER BY name ASC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("échec de la requête SQL pour les catégories : %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Categories
		if err := rows.Scan(&category.Category_id, &category.Name, &category.Description); err != nil {
			log.Printf("Erreur de scan pour une catégorie : %v", err)
			continue
		}
		categories = append(categories, category)
	}
	return categories, nil
}

// GetByID récupère une seule catégorie par son ID.
func (r *CategoriesRepositories) GetByID(id int) (models.Categories, error) {
	var category models.Categories
	query := `SELECT category_id, name, description FROM categories WHERE category_id = ?`
	err := r.db.QueryRow(query, id).Scan(&category.Category_id, &category.Name, &category.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Categories{}, fmt.Errorf("aucune catégorie trouvée avec l'ID %d", id)
		}
		return models.Categories{}, err
	}
	return category, nil
}
