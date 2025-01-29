package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/litegral/freepass-2025/internal/lib/config"
	"github.com/litegral/freepass-2025/internal/lib/db"
	"github.com/litegral/freepass-2025/internal/model"
)

// SessionService is a service for managing sessions
type SessionService struct {
	queries *db.Queries
}

func NewSessionService(queries *db.Queries) *SessionService {
	return &SessionService{
		queries: queries,
	}
}

// CreateProposal creates a new session proposal
func (s *SessionService) CreateProposal(ctx context.Context, userID int32, proposal model.SessionProposal) (model.Session, error) {
	// Check if user already has a pending proposal
	sessions, err := s.queries.ListSessions(ctx)
	if err != nil {
		return model.Session{}, err
	}

	// Check for existing proposals in the current conference cycle
	for _, session := range sessions {
		if session.ProposerID.Int32 == userID && session.Status == "pending" {
			return model.Session{}, errors.New(config.ErrPendingProposalExists)
		}
	}

	// Get conference config to check if we're in the proposal period
	c, err := s.queries.GetConferenceConfig(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Session{}, errors.New(config.ErrNoConferenceConfig)
		}
		return model.Session{}, err
	}

	now := time.Now()
	if now.Before(c.SessionProposalStart.Time) || now.After(c.SessionProposalEnd.Time) {
		return model.Session{}, errors.New(config.ErrProposalsNotOpen)
	}

	// Create the session proposal
	session, err := s.queries.CreateSessionProposal(ctx, db.CreateSessionProposalParams{
		Title:           proposal.Title,
		Description:     pgtype.Text{String: proposal.Description, Valid: true},
		StartTime:       pgtype.Timestamptz{Time: proposal.StartTime, Valid: true},
		EndTime:         pgtype.Timestamptz{Time: proposal.EndTime, Valid: true},
		Room:            pgtype.Text{String: proposal.Room, Valid: true},
		SeatingCapacity: int32(proposal.SeatingCapacity),
		ProposerID:      pgtype.Int4{Int32: userID, Valid: true},
	})
	if err != nil {
		return model.Session{}, err
	}

	return convertDBSessionToModel(session), nil
}

// UpdateProposal updates an existing session proposal
func (s *SessionService) UpdateProposal(ctx context.Context, userID int32, sessionID int32, update model.SessionUpdate) (model.Session, error) {
	// Get the session to verify ownership
	session, err := s.queries.GetSessionByID(ctx, sessionID)
	if err != nil {
		return model.Session{}, err
	}

	// Verify ownership
	if session.ProposerID.Int32 != userID {
		return model.Session{}, errors.New(config.ErrUnauthorizedSession)
	}

	// Verify session is still pending
	if session.Status != "pending" {
		return model.Session{}, errors.New(config.ErrOnlyPendingUpdate)
	}

	// Update the session
	updatedSession, err := s.queries.UpdateSession(ctx, db.UpdateSessionParams{
		ID:              sessionID,
		Title:           update.Title,
		Description:     pgtype.Text{String: update.Description, Valid: true},
		StartTime:       pgtype.Timestamptz{Time: update.StartTime, Valid: true},
		EndTime:         pgtype.Timestamptz{Time: update.EndTime, Valid: true},
		Room:            pgtype.Text{String: update.Room, Valid: true},
		SeatingCapacity: int32(update.SeatingCapacity),
		Status:          "pending",
	})
	if err != nil {
		return model.Session{}, err
	}

	return convertDBSessionToModel(updatedSession), nil
}

// DeleteProposal deletes an existing session proposal
func (s *SessionService) DeleteProposal(ctx context.Context, userID int32, sessionID int32) error {
	// Add validation for session ID
	if sessionID <= 0 {
		return errors.New("invalid session ID")
	}

	// Get the session to verify ownership
	session, err := s.queries.GetSessionByID(ctx, sessionID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New(config.ErrSessionNotFound)
		}
		return err
	}

	// Verify ownership
	if session.ProposerID.Int32 != userID {
		return errors.New(config.ErrUnauthorizedSessionDel)
	}

	// Verify session is still pending
	if session.Status != "pending" {
		return errors.New(config.ErrOnlyPendingDelete)
	}

	// Delete the session
	err = s.queries.SoftDeleteSession(ctx, sessionID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New(config.ErrSessionNotFound)
		}
		return err
	}

	return nil
}

