package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course/internal/product/dto"
	"go_online_course/internal/product/usecase"
	"go_online_course/pkg/utils"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	usecase usecase.ProductUseCase
}

func NewProductHandler(usecase usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{usecase}
}

func (handler *ProductHandler) FindAll(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	data := handler.usecase.FindAll(offset, limit)

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}

func (handler *ProductHandler) FindByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := handler.usecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, "not found", "not found"))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}

func (handler *ProductHandler) Create(ctx *gin.Context) {
	var input dto.ProductRequestBody

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", "bad request"))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)
	input.CreatedBy = user.ID
	data, err := handler.usecase.Create(input)
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", "internal server error"))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, utils.Response(http.StatusCreated, "created", data))
}

func (handler *ProductHandler) Update(ctx *gin.Context) {
	//id, _ := strconv.Atoi(ctx.Param("id"))
	//	TODO IMPLEMENT Update Handler
}

func (handler *ProductHandler) Delete(ctx *gin.Context) {
	//id, _ := strconv.Atoi(ctx.Param("id"))
	//	TODO IMPLEMENT Delete Handler
}
