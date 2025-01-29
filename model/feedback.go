package model

import (
	"time"
)

type FeedbackParam struct {
	Comment string `json:"comment" binding:"required"`
}

type FeedbackResponse struct {
	Comment string `json:"comment"`
}

type FeedbackResponseOnSession struct {
	ID        int       `json:"id"`
	SessionID int       `json:"session_id"`
	UserName  string    `json:"user_name"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
