package model

// ProposalStatusUpdate represents the request body for updating a proposal's status
type ProposalStatusUpdate struct {
	Status string `json:"status" validate:"required,oneof=accepted rejected"`
}