// ListSessions lists all sessions
func (s *SessionService) ListSessions(ctx context.Context, userID int32) ([]model.SessionWithDetails, error) {
	// Get all sessions
	sessions, err := s.queries.ListSessions(ctx)
	if err != nil {
		return nil, err
	}

	var result []model.SessionWithDetails
	for _, session := range sessions {
		// Get session details including proposer info
		details, err := s.queries.GetSessionByID(ctx, session.ID)
		if err != nil {
			continue
		}

		// Get feedback for the session
		feedback, err := s.queries.ListSessionFeedback(ctx, session.ID)
		if err != nil {
			continue
		}

		// Check if user is registered
		isRegistered := false
		reg, err := s.queries.GetRegistration(ctx, db.GetRegistrationParams{
			UserID:    userID,
			SessionID: session.ID,
		})
		if err == nil {
			isRegistered = reg.ID > 0
		}

		// Convert to model
		sessionDetails := model.SessionWithDetails{
			Session:             convertDBSessionToModel(session),
			ProposerName:        details.ProposerName.String,
			ProposerAffiliation: details.ProposerAffiliation.String,
			IsRegistered:        isRegistered,
			AvailableSeats:      int(session.SeatingCapacity),
			Feedback:            convertDBFeedbackToModel(feedback),
		}

		result = append(result, sessionDetails)
	}

	return result, nil
}

// GetSession gets a specific session by ID
func (s *SessionService) GetSession(ctx context.Context, sessionID int32, userID int32) (model.SessionWithDetails, error) {
	// Get session details
	session, err := s.queries.GetSessionByID(ctx, sessionID)
	if err != nil {
		return model.SessionWithDetails{}, err
	}

	// Get feedback
	feedback, err := s.queries.ListSessionFeedback(ctx, sessionID)
	if err != nil {
		return model.SessionWithDetails{}, err
	}

	// Check registration
	isRegistered := false
	reg, err := s.queries.GetRegistration(ctx, db.GetRegistrationParams{
		UserID:    userID,
		SessionID: sessionID,
	})
	if err == nil {
		isRegistered = reg.ID > 0
	}

	return model.SessionWithDetails{
		Session:             convertSessionRowToModel(session),
		ProposerName:        session.ProposerName.String,
		ProposerAffiliation: session.ProposerAffiliation.String,
		Feedback:            convertDBFeedbackToModel(feedback),
		IsRegistered:        isRegistered,
		AvailableSeats:      int(session.SeatingCapacity),
	}, nil
}

// RegisterForSession registers a user for a specific session
func (s *SessionService) RegisterForSession(ctx context.Context, sessionID int32, userID int32) error {
	// Get conference config to check registration period
	conf, err := s.queries.GetConferenceConfig(ctx)
	if err != nil {
		return errors.New("conference configuration not found")
	}

	now := time.Now()
	if now.Before(conf.RegistrationStart.Time) || now.After(conf.RegistrationEnd.Time) {
		return errors.New("registration is not currently open")
	}

	// Check if session exists and is accepted
	session, err := s.queries.GetSessionByID(ctx, sessionID)
	if err != nil {
		return errors.New("session not found")
	}

	if session.Status != "accepted" {
		return errors.New("can only register for accepted sessions")
	}

	// Count current registrations for the session
	registrations, err := s.queries.CountSessionRegistrations(ctx, sessionID)
	if err != nil {
		return fmt.Errorf("error checking seat availability: %w", err)
	}

	// Check if seats are available
	if int32(registrations) >= session.SeatingCapacity {
		return errors.New("session is full")
	}

	// Check if user is already registered
	_, err = s.queries.GetRegistration(ctx, db.GetRegistrationParams{
		UserID:    userID,
		SessionID: sessionID,
	})
	if err == nil {
		return errors.New("already registered for the session")
	}

	// Create registration
	_, err = s.queries.RegisterForSession(ctx, db.RegisterForSessionParams{
		UserID:    userID,
		SessionID: sessionID,
	})
	
	return err
}

