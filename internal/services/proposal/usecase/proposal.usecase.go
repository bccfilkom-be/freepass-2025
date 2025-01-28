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

type SessionDates struct {
	RegistrationStart time.Time
	RegistrationEnd   time.Time
	SessionStart      time.Time
	SessionEnd        time.Time
}

func NewProposalUsecase(sessionRepo session.SessionRepository) proposal.ProposalUsecase {
	return &ProposalUsecase{sessionRepo}
}

func (v *ProposalUsecase) CreateProposal(userId uint, req *dto.CreateProposalRequest) error {
	dates, err := parseDatesFromRequest(
		req.RegistrationStartDate,
		req.RegistrationEndDate,
		req.SessionStartDate,
		req.SessionEndDate,
	)

	if err != nil {
		return err
	}

	if err := validateDates(dates); err != nil {
		return err
	}

	if err := v.sessionRepo.DateInBetweenSession(dates.SessionStart, dates.SessionEnd, session.SessionFilter{
		UserID: userId,
	}); err != nil {
		return err
	}

	data := domain.Session{
		UserID:                userId,
		Title:                 req.Title,
		Description:           req.Description,
		RegistrationStartDate: dates.RegistrationStart,
		RegistrationEndDate:   dates.RegistrationEnd,

		SessionStartDate: dates.SessionStart,
		SessionEndDate:   dates.SessionEnd,
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

	dates, err := parseDatesFromRequest(
		req.RegistrationStartDate,
		req.RegistrationEndDate,
		req.SessionStartDate,
		req.SessionEndDate,
	)

	if err != nil {
		return err
	}

	if err := validateDates(dates); err != nil {
		return err
	}

	if err := v.sessionRepo.DateInBetweenSession(dates.SessionStart, dates.SessionEnd, session.SessionFilter{
		UserID:    userId,
		ExcludeID: []uint{sessionId},
	}); err != nil {
		return err
	}

	data := domain.Session{
		ID:                    sessionId,
		Title:                 req.Title,
		Description:           req.Description,
		RegistrationStartDate: dates.RegistrationStart,
		RegistrationEndDate:   dates.RegistrationEnd,

		SessionStartDate: dates.SessionStart,
		SessionEndDate:   dates.SessionEnd,
		MaxSeat:          req.MaxSeat,
		Status:           constant.STATUS_SESSION_PENDING,
	}

	return v.sessionRepo.Update(data)
}

func (v *ProposalUsecase) GetAllProposal(ctx *gin.Context) ([]dto.GetProposalResponse, error) {
	filter := session.SessionFilter{
		Status: []string{constant.STATUS_SESSION_PENDING, constant.STATUS_SESSION_REJECTED},
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

func (v *ProposalUsecase) ApproveProposal(sessionId uint) error {
	session, err := v.sessionRepo.GetById(sessionId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Proposal not found!")
		} else {
			return err
		}
	}

	if session.Status != constant.STATUS_SESSION_PENDING {
		return errors.New("Proposal not found! Either it is already accepted or rejected")
	}

	return v.sessionRepo.Update(domain.Session{
		ID:     sessionId,
		Status: constant.STATUS_SESSION_ACCEPTED,
	})
}

func (v *ProposalUsecase) DeclineProposal(sessionId uint, req *dto.DecliendProposalRequest) error {
	session, err := v.sessionRepo.GetById(sessionId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Proposal not found!")
		} else {
			return err
		}
	}

	if session.Status != constant.STATUS_SESSION_PENDING {
		return errors.New("Proposal not found! Either it is already accepted or rejected")
	}

	return v.sessionRepo.Update(domain.Session{
		ID:              sessionId,
		Status:          constant.STATUS_SESSION_REJECTED,
		RejectedMessage: req.RejectedMessage,
	})
}

// Helper
func validateDates(dates SessionDates) error {
	if dates.RegistrationStart.Before(time.Now()) {
		return errors.New("registration start date should be after today")
	}

	if dates.RegistrationStart.After(dates.RegistrationEnd) {
		return errors.New("registration start date should be before the registration end date")
	}

	if dates.SessionStart.Before(dates.RegistrationEnd) {
		return errors.New("session start date should be after the registration end date")
	}

	if dates.SessionStart.After(dates.SessionEnd) {
		return errors.New("session start date should be before the session end date")
	}

	return nil
}

func parseDatesFromRequest(
	RegistrationStarts, RegistrationEnds, SessionStarts, SessionEnds string,
) (SessionDates, error) {
	registrationStart, err := helper.StringISOToDateTime(RegistrationStarts)
	if err != nil {
		return SessionDates{}, err
	}

	registrationEnd, err := helper.StringISOToDateTime(RegistrationEnds)
	if err != nil {
		return SessionDates{}, err
	}

	sessionStart, err := helper.StringISOToDateTime(SessionStarts)
	if err != nil {
		return SessionDates{}, err
	}

	sessionEnd, err := helper.StringISOToDateTime(SessionEnds)
	if err != nil {
		return SessionDates{}, err
	}

	return SessionDates{
		RegistrationStart: registrationStart,
		RegistrationEnd:   registrationEnd,
		SessionStart:      sessionStart,
		SessionEnd:        sessionEnd,
	}, nil
}
