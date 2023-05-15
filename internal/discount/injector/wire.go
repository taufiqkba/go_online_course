//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"go_online_course/internal/discount/delivery/http"
	"go_online_course/internal/discount/repository"
	"go_online_course/internal/discount/usecase"
	"gorm.io/gorm"
)

func InitializeService(db *gorm.DB) *http.DiscountHandler {
	wire.Build(
		http.NewDiscountHandler,
		repository.NewDiscountRepository,
		usecase.NewDiscountUseCase,
	)
	return &http.DiscountHandler{}
}