// CreateFeedback creates feedback for a specific session
func (s *SessionService) CreateFeedback(ctx context.Context, sessionID int32, userID int32, comment string) (model.SessionFeedbackResponse, error) {
	// Check if session exists
	_, err := s.queries.GetSessionByID(ctx, sessionID)
	if err != nil {
		return model.SessionFeedbackResponse{}, errors.New("session not found")
	}

	// Check if user is registered for the session
	_, err = s.queries.GetRegistration(ctx, db.GetRegistrationParams{
		UserID:    userID,
		SessionID: sessionID,
	})
	if err != nil {
		return model.SessionFeedbackResponse{}, errors.New("must be registered for session to leave feedback")
	}

	// Create feedback
	feedback, err := s.queries.CreateFeedback(ctx, db.CreateFeedbackParams{
		UserID:    userID,
		SessionID: sessionID,
		Comment:   comment,
	})
	if err != nil {
		return model.SessionFeedbackResponse{}, err
	}

	// Get user details for response
	user, err := s.queries.GetUserByID(ctx, userID)
	if err != nil {
		return model.SessionFeedbackResponse{}, err
	}

	return model.SessionFeedbackResponse{
		ID:          int(feedback.ID),
		UserID:      int(feedback.UserID),
		SessionID:   int(feedback.SessionID),
		Comment:     feedback.Comment,
		CreatedAt:   feedback.CreatedAt.Time,
		UserName:    user.FullName.String,
		Affiliation: user.Affiliation.String,
	}, nil
}

// convertDBFeedbackToModel converts feedback from the database to the model
func convertDBFeedbackToModel(feedback []db.ListSessionFeedbackRow) []model.SessionFeedbackResponse {
	result := make([]model.SessionFeedbackResponse, len(feedback))
	for i, f := range feedback {
		result[i] = model.SessionFeedbackResponse{
			ID:          int(f.ID),
			UserID:      int(f.UserID),
			SessionID:   int(f.SessionID),
			Comment:     f.Comment,
			CreatedAt:   f.CreatedAt.Time,
			UserName:    f.FullName.String,
			Affiliation: f.Affiliation.String,
		}
	}
	return result
}

// convertDBSessionToModel converts a session from the database to the model
func convertDBSessionToModel(s db.Session) model.Session {
	return model.Session{
		ID:              int(s.ID),
		Title:           s.Title,
		Description:     s.Description.String,
		StartTime:       s.StartTime.Time,
		EndTime:         s.EndTime.Time,
		Room:            s.Room.String,
		Status:          string(s.Status),
		SeatingCapacity: int(s.SeatingCapacity),
		ProposerID:      int(s.ProposerID.Int32),
		CreatedAt:       s.CreatedAt.Time,
		UpdatedAt:       s.UpdatedAt.Time,
	}
}

// convertSessionRowToModel converts a session from the database to the model
func convertSessionRowToModel(s db.GetSessionByIDRow) model.Session {
	return model.Session{
		ID:              int(s.ID),
		Title:           s.Title,
		Description:     s.Description.String,
		StartTime:       s.StartTime.Time,
		EndTime:         s.EndTime.Time,
		Room:            s.Room.String,
		Status:          string(s.Status),
		SeatingCapacity: int(s.SeatingCapacity),
		ProposerID:      int(s.ProposerID.Int32),
		CreatedAt:       s.CreatedAt.Time,
		UpdatedAt:       s.UpdatedAt.Time,
	}
}
