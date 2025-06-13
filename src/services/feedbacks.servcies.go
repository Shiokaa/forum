package services

import (
	"database/sql"
	"forum/src/models"
	"forum/src/repositories"
)

type FeedbacksServices struct {
	repo *repositories.FeedbacksRepository
}

func FeedbacksServicesInit(db *sql.DB) *FeedbacksServices {
	return &FeedbacksServices{repo: repositories.FeedbacksRepositoryInit(db)}
}

// HandleFeedback gère la logique d'un vote.
func (s *FeedbacksServices) HandleFeedback(userID, messageID int, newVote string) error {
	currentVote, err := s.repo.GetFeedbackForMessageByUser(userID, messageID)
	if err != nil {
		return err
	}

	if currentVote == newVote {
		// L'utilisateur clique sur le même bouton, on annule le vote
		return s.repo.RemoveFeedback(userID, messageID)
	} else {
		// Nouveau vote ou changement de vote
		return s.repo.SetFeedback(userID, messageID, newVote)
	}
}

// GetFeedbackInfoForMessages traite les votes bruts pour les afficher.
func (s *FeedbacksServices) GetFeedbackInfoForMessages(messageIDs []int, currentUserID int) (map[int]models.FeedbackInfo, error) {
	// Le modèle FeedbackInfo doit être ajouté à un fichier de modèles, par ex: join.models.go
	results := make(map[int]models.FeedbackInfo)
	if len(messageIDs) == 0 {
		return results, nil
	}

	rawFeedbacks, err := s.repo.GetFeedbacksForMessageIDs(messageIDs)
	if err != nil {
		return nil, err
	}

	// Initialiser la map pour tous les messages
	for _, id := range messageIDs {
		results[id] = models.FeedbackInfo{}
	}

	// Compter les votes et trouver le vote de l'utilisateur actuel
	for _, feedback := range rawFeedbacks {
		info := results[feedback.MessageID]
		if feedback.FeedbackType == "like" {
			info.LikeCount++
		} else if feedback.FeedbackType == "dislike" {
			info.DislikeCount++
		}
		// Ce n'est pas la bonne approche pour le vote de l'utilisateur, à revoir.
		// La logique pour le vote de l'utilisateur sera gérée séparément pour la simplicité.
		results[feedback.MessageID] = info
	}
	return results, nil
}

// GetUserVotesForMessages récupère les votes d'un utilisateur pour une liste de messages.
func (s *FeedbacksServices) GetUserVotesForMessages(messageIDs []int, userID int) (map[int]string, error) {
	userVotes := make(map[int]string)
	if userID == 0 || len(messageIDs) == 0 {
		return userVotes, nil
	}

	for _, msgID := range messageIDs {
		vote, _ := s.repo.GetFeedbackForMessageByUser(userID, msgID)
		if vote != "" {
			userVotes[msgID] = vote
		}
	}
	return userVotes, nil
}
