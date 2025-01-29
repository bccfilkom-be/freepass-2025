package service

import (
	"freepass-bcc/internal/repository"
	"freepass-bcc/pkg/bcrypt"
	"freepass-bcc/pkg/jwt"
)

type Service struct {
	UserService     IUserService
	SessionService  ISessionService
	FeedbackService IFeedbackService
	ProposalService IProposalService
}

func NewService(repository *repository.Repository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface) *Service {
	return &Service{
		UserService:     NewUserService(repository.UserRepository, bcrypt, jwtAuth),
		SessionService:  NewSessionService(repository.SessionRepository, repository.RegisterRepository),
		FeedbackService: NewFeedbackService(repository.FeedbackRepository),
		ProposalService: NewProposalService(repository.ProposalRepository),
	}
}
