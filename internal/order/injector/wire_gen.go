// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	repository2 "go_online_course/internal/cart/repository"
	"go_online_course/internal/cart/usecase"
	repository3 "go_online_course/internal/discount/repository"
	usecase2 "go_online_course/internal/discount/usecase"
	"go_online_course/internal/order/delivery/http"
	"go_online_course/internal/order/repository"
	usecase6 "go_online_course/internal/order/usecase"
	repository5 "go_online_course/internal/order_detail/repository"
	usecase4 "go_online_course/internal/order_detail/usecase"
	usecase5 "go_online_course/internal/payment/usecase"
	repository4 "go_online_course/internal/product/repository"
	usecase3 "go_online_course/internal/product/usecase"
	"go_online_course/pkg/file_upload/cloudinary"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *http.OrderHandler {
	orderRepository := repository.NewOrderRepository(db)
	cartRepository := repository2.NewCartRepository(db)
	cartUseCase := usecase.NewCartUseCase(cartRepository)
	discountRepository := repository3.NewDiscountRepository(db)
	discountUseCase := usecase2.NewDiscountUseCase(discountRepository)
	productRepository := repository4.NewProductRepository(db)
	fileUpload := cloudinary.NewFileUpload()
	productUseCase := usecase3.NewProductUseCase(productRepository, fileUpload)
	orderDetailRepository := repository5.NewOrderDetailRepository(db)
	orderDetailUseCase := usecase4.NewOrderDetailUseCase(orderDetailRepository)
	paymentUseCase := usecase5.NewPaymentUseCase()
	orderUseCase := usecase6.NewOrderUseCase(orderRepository, cartUseCase, discountUseCase, productUseCase, orderDetailUseCase, paymentUseCase)
	orderHandler := http.NewOrderHandler(orderUseCase)
	return orderHandler
}
