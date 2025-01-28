package dto

type GetAllSessionResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`

	RegistrationStartDate string `json:"registration_start_date"`
	RegistrationEndDate   string `json:"registration_end_date"`

	SessionStartDate string `json:"session_start_date"`
	SessionEndDate   string `json:"session_end_date"`

	MaxSeat int                   `json:"max_seat"`
	User    GetUserDetailResponse `json:"user"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateSessionRequest struct {
	Title                 string `json:"title" validate:"required"`
	Description           string `json:"description" validate:"required"`
	RegistrationStartDate string `json:"registration_start_date" validate:"required"`
	RegistrationEndDate   string `json:"registration_end_date" validate:"required"`

	SessionStartDate string `json:"session_start_date" validate:"required"`
	SessionEndDate   string `json:"session_end_date" validate:"required"`

	MaxSeat int `json:"max_seat" validate:"required"`
}
