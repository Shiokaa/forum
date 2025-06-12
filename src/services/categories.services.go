package services

import (
	"database/sql"
	"forum/src/models"
	"forum/src/repositories"
)

type CategoriesServices struct {
	categoriesRepositories *repositories.CategoriesRepositories
}

func CategoriesServicesInit(db *sql.DB) *CategoriesServices {
	return &CategoriesServices{categoriesRepositories: repositories.CategoriesRepositoriesInit(db)}
}

func (s *CategoriesServices) GetAll() ([]models.Categories, error) {
	return s.categoriesRepositories.GetAllCategories()
}

func (s *CategoriesServices) GetCategoriesWithForums(forumService *ForumsServices) ([]models.CategoryWithForums, error) {
	// 1. Récupérer toutes les catégories
	categories, err := s.GetAll()
	if err != nil {
		return nil, err
	}

	// 2. Récupérer tous les forums
	forums, err := forumService.GetAll()
	if err != nil {
		return nil, err
	}

	// 3. Créer une map pour associer les forums à leur category_id
	forumsByCategoryID := make(map[int][]models.Forums)
	for _, forum := range forums {
		forumsByCategoryID[forum.Categorie_id] = append(forumsByCategoryID[forum.Categorie_id], forum)
	}

	// 4. Construire la structure finale
	var result []models.CategoryWithForums
	for _, category := range categories {
		result = append(result, models.CategoryWithForums{
			Categories: category,
			Forums:     forumsByCategoryID[category.Category_id],
		})
	}

	return result, nil
}

func (s *CategoriesServices) GetByID(id int) (models.Categories, error) {
	return s.categoriesRepositories.GetByID(id)
}
