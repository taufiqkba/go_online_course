package http

import (
	"go_online_course/internal/admin/dto"
	"go_online_course/internal/admin/usecase"
	"go_online_course/internal/middleware"
	"go_online_course/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	usecase usecase.AdminUseCase
}

func NewAdminHandler(usecase usecase.AdminUseCase) *AdminHandler {
	return &AdminHandler{usecase}
}

func (handler *AdminHandler) Route(r *gin.RouterGroup) {
	adminRouter := r.Group("/api/v1")

	adminRouter.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		adminRouter.GET("/admins", handler.FindAll)
		adminRouter.GET("/admins/:id", handler.FindByID)
		adminRouter.POST("/admins", handler.Create)
		adminRouter.PATCH("/admins/:id", handler.Update)
		adminRouter.DELETE("/admins/:id", handler.Delete)
	}

}

func (handler *AdminHandler) Create(ctx *gin.Context) {
	var input dto.AdminRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.CreatedBy = user.ID

	// Create data
	_, err := handler.usecase.Create(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, utils.Response(http.StatusCreated, "created", "created"))
}
func (handler *AdminHandler) Update(ctx *gin.Context) {
	// api/v1/admin/:id

	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.AdminRequestBody
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)
	input.UpdatedBy = user.ID

	//	Update data with call repository
	data, err := handler.usecase.Update(id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}

func (handler *AdminHandler) FindAll(ctx *gin.Context) {
	// api/v1/admins?offset=1&limit=1
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	data := handler.usecase.FindAll(offset, limit)
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}

func (handler *AdminHandler) FindByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := handler.usecase.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, "not found", "not found"))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}

func (handler *AdminHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := handler.usecase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, "not found", "not found"))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", "ok"))
}
