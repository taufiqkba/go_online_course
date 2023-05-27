//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	repository3 "go_online_course/internal/cart/repository"
	usecase4 "go_online_course/internal/cart/usecase"
	"go_online_course/internal/class_room/repository"
	usecase2 "go_online_course/internal/class_room/usecase"
	repository4 "go_online_course/internal/discount/repository"
	usecase5 "go_online_course/internal/discount/usecase"
	repository2 "go_online_course/internal/order/repository"
	usecase3 "go_online_course/internal/order/usecase"
	repository5 "go_online_course/internal/order_detail/repository"
	usecase6 "go_online_course/internal/order_detail/usecase"
	usecase7 "go_online_course/internal/payment/usecase"
	repository6 "go_online_course/internal/product/repository"
	usecase8 "go_online_course/internal/product/usecase"
	"go_online_course/internal/web_hook/delivery/http"
	"go_online_course/internal/web_hook/usecase"
	"go_online_course/pkg/file_upload/cloudinary"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.WebHookHandler {
	wire.Build(
		repository2.NewOrderRepository,
		usecase3.NewOrderUseCase,
		usecase2.NewClassRoomUseCase,
		repository.NewClassRoomRepository,
		usecase.NewWebHookUseCase,
		http.NewWebHookHandler,
		repository3.NewCartRepository,
		usecase4.NewCartUseCase,
		repository4.NewDiscountRepository,
		usecase5.NewDiscountUseCase,
		repository5.NewOrderDetailRepository,
		usecase6.NewOrderDetailUseCase,
		usecase7.NewPaymentUseCase,
		repository6.NewProductRepository,
		usecase8.NewProductUseCase,
		cloudinary.NewFileUpload,
	)
	return &http.WebHookHandler{}
}
