package dto

import "time"

type GetSessionRegistrationResponse struct {
	ID        uint
	SessionID uint
	Session   GetAllSessionResponse

	CreatedAt time.Time
	UpdatedAt time.Time
}
