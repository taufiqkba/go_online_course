package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course/internal/dashboard/usecase"
	"go_online_course/internal/middleware"
	"go_online_course/pkg/utils"
	"net/http"
)

type DashboardHandler struct {
	dashboardUseCase usecase.DashboardUseCase
}

func NewDashboardHandler(dashboardUseCase usecase.DashboardUseCase) *DashboardHandler {
	return &DashboardHandler{dashboardUseCase: dashboardUseCase}
}

func (handler *DashboardHandler) Route(r *gin.RouterGroup) {
	dashboardHandler := r.Group("/api/v1")

	dashboardHandler.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		dashboardHandler.GET("/dashboard", handler.GetDataDashboard)
	}
}

func (handler *DashboardHandler) GetDataDashboard(ctx *gin.Context) {
	data := handler.dashboardUseCase.GetDataDashboard()

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}
