package http

import (
	"go_online_course/internal/middleware"
	"go_online_course/internal/profile/usecase"
	"go_online_course/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	useCase usecase.ProfileUseCase
}

func NewProfileHandler(useCase usecase.ProfileUseCase) *ProfileHandler {
	return &ProfileHandler{useCase}
}

func (handler *ProfileHandler) Route(r *gin.RouterGroup) {
	authorized := r.Group("/api/v1")
	authorized.Use(middleware.AuthJwt)
	{
		authorized.GET("/profiles", handler.GetProfile)
	}
}

func (handler *ProfileHandler) GetProfile(ctx *gin.Context) {
	user := utils.GetCurrentUser(ctx)

	//	Get Profile user
	profile, err := handler.useCase.GetProfile(int(user.ID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", profile))
}
