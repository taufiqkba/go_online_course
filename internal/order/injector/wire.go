//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"go_online_course/internal/cart/repository"
	"go_online_course/internal/cart/usecase"
	repository2 "go_online_course/internal/discount/repository"
	usecase2 "go_online_course/internal/discount/usecase"
	"go_online_course/internal/order/delivery/http"
	repository3 "go_online_course/internal/order/repository"
	usecase3 "go_online_course/internal/order/usecase"
	repository4 "go_online_course/internal/order_detail/repository"
	usecase4 "go_online_course/internal/order_detail/usecase"
	usecase5 "go_online_course/internal/payment/usecase"
	repository5 "go_online_course/internal/product/repository"
	usecase6 "go_online_course/internal/product/usecase"
	"go_online_course/pkg/file_upload/cloudinary"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.OrderHandler {
	wire.Build(
		repository.NewCartRepository,
		usecase.NewCartUseCase,
		repository2.NewDiscountRepository,
		usecase2.NewDiscountUseCase,
		http.NewOrderHandler,
		repository3.NewOrderRepository,
		usecase3.NewOrderUseCase,
		repository4.NewOrderDetailRepository,
		usecase4.NewOrderDetailUseCase,
		usecase5.NewPaymentUseCase,
		repository5.NewProductRepository,
		usecase6.NewProductUseCase,
		cloudinary.NewFileUpload,
	)
	return &http.OrderHandler{}
}
