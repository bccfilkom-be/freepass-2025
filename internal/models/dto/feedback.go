package dto

type CreateFeedbackRequest struct {
	Content string `json:"content" validate:"required"`
	Rating  int    `json:"rating" validate:"required,gte=1,lte=5"`
}
