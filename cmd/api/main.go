package main

import (
	admin "go_online_course/internal/admin/injector"
	oauth "go_online_course/internal/oauth/injector"
	productCategory "go_online_course/internal/product_category/injector"
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
	admin.InitializedService(db).Route(&r.RouterGroup)
	productCategory.InitializedService(db).Route(&r.RouterGroup)

	r.Run() //0.0.0.0:8080
}
