//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	repository7 "go_online_course/internal/admin/repository"
	usecase9 "go_online_course/internal/admin/usecase"
	repository3 "go_online_course/internal/cart/repository"
	usecase4 "go_online_course/internal/cart/usecase"
	"go_online_course/internal/dashboard/delivery/http"
	"go_online_course/internal/dashboard/usecase"
	repository4 "go_online_course/internal/discount/repository"
	usecase5 "go_online_course/internal/discount/usecase"
	repository2 "go_online_course/internal/order/repository"
	usecase3 "go_online_course/internal/order/usecase"
	repository5 "go_online_course/internal/order_detail/repository"
	usecase6 "go_online_course/internal/order_detail/usecase"
	usecase7 "go_online_course/internal/payment/usecase"
	repository6 "go_online_course/internal/product/repository"
	usecase8 "go_online_course/internal/product/usecase"
	repository8 "go_online_course/internal/user/repository"
	usecase10 "go_online_course/internal/user/usecase"
	"go_online_course/pkg/file_upload/cloudinary"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.DashboardHandler {
	wire.Build(
		//dashboard service
		http.NewDashboardHandler,
		usecase.NewDashboardUseCase,

		//order service
		repository2.NewOrderRepository,
		usecase3.NewOrderUseCase,

		//cart service
		repository3.NewCartRepository,
		usecase4.NewCartUseCase,

		//discount service
		repository4.NewDiscountRepository,
		usecase5.NewDiscountUseCase,

		//discount service
		repository5.NewOrderDetailRepository,
		usecase6.NewOrderDetailUseCase,

		//payment gateway service
		usecase7.NewPaymentUseCase,

		//product service
		repository6.NewProductRepository,
		usecase8.NewProductUseCase,

		//upload file to cloudinary service
		cloudinary.NewFileUpload,

		//admin service
		repository7.NewAdminRepository,
		usecase9.NewAdminUseCase,

		//user service
		repository8.NewUserRepositoryImpl,
		usecase10.NewUserUseCase,
	)
	return &http.DashboardHandler{}
}
