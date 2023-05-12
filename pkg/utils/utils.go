package utils

import (
	"github.com/gin-gonic/gin"
	"go_online_course/internal/oauth/dto"
	"gorm.io/gorm"
	"math/rand"
	"path/filepath"
)

func RandomString(number int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, number)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Paginate(offset int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := offset

		//if value on page less than or equal 0
		if page <= 0 {
			page = 1
		}
		pageSize := limit

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(pageSize)
	}
}

func GetCurrentUser(ctx *gin.Context) *dto.MapClaimResponse {
	user, _ := ctx.Get("user")
	return user.(*dto.MapClaimResponse)
}

func GetFileName(fileName string) string {
	file := filepath.Base(fileName)

	return file[:len(file)-len(filepath.Ext(file))]
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
