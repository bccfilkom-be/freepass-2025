package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseHandler interface {
	OK(ctx *gin.Context, data interface{}, message string, code int)
	BadRequest(ctx *gin.Context, data interface{}, message string)
	NotFound(ctx *gin.Context)
	Forbidden(ctx *gin.Context, message string)
	InternalServerError(ctx *gin.Context, message string)
}

type JSONResponseModel struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func NewResponseHandler() ResponseHandler {
	return &JSONResponseModel{}
}

func (r *JSONResponseModel) OK(ctx *gin.Context, data interface{}, message string, code int) {
	ctx.JSON(code, JSONResponseModel{
		Message: message,
		Success: true,
		Data:    data,
	})
}

func (r *JSONResponseModel) BadRequest(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusBadRequest, JSONResponseModel{
		Message: message,
		Success: false,
		Data:    data,
	})
}

func (r *JSONResponseModel) InternalServerError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, JSONResponseModel{
		Message: fmt.Sprintf("Internal Server Error: %s", message),
	})
}

func (r *JSONResponseModel) Forbidden(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusForbidden, JSONResponseModel{
		Message: fmt.Sprintf("Forbidden: %s", message),
	})
}

func (r *JSONResponseModel) NotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, JSONResponseModel{
		Message: "Not Found!",
	})
}
