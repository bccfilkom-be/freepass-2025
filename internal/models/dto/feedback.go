package dto

type CreateFeedbackRequest struct {
	Content string `json:"content" validate:"required"`
	Rating  int    `json:"rating" validate:"required,gte=1,lte=5"`
}

type GetFeedbackResponse struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	Rating    int    `json:"rating"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`

	User GetUserDetailResponse `json:"user"`
}
