package dto

import "time"

type GetSessionRegistrationResponse struct {
	ID        uint                  `json:"id"`
	SessionID uint                  `json:"session_id"`
	Session   GetAllSessionResponse `json:"session"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
