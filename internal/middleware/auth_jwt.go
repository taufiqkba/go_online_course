package middleware

import (
	"go_online_course/internal/oauth/dto"
	"go_online_course/pkg/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Header struct {
	Authorization string `header:"authorization" binding:"required"`
}

func AuthJwt(ctx *gin.Context) {
	var input Header

	if err := ctx.ShouldBindHeader(&input); err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.Response(http.StatusUnauthorized, "unauthorized", err.Error()))
		ctx.Abort()
		return
	}

	reqToken := input.Authorization
	splitToken := strings.Split(reqToken, "Bearer ")

	if len(splitToken) != 2 {
		ctx.JSON(http.StatusUnauthorized, utils.Response(http.StatusUnauthorized, "unauthorized", "unauthorized"))
		ctx.Abort()
		return
	}

	reqToken = splitToken[1]
	claims := &dto.MapClaimResponse{}

	token, err := jwt.ParseWithClaims(reqToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.Response(http.StatusUnauthorized, "unauthorized", err.Error()))
		ctx.Abort()
		return
	}

	if !token.Valid {
		ctx.JSON(http.StatusUnauthorized, utils.Response(http.StatusUnauthorized, "unauthorized", err.Error()))
		ctx.Abort()
		return
	}

	claims = token.Claims.(*dto.MapClaimResponse)
	ctx.Set("user", claims)
	ctx.Next()

}
