package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course/internal/register/usecase"
	"go_online_course/internal/user/dto"
	"go_online_course/pkg/utils"
	"net/http"
)

type RegisterHandler struct {
	registerUseCase usecase.RegisterUseCase
}

func NewRegisterHandler(registerUseCase usecase.RegisterUseCase) *RegisterHandler {
	return &RegisterHandler{registerUseCase}
}

func (rh *RegisterHandler) Route(r *gin.RouterGroup) {
	r.POST("/api/v1/register", rh.Register)
}

func (rh RegisterHandler) Register(ctx *gin.Context) {
	//	validate input
	var registerRequestInput dto.UserRequestBody

	//validation from body using JSON format
	if err := ctx.ShouldBindJSON(&registerRequestInput); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(400, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	err := rh.registerUseCase.Register(registerRequestInput)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(500, "internal sever error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, utils.Response(201, "created", "success, check your email"))
}
