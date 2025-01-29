package dto

type CreateProposalRequest struct {
	Title                 string `json:"title" validate:"required" example:"Tech Innovations Conference 2025"`
	Description           string `json:"description" validate:"required" example:"A conference focused on the latest advancements in technology and innovation."`
	RegistrationStartDate string `json:"registration_start_date" validate:"required" example:"2025-02-01T09:00:00Z"`
	RegistrationEndDate   string `json:"registration_end_date" validate:"required" example:"2025-02-15T12:00:00Z"`

	SessionStartDate string `json:"session_start_date" validate:"required" example:"2025-02-25T12:00:00Z"`
	SessionEndDate   string `json:"session_end_date" validate:"required" example:"2025-02-25T17:00:00Z"`

	MaxSeat int `json:"max_seat" validate:"required" example:"100"`
}

type UpdateProposalRequest struct {
	Title                 string `json:"title" validate:"required" example:"Tech Innovations Conference 2026"`
	Description           string `json:"description" validate:"required" example:"A conference focused on the latest advancements in technology and innovation."`
	RegistrationStartDate string `json:"registration_start_date" validate:"required" example:"2026-02-01T09:00:00Z"`
	RegistrationEndDate   string `json:"registration_end_date" validate:"required" example:"2026-02-15T12:00:00Z"`

	SessionStartDate string `json:"session_start_date" validate:"required" example:"2026-02-25T12:00:00Z"`
	SessionEndDate   string `json:"session_end_date" validate:"required" example:"2026-02-25T17:00:00Z"`

	MaxSeat int `json:"max_seat" validate:"required" example:"50"`
}

type GetProposalResponse struct {
	ID                    uint   `json:"id"`
	Title                 string `json:"title"`
	Description           string `json:"description"`
	RegistrationStartDate string `json:"registration_start_date"`
	RegistrationEndDate   string `json:"registration_end_date"`

	SessionStartDate string `json:"session_start_date"`
	SessionEndDate   string `json:"session_end_date"`

	MaxSeat         int    `json:"max_seat"`
	Status          string `json:"status"`
	RejectedMessage string `json:"rejected_message"`

	User GetUserDetailResponse `json:"user"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DecliendProposalRequest struct {
	RejectedMessage string `json:"rejected_message" validate:"required" example:"The title is not clear."`
}
