package http

import (
	"github.com/gin-gonic/gin"
	"go_online_course/internal/class_room/usecase"
	"go_online_course/internal/middleware"
	"go_online_course/pkg/utils"
	"net/http"
	"strconv"
)

type ClassRoomHandler struct {
	useCase usecase.ClassRoomUseCase
}

func NewClassRoomHandler(useCase usecase.ClassRoomUseCase) *ClassRoomHandler {
	return &ClassRoomHandler{useCase: useCase}
}

func (handler *ClassRoomHandler) Route(r *gin.RouterGroup) {
	classRoomHandler := r.Group("/api/v1")

	classRoomHandler.Use(middleware.AuthJwt)
	{
		classRoomHandler.GET("/class_rooms", handler.FindAllByUseriD)
	}
}

func (handler *ClassRoomHandler) FindAllByUseriD(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	user := utils.GetCurrentUser(ctx)

	data := handler.useCase.FindAllByUserID(offset, limit, int(user.ID))
	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}
