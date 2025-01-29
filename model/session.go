package model

import "time"

type GetSession struct {
	SessionID      int       `json:"sessionID"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	SessionOwner   string    `json:"sessinOwner"`
	TimeSlot       time.Time `json:"timeSlots"`
	MaxSeats       int       `json:"maxSeats"`
	AvailableSeats int       `json:"availableSeats"`
	CreatedAt      time.Time `json:"createdAt"`
}
