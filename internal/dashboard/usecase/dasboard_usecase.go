package usecase

import (
	usecase2 "go_online_course/internal/admin/usecase"
	"go_online_course/internal/dashboard/dto"
	usecase4 "go_online_course/internal/order/usecase"
	usecase3 "go_online_course/internal/product/usecase"
	"go_online_course/internal/user/usecase"
)

type DashboardUseCase interface {
	GetDataDashboard() dto.DashboardResponseBody
}
type DashboardUseCaseImpl struct {
	userUseCase    usecase.UserUseCase
	adminUseCase   usecase2.AdminUseCase
	productUseCase usecase3.ProductUseCase
	orderUseCase   usecase4.OrderUseCase
}

func (useCase *DashboardUseCaseImpl) GetDataDashboard() dto.DashboardResponseBody {
	//TODO count total admin
	totalUser := useCase.userUseCase.Count()
	totalAdmin := useCase.adminUseCase.Count()
	totalOrder := useCase.orderUseCase.Count()
	totalProduct := useCase.productUseCase.Count()

	return dto.DashboardResponseBody{
		TotalUser:    int64(totalUser),
		TotalProduct: int64(totalProduct),
		TotalOrder:   int64(totalOrder),
		TotalAdmin:   int64(totalAdmin),
	}
}

func NewDashboardUseCase(
	userUseCase usecase.UserUseCase,
	adminUseCase usecase2.AdminUseCase,
	productUseCase usecase3.ProductUseCase,
	orderUseCase usecase4.OrderUseCase,
) DashboardUseCase {
	return &DashboardUseCaseImpl{
		userUseCase:    userUseCase,
		adminUseCase:   adminUseCase,
		productUseCase: productUseCase,
		orderUseCase:   orderUseCase,
	}
}
