//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"go_online_course/internal/cart/delivery/http"
	"go_online_course/internal/cart/repository"
	"go_online_course/internal/cart/usecase"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.CartHandler {
	wire.Build(
		http.NewCartHandler,
		repository.NewCartRepository,
		usecase.NewCartUseCase,
	)
	return &http.CartHandler{}
}
