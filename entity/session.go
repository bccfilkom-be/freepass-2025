package entity

import "time"

type Session struct {
	SessionID      int                   `json:"sessionID" gorm:"type:int;primary_key;autoIncrement"`
	Title          string                `json:"title" gorm:"type:varchar(50);not null"`
	Description    string                `json:"description" gorm:"type:text"`
	SessionOwner   string                `json:"sessinOwner" gorm:"type:varchar(100);not null"`
	TimeSlot       time.Time             `json:"timeSlots" gorm:"type:datetime;not null"`
	MaxSeats       int                   `json:"maxSeats" gorm:"type:int;not null"`
	AvailableSeats int                   `json:"availableSeats" gorm:"type:int;not null"`
	CreatedAt      time.Time             `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt      time.Time             `json:"updatedAt" gorm:"autoUpdateTime"`
	Proposal       SessionProposal       `json:"proposal" gorm:"foreignKey:SessionID"`
	ProposalID     int                   `json:"proposalID"`
	Feedbacks      []Feedback            `json:"feedbacks" gorm:"foreignKey:SessionID"`
	Registration   []SessionRegistration `json:"registration" gorm:"foreignKey:SessionID"`
}
