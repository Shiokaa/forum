package services

import (
	"database/sql"
	"forum/src/models"
	"forum/src/repositories"
)

type ForumsServices struct {
	forumsRepositories *repositories.ForumsRepositories
}

func ForumsServicesInit(db *sql.DB) *ForumsServices {
	return &ForumsServices{forumsRepositories: repositories.ForumsRepositoriesInit(db)}
}

func (s *ForumsServices) GetAll() ([]models.Forums, error) {
	return s.forumsRepositories.GetAllForums()
}

func (s *ForumsServices) GetByCategoryID(categoryID int) ([]models.Forums, error) {
	return s.forumsRepositories.GetByCategoryID(categoryID)
}

func (s *ForumsServices) GetByIDWithCategory(id int) (models.ForumWithCategory, error) {
	return s.forumsRepositories.GetByIDWithCategory(id)
}
