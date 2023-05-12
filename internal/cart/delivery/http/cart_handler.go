package http

import (
	"go_online_course/internal/cart/dto"
	"go_online_course/internal/cart/usecase"
	"go_online_course/internal/middleware"
	"go_online_course/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	usecase usecase.CartUseCase
}

func NewCartHandler(usecase usecase.CartUseCase) *CartHandler {
	return &CartHandler{usecase}
}

func (handler *CartHandler) Route(r *gin.RouterGroup) {
	cartHandler := r.Group("/api/v1")

	cartHandler.Use(middleware.AuthJwt)
	{
		cartHandler.GET("/carts", handler.FindByUserID)
		cartHandler.POST("/carts", handler.Create)
		cartHandler.DELETE("/carts/:id", handler.Delete)
	}
}

func (handler *CartHandler) FindByUserID(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	user := utils.GetCurrentUser(ctx)

	data := handler.usecase.FindByUserID(int(user.ID), offset, limit)
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}

func (handler *CartHandler) Create(ctx *gin.Context) {
	var input dto.CartRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.UserID = user.ID

	data, err := handler.usecase.Create(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, utils.Response(http.StatusCreated, "created", data))

}

func (handler *CartHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	user := utils.GetCurrentUser(ctx)
	err := handler.usecase.Delete(id, int(user.ID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", "success"))
}
