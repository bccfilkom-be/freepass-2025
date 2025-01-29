package controller

import (
	"strconv"

	"github.com/go-fuego/fuego"
	"github.com/litegral/freepass-2025/internal/lib/config"
	"github.com/litegral/freepass-2025/internal/lib/jwt"
	"github.com/litegral/freepass-2025/internal/model"
	"github.com/litegral/freepass-2025/internal/service"
)

// SessionController handles session-related operations
type SessionController struct {
	sessionService *service.SessionService
}

func NewSessionController(sessionService *service.SessionService) *SessionController {
	return &SessionController{
		sessionService: sessionService,
	}
}

// CreateProposal handles creating a new session proposal
func (c *SessionController) CreateProposal(ctx fuego.ContextWithBody[model.SessionProposal]) (model.Session, error) {
	// Get user claims from context
	claims, ok := ctx.Value(jwt.ClaimsContextKey).(*jwt.Claims)
	if !ok {
		return model.Session{}, fuego.UnauthorizedError{Title: "Invalid token claims"}
	}

	// Get request body
	proposal, err := ctx.Body()
	if err != nil {
		return model.Session{}, err
	}

	// Create proposal
	session, err := c.sessionService.CreateProposal(ctx.Context(), int32(claims.UserID), proposal)
	if err != nil {
		return model.Session{}, fuego.BadRequestError{Title: err.Error()}
	}

	return session, nil
}

// UpdateProposal handles updating an existing session proposal
func (c *SessionController) UpdateProposal(ctx fuego.ContextWithBody[model.SessionUpdate]) (model.Session, error) {
	// Get user claims from context
	claims, ok := ctx.Value(jwt.ClaimsContextKey).(*jwt.Claims)
	if !ok {
		return model.Session{}, fuego.UnauthorizedError{Title: "Invalid token claims"}
	}

	// Get session ID from path
	sessionID, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return model.Session{}, fuego.BadRequestError{Title: "Invalid session ID"}
	}

	// Get request body
	update, err := ctx.Body()
	if err != nil {
		return model.Session{}, err
	}

	// Update proposal
	session, err := c.sessionService.UpdateProposal(ctx.Context(), int32(claims.UserID), int32(sessionID), update)
	if err != nil {
		return model.Session{}, fuego.BadRequestError{Title: err.Error()}
	}

	return session, nil
}

// DeleteProposal handles deleting an existing session proposal
func (c *SessionController) DeleteProposal(ctx fuego.ContextWithBody[any]) (any, error) {
	// Get user claims from context
	claims, ok := ctx.Value(jwt.ClaimsContextKey).(*jwt.Claims)
	if !ok {
		return nil, fuego.UnauthorizedError{Title: "Invalid token claims"}
	}

	// Get session ID from path and convert to int
	sessionID, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{Title: "Invalid session ID"}
	}

	// Delete proposal
	err = c.sessionService.DeleteProposal(ctx.Context(), int32(claims.UserID), int32(sessionID))
	if err != nil {
		if err.Error() == config.ErrSessionNotFound {
			return nil, fuego.NotFoundError{Title: err.Error()}
		}
		return nil, fuego.BadRequestError{Title: err.Error()}
	}

	return nil, nil
}

// ListSessions handles listing all conference sessions
func (c *SessionController) ListSessions(ctx fuego.ContextNoBody) ([]model.SessionWithDetails, error) {
	// Get user claims from context
	claims, ok := ctx.Value(jwt.ClaimsContextKey).(*jwt.Claims)
	if !ok {
		return nil, fuego.UnauthorizedError{Title: "Invalid token claims"}
	}

	// Get all sessions
	sessions, err := c.sessionService.ListSessions(ctx.Context(), int32(claims.UserID))
	if err != nil {
		return nil, fuego.InternalServerError{Title: err.Error()}
	}

	return sessions, nil
}

// GetSession handles getting a specific session details by ID
func (c *SessionController) GetSession(ctx fuego.ContextNoBody) (model.SessionWithDetails, error) {
	// Get user claims from context
	claims, ok := ctx.Value(jwt.ClaimsContextKey).(*jwt.Claims)
	if !ok {
		return model.SessionWithDetails{}, fuego.UnauthorizedError{Title: "Invalid token claims"}
	}

	// Get session ID from path and convert to int
	sessionID, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return model.SessionWithDetails{}, fuego.BadRequestError{Title: "Invalid session ID"}
	}

	// Get session details
	session, err := c.sessionService.GetSession(ctx.Context(), int32(sessionID), int32(claims.UserID))
	if err != nil {
		return model.SessionWithDetails{}, fuego.NotFoundError{Title: err.Error()}
	}

	return session, nil
}

// RegisterForSession handles registering a user for a specific session
func (c *SessionController) RegisterForSession(ctx fuego.ContextNoBody) (any, error) {
	// Get user claims from context
	claims, ok := ctx.Value(jwt.ClaimsContextKey).(*jwt.Claims)
	if !ok {
		return nil, fuego.UnauthorizedError{Title: "Invalid token claims"}
	}

	// Get session ID from path and convert to int
	sessionID, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{Title: "Invalid session ID"}
	}

	// Register user for the session
	err = c.sessionService.RegisterForSession(ctx.Context(), int32(sessionID), int32(claims.UserID))
	if err != nil {
		return nil, fuego.BadRequestError{Title: err.Error()}
	}

	return map[string]string{"message": "Successfully registered for session"}, nil
}

// CreateFeedback handles creating feedback for a specific session
func (c *SessionController) CreateFeedback(ctx fuego.ContextWithBody[model.SessionFeedback]) (model.SessionFeedbackResponse, error) {
	// Get user claims from context
	claims, ok := ctx.Value(jwt.ClaimsContextKey).(*jwt.Claims)
	if !ok {
		return model.SessionFeedbackResponse{}, fuego.UnauthorizedError{Title: "Invalid token claims"}
	}

	// Get session ID from path and convert to int
	sessionID, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return model.SessionFeedbackResponse{}, fuego.BadRequestError{Title: "Invalid session ID"}
	}

	// Get request body
	feedback, err := ctx.Body()
	if err != nil {
		return model.SessionFeedbackResponse{}, err
	}

	// Create feedback
	response, err := c.sessionService.CreateFeedback(ctx.Context(), int32(sessionID), int32(claims.UserID), feedback.Comment)
	if err != nil {
		if err.Error() == config.ErrDuplicateFeedback {
			return model.SessionFeedbackResponse{}, fuego.ConflictError{Title: err.Error()}
		}
		return model.SessionFeedbackResponse{}, fuego.BadRequestError{Title: err.Error()}
	}

	return response, nil
} 