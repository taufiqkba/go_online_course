package main

import (
	"github.com/gin-gonic/gin"
	"go_online_course/internal/register/delivery/http"
	usecase2 "go_online_course/internal/register/usecase"
	"go_online_course/internal/user/repository"
	"go_online_course/internal/user/usecase"
	"go_online_course/pkg/db/mysql"
	"go_online_course/pkg/mail/sendgrid"
)

func main() {
	db := mysql.DB()
	r := gin.Default()

	mail := sendgrid.NewMail()
	userRepository := repository.NewUserRepositoryImpl(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	registerUseCase := usecase2.NewRegisterUseCase(userUseCase, mail)
	http.NewRegisterHandler(registerUseCase).Route(&r.RouterGroup)

	r.Run() //0.0.0.0:8080
}
