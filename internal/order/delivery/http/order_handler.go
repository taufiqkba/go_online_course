package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course/internal/middleware"
	"go_online_course/internal/order/dto"
	"go_online_course/internal/order/usecase"
	"go_online_course/pkg/utils"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	useCase usecase.OrderUseCase
}

func NewOrderHandler(useCase usecase.OrderUseCase) *OrderHandler {
	return &OrderHandler{useCase: useCase}
}

func (handler *OrderHandler) Route(r *gin.RouterGroup) {
	orderHandler := r.Group("/api/v1")
	orderHandler.Use(middleware.AuthJwt)
	{
		orderHandler.POST("/orders", handler.Create)
		orderHandler.GET("/orders", handler.FindAllByUserID)

	}
}
func (handler *OrderHandler) Create(ctx *gin.Context) {
	var input dto.OrderRequestBody
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	//set data current user
	user := utils.GetCurrentUser(ctx)
	input.UserID = user.ID
	input.Email = user.Email

	//	call create
	data, err := handler.useCase.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, utils.Response(http.StatusCreated, "created", data))

}

func (handler *OrderHandler) FindAllByUserID(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Param("offset"))
	limit, _ := strconv.Atoi(ctx.Param("limit"))

	user := utils.GetCurrentUser(ctx)
	data := handler.useCase.FindAllByUserID(offset, limit, int(user.ID))

	//return data
	ctx.JSON(http.StatusCreated, utils.Response(http.StatusCreated, "created", data))
}

func (handler *OrderHandler) FindByID(ctx *gin.Context) {
	//	TODO Implement me
	panic("unimplemented")
}
