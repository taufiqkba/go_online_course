//go:build wireinject
// +build wireinject

package injector

import (
	"go_online_course/internal/product_category/delivery/http"
	"go_online_course/internal/product_category/repository"
	"go_online_course/internal/product_category/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.ProductCategoryHandler {
	wire.Build(
		http.NewProductCategoryHandler,
		repository.NewProductCategoryRepository,
		usecase.NewProductCategoryUseCase,
	)
	return &http.ProductCategoryHandler{}
}
