package entity

import (
	"time"

	"github.com/google/uuid"
)

type SessionProposal struct {
	ProposalID  int       `json:"proposalID" gorm:"type:int;primary_key;autoIncrement"`
	Title       string    `json:"title" gorm:"type:varchar(50);not null"`
	Description string    `json:"description" gorm:"type:text"`
	TimeSlot    time.Time `json:"timeSlot" gorm:"datetime;not null"`
	Status      int       `json:"status" gorm:"int;not null"`
	MaxSeats    int       `json:"maxSeats" gorm:"int;not null"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	UserID      uuid.UUID `json:"userID"`
}
