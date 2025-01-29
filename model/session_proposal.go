package model

import (
	"time"

	"github.com/google/uuid"
)

type CreateProposal struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	TimeSlot    time.Time `json:"timeSlot" binding:"required"`
	MaxSeats    int       `json:"maxSeats" binding:"required"`
}

type CreateProposalResponse struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TimeSlot    time.Time `json:"timeSlot"`
	Status      int       `json:"status"`
	MaxSeats    int       `json:"maxSeats"`
	CreatedAt   time.Time `json:"createdAt"`
	UserName    string    `json:"userName"`
}

type GetAllProposal struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TimeSlot    time.Time `json:"timeSlot"`
	Status      int       `json:"status"`
	MaxSeats    int       `json:"maxSeats"`
	CreatedAt   time.Time `json:"createdAt"`
	UserID      uuid.UUID `json:"userID"`
}

type UpdateProposal struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	TimeSlot    *time.Time `json:"timeSlot"`
	MaxSeats    *int       `json:"maxSeats"`
}
