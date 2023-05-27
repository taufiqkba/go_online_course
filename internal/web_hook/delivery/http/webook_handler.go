package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course/internal/web_hook/dto"
	"go_online_course/internal/web_hook/usecase"
	"go_online_course/pkg/utils"
	"net/http"
)

type WebHookHandler struct {
	useCase usecase.WebHookUseCase
}

func NewWebHookHandler(useCase usecase.WebHookUseCase) *WebHookHandler {
	return &WebHookHandler{useCase: useCase}
}

func (handler *WebHookHandler) Route(r *gin.RouterGroup) {
	webHookHandler := r.Group("/api/v1")

	webHookHandler.POST("/webhooks", handler.Xendit)
}

func (handler *WebHookHandler) Xendit(ctx *gin.Context) {
	var input dto.WebHookRequestBody

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	err = handler.useCase.UpdatePayment(input.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", "ok"))
}
