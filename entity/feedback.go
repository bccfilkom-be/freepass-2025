package entity

import (
	"time"

	"github.com/google/uuid"
)

type Feedback struct {
	FeedbackID int       `json:"feedbackID" gorm:"type:int;primary_key;autoIncrement"`
	Comment    string    `json:"comment" gorm:"type:varchar(255);not null"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"`
	SessionID  int       `json:"sessionID"`
	User       User      `json:"user"`
	UserID     uuid.UUID `json:"userID"`
}
