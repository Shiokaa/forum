package services

import (
	"database/sql"
	"fmt"
	"forum/src/models"
	"forum/src/repositories"
)

type SearchServices struct {
	searchRepositories *repositories.SearchRepositories
}

func SearchServicesInit(db *sql.DB) *SearchServices {
	return &SearchServices{searchRepositories: repositories.SearchRepositoriesInit(db)}
}

// Remplacer FindTopics par PerformGlobalSearch
func (s *SearchServices) PerformGlobalSearch(query string) ([]models.SearchResult, error) {
	if len(query) < 3 {
		return nil, fmt.Errorf("la requête de recherche doit contenir au moins 3 caractères")
	}
	return s.searchRepositories.GlobalSearch(query)
}
