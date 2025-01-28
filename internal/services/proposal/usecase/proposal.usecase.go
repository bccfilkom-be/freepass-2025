package usecase

import (
	"errors"
	"jevvonn/bcc-be-freepass-2025/internal/constant"
	"jevvonn/bcc-be-freepass-2025/internal/helper"
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"
	"jevvonn/bcc-be-freepass-2025/internal/services/proposal"
	"jevvonn/bcc-be-freepass-2025/internal/services/session"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProposalUsecase struct {
	sessionRepo session.SessionRepository
}

func NewProposalUsecase(sessionRepo session.SessionRepository) proposal.ProposalUsecase {
	return &ProposalUsecase{sessionRepo}
}

func (v *ProposalUsecase) CreateProposal(userId uint, req *dto.CreateProposalRequest) error {
	registrationStarDate, err := helper.StringISOToDateTime(req.RegistrationStartDate)
	if err != nil {
		return err
	}

	registrationEndDate, err := helper.StringISOToDateTime(req.RegistrationEndDate)
	if err != nil {
		return err
	}

	sessionStartDate, err := helper.StringISOToDateTime(req.SessionStartDate)
	if err != nil {
		return err
	}

	sessionEndDate, err := helper.StringISOToDateTime(req.SessionEndDate)
	if err != nil {
		return err
	}

	if registrationStarDate.Before(time.Now()) {
		return errors.New("Registration start date should be after today!")
	}

	if registrationStarDate.After(registrationEndDate) {
		return errors.New("Registration start date should be before the registration end date!")
	}

	if sessionStartDate.Before(registrationEndDate) {
		return errors.New("Session start date should be after the registration end date!")
	}

	if sessionStartDate.After(sessionEndDate) {
		return errors.New("Session start date should be before the session end date!")
	}

	if err := v.sessionRepo.DateInBetweenSession(sessionStartDate, sessionEndDate, session.SessionFilter{
		Status: constant.STATUS_SESSION_PENDING,
		UserID: userId,
	}); err != nil {
		return err
	}

	data := domain.Session{
		UserID:                userId,
		Title:                 req.Title,
		Description:           req.Description,
		RegistrationStartDate: registrationStarDate,
		RegistrationEndDate:   registrationEndDate,

		SessionStartDate: sessionStartDate,
		SessionEndDate:   sessionEndDate,
		MaxSeat:          req.MaxSeat,
	}

	return v.sessionRepo.Create(data)
}

func (v *ProposalUsecase) UpdateProposal(sessionId, userId uint, req *dto.UpdateProposalRequest) error {
	sessionExist, err := v.sessionRepo.GetById(sessionId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Session not found!")
		} else {
			return err
		}
	}

	if sessionExist.UserID != userId {
		return errors.New("You are not authorized to update this session!")
	}

	registrationStarDate, err := helper.StringISOToDateTime(req.RegistrationStartDate)
	if err != nil {
		return err
	}

	registrationEndDate, err := helper.StringISOToDateTime(req.RegistrationEndDate)
	if err != nil {
		return err
	}

	sessionStartDate, err := helper.StringISOToDateTime(req.SessionStartDate)
	if err != nil {
		return err
	}

	sessionEndDate, err := helper.StringISOToDateTime(req.SessionEndDate)
	if err != nil {
		return err
	}

	if registrationStarDate.Before(time.Now()) {
		return errors.New("Registration start date should be after today!")
	}

	if registrationStarDate.After(registrationEndDate) {
		return errors.New("Registration start date should be before the registration end date!")
	}

	if sessionStartDate.Before(registrationEndDate) {
		return errors.New("Session start date should be after the registration end date!")
	}

	if sessionStartDate.After(sessionEndDate) {
		return errors.New("Session start date should be before the session end date!")
	}

	if err := v.sessionRepo.DateInBetweenSession(sessionStartDate, sessionEndDate, session.SessionFilter{
		Status:    constant.STATUS_SESSION_PENDING,
		UserID:    userId,
		ExcludeID: []uint{sessionId, 2},
	}); err != nil {
		return err
	}

	data := domain.Session{
		ID:                    sessionId,
		Title:                 req.Title,
		Description:           req.Description,
		RegistrationStartDate: registrationStarDate,
		RegistrationEndDate:   registrationEndDate,

		SessionStartDate: sessionStartDate,
		SessionEndDate:   sessionEndDate,
		MaxSeat:          req.MaxSeat,
	}

	return v.sessionRepo.Update(data)
}

