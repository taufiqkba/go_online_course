package middleware

import (
	"github.com/gin-gonic/gin"
	"go_online_course/pkg/utils"
	"net/http"
)

func AuthAdmin(ctx *gin.Context) {
	admin := utils.GetCurrentUser(ctx)

	if !admin.IsAdmin {
		ctx.JSON(http.StatusUnauthorized, utils.Response(http.StatusUnauthorized, "unauthorized", "unauthorized"))
		ctx.Abort()
		return
	}
	//	TODO return middleware Auth Admin
}
