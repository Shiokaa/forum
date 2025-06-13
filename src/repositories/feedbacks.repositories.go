package repositories

import (
	"database/sql"
	"fmt"
	"strings"
)

type FeedbacksRepository struct {
	db *sql.DB
}

func FeedbacksRepositoryInit(db *sql.DB) *FeedbacksRepository {
	return &FeedbacksRepository{db: db}
}

// GetFeedbackForMessageByUser récupère le vote d'un utilisateur pour un message.
func (r *FeedbacksRepository) GetFeedbackForMessageByUser(userID, messageID int) (string, error) {
	var feedbackType string
	query := "SELECT type FROM feedbacks WHERE user_id = ? AND message_id = ?"
	err := r.db.QueryRow(query, userID, messageID).Scan(&feedbackType)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // Pas de vote, pas d'erreur
		}
		return "", err
	}
	return feedbackType, nil
}

// SetFeedback insère ou met à jour un vote.
func (r *FeedbacksRepository) SetFeedback(userID, messageID int, feedbackType string) error {
	query := `
	INSERT INTO feedbacks (user_id, message_id, type)
	VALUES (?, ?, ?)
	ON DUPLICATE KEY UPDATE type = ?
	`
	_, err := r.db.Exec(query, userID, messageID, feedbackType, feedbackType)
	return err
}

// RemoveFeedback supprime le vote d'un utilisateur pour un message.
func (r *FeedbacksRepository) RemoveFeedback(userID, messageID int) error {
	query := "DELETE FROM feedbacks WHERE user_id = ? AND message_id = ?"
	_, err := r.db.Exec(query, userID, messageID)
	return err
}

// RawFeedback est une structure pour la récupération de données brutes.
type RawFeedback struct {
	MessageID    int
	FeedbackType string
}

// GetFeedbacksForMessageIDs récupère tous les votes pour une liste de messages.
func (r *FeedbacksRepository) GetFeedbacksForMessageIDs(messageIDs []int) ([]RawFeedback, error) {
	var feedbacks []RawFeedback
	if len(messageIDs) == 0 {
		return feedbacks, nil
	}
	// Création de la chaîne de placeholders (?, ?, ?)
	placeholders := strings.Repeat("?,", len(messageIDs)-1) + "?"
	query := fmt.Sprintf("SELECT message_id, type FROM feedbacks WHERE message_id IN (%s)", placeholders)

	// Conversion des IDs en []interface{} pour la requête variadique
	args := make([]interface{}, len(messageIDs))
	for i, id := range messageIDs {
		args[i] = id
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var feedback RawFeedback
		if err := rows.Scan(&feedback.MessageID, &feedback.FeedbackType); err != nil {
			continue
		}
		feedbacks = append(feedbacks, feedback)
	}
	return feedbacks, nil
}
