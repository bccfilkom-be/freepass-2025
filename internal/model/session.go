package model

import "time"

type Session struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	Room            string    `json:"room"`
	Status          string    `json:"status"`
	SeatingCapacity int       `json:"seating_capacity"`
	ProposerID      int       `json:"proposer_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type SessionProposal struct {
	Title           string    `json:"title" validate:"required"`
	Description     string    `json:"description" validate:"required"`
	StartTime       time.Time `json:"start_time" validate:"required"`
	EndTime         time.Time `json:"end_time" validate:"required,gtfield=StartTime"`
	Room            string    `json:"room" validate:"required"`
	SeatingCapacity int       `json:"seating_capacity" validate:"required,min=1"`
}

type SessionUpdate struct {
	Title           string    `json:"title" validate:"required"`
	Description     string    `json:"description" validate:"required"`
	StartTime       time.Time `json:"start_time" validate:"required"`
	EndTime         time.Time `json:"end_time" validate:"required,gtfield=StartTime"`
	Room            string    `json:"room" validate:"required"`
	SeatingCapacity int       `json:"seating_capacity" validate:"required,min=1"`
} 