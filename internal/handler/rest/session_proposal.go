package rest

import (
	"freepass-bcc/entity"
	"freepass-bcc/model"
	"freepass-bcc/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateProposal(ctx *gin.Context) {
	var param model.CreateProposal
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid input", err)
		return
	}

	user := ctx.MustGet("user").(entity.User)
	userID := user.UserID

	proposal, err := r.service.ProposalService.CreateProposal(userID, param)
	if err != nil {
		if err.Error() == "already created a proposal today" {
			response.Error(ctx, http.StatusForbidden, "you can only create one proposal per day", err)
			return
		}

		response.Error(ctx, http.StatusInternalServerError, "failed to create proposal", err)
		return
	}

	responses := model.CreateProposalResponse{
		Title:       proposal.Title,
		Description: proposal.Description,
		TimeSlot:    proposal.TimeSlot,
		Status:      proposal.Status,
		MaxSeats:    proposal.MaxSeats,
		CreatedAt:   proposal.CreatedAt,
		UserName:    user.Name,
	}

	response.Success(ctx, http.StatusCreated, "session proposal created successfully", responses)

}

func (r *Rest) UpdateProposal(ctx *gin.Context) {
	id := ctx.Query("proposalID")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid id", err)
		return
	}

	var param model.UpdateProposal
	err = ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid input", err)
		return
	}

	proposal, err := r.service.ProposalService.UpdateProposal(idInt, &param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to update proposal", err)
		return
	}

	responses := model.UpdateProposal{
		Title:       proposal.Title,
		Description: proposal.Description,
		TimeSlot:    &proposal.TimeSlot,
		MaxSeats:    &proposal.MaxSeats,
	}

	response.Success(ctx, http.StatusOK, "success to update proposal", responses)
}

func (r *Rest) DeleteProposal(ctx *gin.Context) {
	id := ctx.Query("proposalID")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid proposal id", err)
		return
	}

	err = r.service.ProposalService.DeleteProposal(idInt)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to delete proposal", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success to delete proposal", nil)
}

func (r *Rest) GetAllProposal(ctx *gin.Context) {
	page := ctx.Query("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "failed to bind input", err)
		return
	}

	proposal, err := r.service.ProposalService.GetAllProposal(pageInt)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get proposal", err)
		return
	}

	var responses []model.GetAllProposal

	for _, v := range proposal {
		responses = append(responses, model.GetAllProposal{
			Title:       v.Title,
			Description: v.Description,
			TimeSlot:    v.TimeSlot,
			Status:      v.Status,
			MaxSeats:    v.MaxSeats,
			CreatedAt:   v.CreatedAt,
			UserID:      v.UserID,
		})
	}

	response.Success(ctx, http.StatusOK, "success to get proposal", responses)

}

func (r *Rest) ApprovedProposal(ctx *gin.Context) {
	id := ctx.Param("proposalID")
	proposalID, err := strconv.Atoi(id)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid proposal id", err)
		return
	}

	session, err := r.service.ProposalService.ApprovedProposal(proposalID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to approve proposal", err)
		return
	}

	responses := model.GetSession{
		SessionID:      session.SessionID,
		Title:          session.Title,
		Description:    session.Description,
		SessionOwner:   session.SessionOwner,
		TimeSlot:       session.TimeSlot,
		MaxSeats:       session.MaxSeats,
		AvailableSeats: session.AvailableSeats,
		CreatedAt:      session.CreatedAt,
	}

	response.Success(ctx, http.StatusOK, "proposal approved", responses)
}

func (r *Rest) RejectedProsal(ctx *gin.Context) {
	id := ctx.Param("proposalID")
	proposalID, err := strconv.Atoi(id)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid proposal id", err)
		return
	}

	err = r.service.ProposalService.RejectedProsal(proposalID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to reject proposal", err)
		return
	}

	response.Success(ctx, http.StatusOK, "proposal rejected", nil)
}
