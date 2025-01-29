package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository     IUserRepository
	SessionRepository  ISessionRepository
	FeedbackRepository IFeedbackRepository
	ProposalRepository IProposalRepository
	RegisterRepository IRegisterRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:     NewUserRepository(db),
		SessionRepository:  NewSessionRepository(db),
		FeedbackRepository: NewFeedbackRepository(db),
		ProposalRepository: NewProposalRepository(db),
		RegisterRepository: NewRegisterRepository(db),
	}
}
