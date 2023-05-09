package utils

import (
	"go_online_course/internal/oauth/dto"
	"math/rand"

	"github.com/gin-gonic/gin"
)

func RandomString(number int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, number)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetCurrentUser(ctx *gin.Context) *dto.MapClaimResponse {
	user, _ := ctx.Get("user")
	return user.(*dto.MapClaimResponse)
}
