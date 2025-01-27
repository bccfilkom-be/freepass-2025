package dto

type CreateProposalRequest struct {
	Title                 string `json:"title" validate:"required"`
	Description           string `json:"description" validate:"required"`
	RegistrationStartDate string `json:"registration_start_date" validate:"required"`
	RegistrationEndDate   string `json:"registration_end_date" validate:"required"`

	SessionStartDate string `json:"session_start_date" validate:"required"`
	SessionEndDate   string `json:"session_end_date" validate:"required"`

	MaxSeat int `json:"max_seat" validate:"required"`
}
