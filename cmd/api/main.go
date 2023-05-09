package main

import (
	oauth "go_online_course/internal/oauth/injector"
	profile "go_online_course/internal/profile/injector"
	register "go_online_course/internal/register/injector"
	"go_online_course/pkg/db/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	db := mysql.DB()
	r := gin.Default()

	register.InitializedService(db).Route(&r.RouterGroup)
	oauth.InitializedService(db).Route(&r.RouterGroup)
	profile.InitializedService(db).Route(&r.RouterGroup)

	r.Run() //0.0.0.0:8080
}
