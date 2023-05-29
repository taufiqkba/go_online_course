package main

import (
	admin "go_online_course/internal/admin/injector"
	cart "go_online_course/internal/cart/injector"
	classRoom "go_online_course/internal/class_room/injector"
	dashboard "go_online_course/internal/dashboard/injector"
	discount "go_online_course/internal/discount/injector"
	oauth "go_online_course/internal/oauth/injector"
	order "go_online_course/internal/order/injector"
	product "go_online_course/internal/product/injector"
	productCategory "go_online_course/internal/product_category/injector"
	profile "go_online_course/internal/profile/injector"
	register "go_online_course/internal/register/injector"
	user "go_online_course/internal/user/injector"
	webhook "go_online_course/internal/web_hook/injector"
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
	product.InitializedServices(db).Route(&r.RouterGroup)
	cart.InitializedService(db).Route(&r.RouterGroup)
	discount.InitializeService(db).Route(&r.RouterGroup)
	order.InitializedService(db).Route(&r.RouterGroup)
	webhook.InitializedService(db).Route(&r.RouterGroup)
	classRoom.InitializedService(db).Route(&r.RouterGroup)
	dashboard.InitializedService(db).Route(&r.RouterGroup)
	user.InitializedService(db).Route(&r.RouterGroup)

	r.Run() //0.0.0.0:8080
}
