package services

import (
	"database/sql"
	"forum/src/repositories"
)

type TopicsServices struct {
	topicsRepositories *repositories.TopicsRepositories
}

func TopicsServicesInit(db *sql.DB) *TopicsServices {
	return &TopicsServices{topicsRepositories: repositories.TopicsRepositoriesInit(db)}
}

func (s *TopicsServices) Display() ([]string, []string, error) {

	topics, users, err := s.topicsRepositories.GetTopicsWithCreators()
	if err != nil {
		return nil, nil, err
	}

	return topics, users, nil
}
