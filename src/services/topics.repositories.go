package services

import (
	"database/sql"
	"forum/src/models"
	"forum/src/repositories"
)

type TopicsServices struct {
	topicsRepositories *repositories.TopicsRepositories
}

func TopicsServicesInit(db *sql.DB) *TopicsServices {
	return &TopicsServices{topicsRepositories: repositories.TopicsRepositoriesInit(db)}
}

func (s *TopicsServices) Display() ([]*models.Topics, []*models.Users, error) {

	topicTitle, userName, err := s.topicsRepositories.GetTopicTitleWithCreatorName()
	if err != nil {
		return nil, nil, err
	}

	return topicTitle, userName, nil
}
