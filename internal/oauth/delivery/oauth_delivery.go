package delivery

import (
	"go_online_course/internal/oauth/dto"
	"go_online_course/internal/oauth/usecase"
	"go_online_course/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OauthHandler struct {
	usecase usecase.OauthUseCase
}

func (handler *OauthHandler) Route(r *gin.RouterGroup) {
	oauthRouter := r.Group("/api/v1")
	oauthRouter.POST("/oauth", handler.Login)
}

func NewOauthHandler(usecase usecase.OauthUseCase) *OauthHandler {
	return &OauthHandler{usecase}
}

func (handler *OauthHandler) Login(ctx *gin.Context) {
	var input dto.LoginRequestBody
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	// call usecase from login
	data, err := handler.usecase.Login(input)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "OK", data))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "OK", data))
}
