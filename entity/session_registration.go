package entity

import (
	"time"

	"github.com/google/uuid"
)

type SessionRegistration struct {
	RegistrationID int       `json:"registrarionID" gorm:"type:int;primary_key;autoIncrement"`
	Timestamp      time.Time `json:"timestamp" gorm:"autoCreateTime"`
	SessionID      int       `json:"sessionID"`
	UserID         uuid.UUID `json:"userID"`
}
