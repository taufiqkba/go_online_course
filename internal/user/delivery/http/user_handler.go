package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course/internal/middleware"
	"go_online_course/internal/user/dto"
	"go_online_course/internal/user/usecase"
	"go_online_course/pkg/utils"
	"net/http"
	"strconv"
)

type UserHandler struct {
	usecase usecase.UserUseCase
}

func NewUserHandler(usecase usecase.UserUseCase) *UserHandler {
	return &UserHandler{usecase}
}

func (handler *UserHandler) Route(r *gin.RouterGroup) {
	userHandler := r.Group("/api/v1")

	userHandler.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		userHandler.GET("/users", handler.FindAll)
		userHandler.GET("/users/:id", handler.FindByID)
		userHandler.POST("/users", handler.Create)
		userHandler.PATCH("/users/:id", handler.Update)
		userHandler.DELETE("/users/:id", handler.Delete)
	}
}

func (handler *UserHandler) FindAll(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	data := handler.usecase.FindAll(offset, limit)
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}

func (handler *UserHandler) FindByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := handler.usecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, "not found", err.Error()))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}

func (handler *UserHandler) Create(ctx *gin.Context) {
	var input dto.UserRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}
	user := utils.GetCurrentUser(ctx)
	input.CreatedBy = &user.ID

	data, err := handler.usecase.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, utils.Response(http.StatusCreated, "created", data))
}

func (handler *UserHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.UserRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}
	user := utils.GetCurrentUser(ctx)
	input.UpdatedBy = &user.ID

	data, err := handler.usecase.Update(id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}

func (handler *UserHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := handler.usecase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, "not found", err.Error()))
	}
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", "ok"))

}
