package controller

import (
	"strconv"

	"github.com/go-fuego/fuego"
	"github.com/litegral/freepass-2025/internal/model"
	"github.com/litegral/freepass-2025/internal/service"
)

// CoordinatorController handles coordinator-specific endpoints
type CoordinatorController struct {
	coordinatorService *service.CoordinatorService
}

func NewCoordinatorController(coordinatorService *service.CoordinatorService) *CoordinatorController {
	return &CoordinatorController{
		coordinatorService: coordinatorService,
	}
}

// ListProposals handles listing all session proposals
func (c *CoordinatorController) ListProposals(ctx fuego.ContextNoBody) ([]model.SessionWithDetails, error) {
	proposals, err := c.coordinatorService.ListSessionProposals(ctx.Context())
	if err != nil {
		return nil, fuego.InternalServerError{Title: err.Error()}
	}

	return proposals, nil
}

// UpdateProposalStatus handles updating a proposal's status
func (c *CoordinatorController) UpdateProposalStatus(ctx fuego.ContextWithBody[model.ProposalStatusUpdate]) (model.Session, error) {
	// Get session ID from path and convert to int
	sessionID, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return model.Session{}, fuego.BadRequestError{Title: "Invalid session ID"}
	}

	// Get request body
	update, err := ctx.Body()
	if err != nil {
		return model.Session{}, err
	}

	// Update status
	session, err := c.coordinatorService.UpdateSessionStatus(ctx.Context(), int32(sessionID), update.Status)
	if err != nil {
		return model.Session{}, fuego.BadRequestError{Title: err.Error()}
	}

	return session, nil
}

// RemoveSession handles removing a session
func (c *CoordinatorController) RemoveSession(ctx fuego.ContextNoBody) (any, error) {
	// Get session ID from path and convert to int
	sessionID, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{Title: "Invalid session ID"}
	}

	// Remove session
	err = c.coordinatorService.RemoveSession(ctx.Context(), int32(sessionID))
	if err != nil {
		return nil, fuego.InternalServerError{Title: err.Error()}
	}

	return nil, nil
}

// RemoveFeedback handles removing inappropriate feedback
func (c *CoordinatorController) RemoveFeedback(ctx fuego.ContextNoBody) (any, error) {
	// Get feedback ID from path and convert to int
	feedbackID, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{Title: "Invalid feedback ID"}
	}

	// Remove feedback
	err = c.coordinatorService.RemoveFeedback(ctx.Context(), int32(feedbackID))
	if err != nil {
		return nil, fuego.InternalServerError{Title: err.Error()}
	}

	return nil, nil
} 