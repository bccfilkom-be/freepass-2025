package dto

type CreateFeedbackRequest struct {
	Content string `json:"content" validate:"required"`
}
