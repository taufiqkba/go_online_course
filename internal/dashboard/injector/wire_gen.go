// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	repository2 "go_online_course/internal/admin/repository"
	usecase2 "go_online_course/internal/admin/usecase"
	repository5 "go_online_course/internal/cart/repository"
	usecase4 "go_online_course/internal/cart/usecase"
	"go_online_course/internal/dashboard/delivery/http"
	usecase9 "go_online_course/internal/dashboard/usecase"
	repository6 "go_online_course/internal/discount/repository"
	usecase5 "go_online_course/internal/discount/usecase"
	repository4 "go_online_course/internal/order/repository"
	usecase8 "go_online_course/internal/order/usecase"
	repository7 "go_online_course/internal/order_detail/repository"
	usecase6 "go_online_course/internal/order_detail/usecase"
	usecase7 "go_online_course/internal/payment/usecase"
	repository3 "go_online_course/internal/product/repository"
	usecase3 "go_online_course/internal/product/usecase"
	"go_online_course/internal/user/repository"
	"go_online_course/internal/user/usecase"
	"go_online_course/pkg/file_upload/cloudinary"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *http.DashboardHandler {
	userRepository := repository.NewUserRepositoryImpl(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	adminRepository := repository2.NewAdminRepository(db)
	adminUseCase := usecase2.NewAdminUseCase(adminRepository)
	productRepository := repository3.NewProductRepository(db)
	fileUpload := cloudinary.NewFileUpload()
	productUseCase := usecase3.NewProductUseCase(productRepository, fileUpload)
	orderRepository := repository4.NewOrderRepository(db)
	cartRepository := repository5.NewCartRepository(db)
	cartUseCase := usecase4.NewCartUseCase(cartRepository)
	discountRepository := repository6.NewDiscountRepository(db)
	discountUseCase := usecase5.NewDiscountUseCase(discountRepository)
	orderDetailRepository := repository7.NewOrderDetailRepository(db)
	orderDetailUseCase := usecase6.NewOrderDetailUseCase(orderDetailRepository)
	paymentUseCase := usecase7.NewPaymentUseCase()
	orderUseCase := usecase8.NewOrderUseCase(orderRepository, cartUseCase, discountUseCase, productUseCase, orderDetailUseCase, paymentUseCase)
	dashboardUseCase := usecase9.NewDashboardUseCase(userUseCase, adminUseCase, productUseCase, orderUseCase)
	dashboardHandler := http.NewDashboardHandler(dashboardUseCase)
	return dashboardHandler
}
