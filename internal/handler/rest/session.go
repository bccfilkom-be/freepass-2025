package rest

import (
	"freepass-bcc/entity"
	"freepass-bcc/model"
	"freepass-bcc/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Rest) GetAllSession(ctx *gin.Context) {
	pageQuery := ctx.Query("page")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "failed to bind request", err)
		return
	}

	session, err := r.service.SessionService.GetAllSession(page)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get session", err)
		return
	}

	var sessionResponse []model.GetSession

	for _, v := range *session {
		sessionResponse = append(sessionResponse, model.GetSession{
			SessionID:      v.SessionID,
			Title:          v.Title,
			Description:    v.Description,
			SessionOwner:   v.SessionOwner,
			TimeSlot:       v.TimeSlot,
			MaxSeats:       v.MaxSeats,
			AvailableSeats: v.AvailableSeats,
			CreatedAt:      v.CreatedAt,
		})
	}

	response.Success(ctx, http.StatusOK, "session retrieved", sessionResponse)
}

func (r *Rest) DeleteSession(ctx *gin.Context) {
	id := ctx.Param("sessionID")
	sessionID, err := strconv.Atoi(id)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid session id", err)
		return
	}

	err = r.service.SessionService.DeleteSession(sessionID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to delete session", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success to delete session", nil)
}

func (r *Rest) RegisterUserToSession(ctx *gin.Context) {
	id := ctx.Query("sessionID")
	sessionID, err := strconv.Atoi(id)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid session id", err)
		return
	}

	user := ctx.MustGet("user").(entity.User)
	regist, err := r.service.SessionService.RegisterUserToSession(user.UserID, sessionID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to register", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "success to register", regist)
}