func (v *ProposalUsecase) GetAllProposal(ctx *gin.Context) ([]dto.GetProposalResponse, error) {
	filter := session.SessionFilter{
		Status: constant.STATUS_SESSION_PENDING,
	}

	role := ctx.GetString("role")

	if role == constant.ROLE_USER {
		filter.UserID = ctx.GetUint("userId")
	}

	sessions, err := v.sessionRepo.GetAll(filter)
	if err != nil {
		return []dto.GetProposalResponse{}, err
	}

	var proposals []dto.GetProposalResponse
	for _, session := range sessions {
		proposals = append(proposals, dto.GetProposalResponse{
			ID:                    session.ID,
			Title:                 session.Title,
			Description:           session.Description,
			RegistrationStartDate: session.RegistrationStartDate.Format(time.RFC3339),
			RegistrationEndDate:   session.RegistrationEndDate.Format(time.RFC3339),

			SessionStartDate: session.SessionStartDate.Format(time.RFC3339),
			SessionEndDate:   session.SessionEndDate.Format(time.RFC3339),

			MaxSeat:         session.MaxSeat,
			Status:          session.Status,
			RejectedMessage: session.RejectedMessage,

			User: dto.GetUserDetailResponse{
				ID:    session.User.ID,
				Name:  session.User.Name,
				Email: session.User.Email,
			},

			CreatedAt: session.CreatedAt.Format(time.RFC3339),
			UpdatedAt: session.UpdatedAt.Format(time.RFC3339),
		})
	}

	return proposals, nil
}

func (v *ProposalUsecase) GetProposalDetail(ctx *gin.Context, sessionId uint) (dto.GetProposalResponse, error) {
	userId := ctx.GetUint("userId")
	role := ctx.GetString("role")

	session, err := v.sessionRepo.GetById(sessionId)
	if err != nil {
		return dto.GetProposalResponse{}, err
	}

	if session.UserID != userId && role != constant.ROLE_COORDINATOR {
		return dto.GetProposalResponse{}, errors.New("You are not authorized to view this proposal!")
	}

	if session.Status != constant.STATUS_SESSION_PENDING {
		return dto.GetProposalResponse{}, errors.New("Proposal not found!")
	}

	proposal := dto.GetProposalResponse{
		ID:                    session.ID,
		Title:                 session.Title,
		Description:           session.Description,
		RegistrationStartDate: session.RegistrationStartDate.Format(time.RFC3339),
		RegistrationEndDate:   session.RegistrationEndDate.Format(time.RFC3339),

		SessionStartDate: session.SessionStartDate.Format(time.RFC3339),
		SessionEndDate:   session.SessionEndDate.Format(time.RFC3339),

		MaxSeat:         session.MaxSeat,
		Status:          session.Status,
		RejectedMessage: session.RejectedMessage,

		User: dto.GetUserDetailResponse{
			ID:    session.User.ID,
			Name:  session.User.Name,
			Email: session.User.Email,
		},

		CreatedAt: session.CreatedAt.Format(time.RFC3339),
		UpdatedAt: session.UpdatedAt.Format(time.RFC3339),
	}

	return proposal, nil
}

func (v *ProposalUsecase) DeleteProposal(ctx *gin.Context, sessionId uint) error {
	userId := ctx.GetUint("userId")

	session, err := v.sessionRepo.GetById(sessionId)
	if err != nil {
		return err
	}

	if session.UserID != userId {
		return errors.New("You are not authorized to delete this proposal!")
	}

	if session.Status != constant.STATUS_SESSION_PENDING {
		return errors.New("Proposal not found!")
	}

	return v.sessionRepo.Delete(sessionId)
}
