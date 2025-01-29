package service

import (
	"fmt"
	"freepass-bcc/entity"
	"freepass-bcc/internal/repository"
	"freepass-bcc/model"

	"github.com/google/uuid"
)

type IProposalService interface {
	CreateProposal(userID uuid.UUID, param model.CreateProposal) (*entity.SessionProposal, error)
	UpdateProposal(id int, request *model.UpdateProposal) (*entity.SessionProposal, error)
	GetAllProposal(page int) ([]*entity.SessionProposal, error)
	ApprovedProposal(proposalID int) (*entity.Session, error)
	RejectedProsal(proposalID int) error
	DeleteProposal(id int) error
}

type ProposalService struct {
	ProposalRepository repository.IProposalRepository
}

func NewProposalService(proposalRepository repository.IProposalRepository) IProposalService {
	return &ProposalService{proposalRepository}
}

func (ps *ProposalService) CreateProposal(userID uuid.UUID, param model.CreateProposal) (*entity.SessionProposal, error) {
	hasCreatedToday, err := ps.ProposalRepository.HasCreatedProposalToday(userID)
	if err != nil {
		return nil, err
	}

	if hasCreatedToday {
		return nil, fmt.Errorf("already created a proposal today")
	}

	proposal := &entity.SessionProposal{
		Title:       param.Title,
		Description: param.Description,
		TimeSlot:    param.TimeSlot,
		MaxSeats:    param.MaxSeats,
		Status:      0,
		UserID:      userID,
	}

	err = ps.ProposalRepository.CreateProposal(proposal)
	if err != nil {
		return nil, err
	}

	return proposal, nil
}

func (ps *ProposalService) UpdateProposal(id int, request *model.UpdateProposal) (*entity.SessionProposal, error) {
	proposal, err := ps.ProposalRepository.UpdateProposal(id, request)
	if err != nil {
		return nil, err
	}

	return proposal, nil
}

func (ps *ProposalService) DeleteProposal(id int) error {
	err := ps.ProposalRepository.DeleteProposal(id)
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProposalService) GetAllProposal(page int) ([]*entity.SessionProposal, error) {
	limit := 2
	offset := (page - 1) * limit

	proposal, err := ps.ProposalRepository.GetAllProposal(limit, offset)
	if err != nil {
		return nil, err
	}

	return proposal, nil
}

func (ps *ProposalService) ApprovedProposal(proposalID int) (*entity.Session, error) {
	proposal, err := ps.ProposalRepository.GetProposalByID(proposalID)
	if err != nil {
		return nil, err
	}

	if proposal.Status != 0 {
		return nil, fmt.Errorf("proposal has been processed")
	}

	session, err := ps.ProposalRepository.CreateSessionFromProposal(proposal)
	if err != nil {
		return nil, err
	}

	err = ps.ProposalRepository.UpdateProposalStatus(proposalID, 1)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (ps *ProposalService) RejectedProsal(proposalID int) error {
	proposal, err := ps.ProposalRepository.GetProposalByID(proposalID)
	if err != nil {
		return err
	}

	if proposal.Status != 0 {
		return fmt.Errorf("proposal has been processed")
	}

	err = ps.ProposalRepository.UpdateProposalStatus(proposalID, 2)
	if err != nil {
		return err
	}

	return nil
}
