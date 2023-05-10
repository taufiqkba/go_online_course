package middleware

import (
	"go_online_course/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthAdmin(ctx *gin.Context) {
	admin := utils.GetCurrentUser(ctx)

	if !admin.IsAdmin {
		ctx.JSON(http.StatusUnauthorized, utils.Response(http.StatusUnauthorized, "unauthorized", "unauthorized"))
		ctx.Abort()
		return
	}
	ctx.Next()
}
