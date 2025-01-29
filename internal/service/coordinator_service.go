package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/litegral/freepass-2025/internal/lib/config"
	"github.com/litegral/freepass-2025/internal/lib/db"
	"github.com/litegral/freepass-2025/internal/model"
)

// CoordinatorService handles coordinator-specific operations
type CoordinatorService struct {
	queries *db.Queries
}

func NewCoordinatorService(queries *db.Queries) *CoordinatorService {
	return &CoordinatorService{
		queries: queries,
	}
}

// ListSessionProposals returns all pending session proposals
func (s *CoordinatorService) ListSessionProposals(ctx context.Context) ([]model.SessionWithDetails, error) {
	// Get proposals
	proposals, err := s.queries.ListSessionProposals(ctx)
	if err != nil {
		return nil, errors.New(config.ErrFetchingProposals)
	}

	// Convert proposals to model
	result := make([]model.SessionWithDetails, 0, len(proposals))
	for _, p := range proposals {
		// Get feedback for each proposal
		feedback, err := s.queries.ListSessionFeedback(ctx, p.ID)
		if err != nil {
			continue
		}

		session := convertToSessionWithDetails(p, feedback)
		result = append(result, session)
	}

	return result, nil
}

// UpdateSessionStatus updates the status of a session proposal
func (s *CoordinatorService) UpdateSessionStatus(ctx context.Context, sessionID int32, status string) (model.Session, error) {
	// Check if status is valid
	if status != "accepted" && status != "rejected" {
		return model.Session{}, errors.New(config.ErrInvalidSessionStatus)
	}

	// Update session status
	session, err := s.queries.UpdateSessionStatusByCoordinator(ctx, db.UpdateSessionStatusByCoordinatorParams{
		ID:     sessionID,
		Status: db.SessionStatus(status),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Session{}, errors.New(config.ErrSessionAlreadyProcessed)
		}
		return model.Session{}, fmt.Errorf("error updating session status: %w", err)
	}

	return convertDBSessionToModel(session), nil
}

// RemoveSession soft deletes a session
func (s *CoordinatorService) RemoveSession(ctx context.Context, sessionID int32) error {
	err := s.queries.SoftDeleteSessionByCoordinator(ctx, sessionID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New(config.ErrSessionNotFound)
		}
		return fmt.Errorf("error removing session: %w", err)
	}
	return nil
}

// RemoveFeedback soft deletes inappropriate feedback
func (s *CoordinatorService) RemoveFeedback(ctx context.Context, feedbackID int32) error {
	err := s.queries.SoftDeleteFeedbackByCoordinator(ctx, feedbackID)
	if err != nil {
		return fmt.Errorf("error removing feedback: %w", err)
	}
	return nil
}

// convertToSessionWithDetails converts db.ListSessionProposalsRow to model.SessionWithDetails
func convertToSessionWithDetails(p db.ListSessionProposalsRow, feedback []db.ListSessionFeedbackRow) model.SessionWithDetails {
	return model.SessionWithDetails{
		Session: model.Session{
			ID:              int(p.ID),
			Title:           p.Title,
			Description:     p.Description.String,
			StartTime:       p.StartTime.Time,
			EndTime:         p.EndTime.Time,
			Room:            p.Room.String,
			Status:          string(p.Status),
			SeatingCapacity: int(p.SeatingCapacity),
			ProposerID:      int(p.ProposerID.Int32),
			CreatedAt:       p.CreatedAt.Time,
			UpdatedAt:       p.UpdatedAt.Time,
		},
		ProposerName:        p.ProposerName.String,
		ProposerAffiliation: p.ProposerAffiliation.String,
		Feedback:            convertDBFeedbackToModel(feedback),
	}
}
