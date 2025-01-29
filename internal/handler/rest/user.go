package rest

import (
	"freepass-bcc/model"
	"freepass-bcc/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) Register(ctx *gin.Context) {
	param := model.UserRegister{}

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	err = r.service.UserService.Register(&param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to register user", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "success to register new user", nil)

}

func (r *Rest) Login(ctx *gin.Context) {
	param := model.UserLogin{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	result, err := r.service.UserService.Login(param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to login user", err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully login to system", result)
}

func (r *Rest) UpdateProfile(ctx *gin.Context) {
	id := ctx.Param("userID")

	var profileRequest model.UpdateProfile
	err := ctx.ShouldBindJSON(&profileRequest)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}

	profile, err := r.service.UserService.UpdateProfile(id, &profileRequest)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to update profile", err)
		return
	}

	updateResponse := model.UpdateProfile{
		UserID:  profile.UserID,
		Name:    profile.Name,
		Email:   profile.Email,
		Address: profile.Address,
	}

	response.Success(ctx, http.StatusOK, "success to update profile", updateResponse)
}

func (r *Rest) AddNewEC(ctx *gin.Context) {
	id := ctx.Param("userID")

	err := r.service.UserService.AddNewEC(id)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to add new event coordinator", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success to add new event coordinator", nil)
}

func (r *Rest) RemoveRole(ctx *gin.Context) {
	id := ctx.Param("userID")

	err := r.service.UserService.RemoveRole(id)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to remove role", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success to remove role from user", nil)
}

func (r *Rest) GetUserByID(ctx *gin.Context) {
	id := ctx.Query("userID")

	user, err := r.service.UserService.GetUserByID(id)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get user", err)
		return
	}

	responses := model.SearchUser{
		Name:    user.Name,
		Address: user.Address,
	}

	response.Success(ctx, http.StatusOK, "user found", responses)
}
