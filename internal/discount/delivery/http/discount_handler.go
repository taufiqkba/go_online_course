package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course/internal/discount/dto"
	"go_online_course/internal/discount/usecase"
	"go_online_course/internal/middleware"
	"go_online_course/pkg/utils"
	"net/http"
	"strconv"
)

type DiscountHandler struct {
	useCase usecase.DiscountUseCase
}

func NewDiscountHandler(useCase usecase.DiscountUseCase) *DiscountHandler {
	return &DiscountHandler{useCase: useCase}
}

func (handler *DiscountHandler) Route(r *gin.RouterGroup) {
	discountHandler := r.Group("/api/v1")

	discountHandler.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		discountHandler.GET("/discounts", handler.FindAll)
		discountHandler.GET("/discounts/:id", handler.FindByID)
		discountHandler.POST("/discounts", handler.Create)
		discountHandler.PATCH("/discounts/:id", handler.Update)
		discountHandler.DELETE("/discounts/:id", handler.Delete)
	}
}

func (handler *DiscountHandler) Create(ctx *gin.Context) {
	var input dto.DiscountRequestBody
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.CreatedBy = user.ID
	data, err := handler.useCase.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, utils.Response(http.StatusCreated, "created", data))
}

func (handler *DiscountHandler) Update(ctx *gin.Context) {
	//get data id
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.DiscountRequestBody
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)
	input.UpdatedBy = user.ID

	data, err := handler.useCase.Update(id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}

func (handler *DiscountHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	//	find data to delete
	err := handler.useCase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, "not found", err.Error()))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", "ok"))

}

func (handler *DiscountHandler) FindAll(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	data := handler.useCase.FindAll(offset, limit)
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}

func (handler *DiscountHandler) FindByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := handler.useCase.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, "not found", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}
