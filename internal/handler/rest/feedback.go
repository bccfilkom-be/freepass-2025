package rest

import (
	"freepass-bcc/entity"
	"freepass-bcc/model"
	"freepass-bcc/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Rest) AddFeedback(ctx *gin.Context) {
	var comment model.FeedbackParam
	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	stringID := ctx.Query("sessionID")
	sessionID, err := strconv.Atoi(stringID)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid session id", err)
		return
	}

	user := ctx.MustGet("user").(entity.User)
	feedback, err := r.service.FeedbackService.AddFeedback(user.UserID, sessionID, comment.Comment)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to add feedback", err)
		return
	}

	responses := model.FeedbackResponse{
		Comment: feedback.Comment,
	}

	response.Success(ctx, http.StatusCreated, "feedback added successfully", responses)
}

func (r *Rest) GetFeedbackBySessionID(ctx *gin.Context) {
	stringID := ctx.Query("session")
	sessionID, err := strconv.Atoi(stringID)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid session id", err)
		return
	}

	feedbacks, err := r.service.FeedbackService.GetFeedbackBySessionID(sessionID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get feedbacks from session", err)
		return
	}

	var responses []model.FeedbackResponseOnSession
	for _, v := range feedbacks {
		responses = append(responses, model.FeedbackResponseOnSession{
			ID:        v.FeedbackID,
			SessionID: v.SessionID,
			UserName:  v.User.Name,
			Comment:   v.Comment,
			CreatedAt: v.CreatedAt,
		})
	}

	response.Success(ctx, http.StatusOK, "success to get feedback", responses)

}

func (r *Rest) DeleteFeedbackInSession(ctx *gin.Context) {
	feedbackID := ctx.Query("feedbackID")

	feedbackInt, err := strconv.Atoi(feedbackID)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid feedback id", err)
		return
	}

	err = r.service.FeedbackService.DeleteFeedbackInSession(feedbackInt)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to delete feedback", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success to delete feedback", nil)

}
