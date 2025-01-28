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