package repository

import (
	"freepass-bcc/entity"
	"freepass-bcc/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IProposalRepository interface {
	CreateProposal(proposal *entity.SessionProposal) error
	UpdateProposal(id int, request *model.UpdateProposal) (*entity.SessionProposal, error)
	DeleteProposal(id int) error
	GetAllProposal(limit, offset int) ([]*entity.SessionProposal, error)
	UpdateProposalStatus(proposalID, status int) error
	GetProposalByID(proposalID int) (*entity.SessionProposal, error)
	CreateSessionFromProposal(proposal *entity.SessionProposal) (*entity.Session, error)
	HasCreatedProposalToday(userID uuid.UUID) (bool, error)
}

type ProposalRepository struct {
	db *gorm.DB
}

func NewProposalRepository(db *gorm.DB) IProposalRepository {
	return &ProposalRepository{db}
}

func (pr *ProposalRepository) CreateProposal(proposal *entity.SessionProposal) error {
	err := pr.db.Debug().Create(&proposal).Error
	if err != nil {
		return err
	}

	return err
}

func (pr *ProposalRepository) UpdateProposal(id int, request *model.UpdateProposal) (*entity.SessionProposal, error) {
	tx := pr.db.Begin()
	var proposal entity.SessionProposal

	proposalUpdate := parseUpdateProposal(request, &proposal)
	err := tx.Debug().Where("proposal_id = ?", id).Updates(proposalUpdate).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &proposal, nil

}

func (pr *ProposalRepository) DeleteProposal(id int) error {
	tx := pr.db.Begin()
	var proposal entity.SessionProposal

	err := tx.Debug().Where("proposal_id = ?", id).First(&proposal).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Debug().Where("proposal_id = ?", id).Delete(proposal).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (pr *ProposalRepository) HasCreatedProposalToday(userID uuid.UUID) (bool, error) {
	var (
		count    int64
		proposal entity.SessionProposal
	)

	today := time.Now().Truncate(24 * time.Hour)
	err := pr.db.Debug().Model(&proposal).Where("user_id = ? AND created_at >= ?", userID, today).Count(&count).Error
	if err != nil {
		return false, nil
	}

	return count > 0, nil
}

func (pr *ProposalRepository) GetAllProposal(limit, offset int) ([]*entity.SessionProposal, error) {
	var proposal []*entity.SessionProposal

	err := pr.db.Debug().Limit(limit).Offset(offset).Find(&proposal).Error
	if err != nil {
		return nil, err
	}

	return proposal, nil
}

func (pr *ProposalRepository) UpdateProposalStatus(proposalID, status int) error {
	var proposal entity.SessionProposal

	err := pr.db.Debug().Model(&proposal).Where("proposal_id = ?", proposalID).Update("status", status).Error
	if err != nil {
		return err
	}

	return nil
}

func (pr *ProposalRepository) GetProposalByID(proposalID int) (*entity.SessionProposal, error) {
	var proposal entity.SessionProposal

	err := pr.db.Where("proposal_id = ?", proposalID).First(&proposal).Error
	if err != nil {
		return nil, err
	}

	return &proposal, nil
}

func (pr *ProposalRepository) CreateSessionFromProposal(proposal *entity.SessionProposal) (*entity.Session, error) {
	session := entity.Session{
		Title:          proposal.Title,
		Description:    proposal.Description,
		SessionOwner:   proposal.UserID.String(),
		TimeSlot:       proposal.TimeSlot,
		MaxSeats:       proposal.MaxSeats,
		AvailableSeats: proposal.MaxSeats,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		ProposalID:     proposal.ProposalID,
	}

	err := pr.db.Debug().Create(&session).Error
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func parseUpdateProposal(model *model.UpdateProposal, proposal *entity.SessionProposal) *entity.SessionProposal {
	if model.Title != "" {
		proposal.Title = model.Title
	}

	if model.Description != "" {
		proposal.Description = model.Description
	}

	if model.MaxSeats != nil {
		proposal.MaxSeats = *model.MaxSeats
	}

	if model.TimeSlot != nil {
		proposal.TimeSlot = *model.TimeSlot
	}

	return proposal
}
