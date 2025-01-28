package service

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/litegral/freepass-2025/internal/lib/config"
	"github.com/litegral/freepass-2025/internal/lib/db"
	"github.com/litegral/freepass-2025/internal/model"
)

type SessionService struct {
	queries *db.Queries
}

func NewSessionService(queries *db.Queries) *SessionService {
	return &SessionService{
		queries: queries,
	}
}

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
