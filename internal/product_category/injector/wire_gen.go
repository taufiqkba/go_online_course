// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"go_online_course/internal/product_category/delivery/http"
	"go_online_course/internal/product_category/repository"
	"go_online_course/internal/product_category/usecase"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *http.ProductCategoryHandler {
	productCategoryRepository := repository.NewProductCategoryRepository(db)
	productCategoryUseCase := usecase.NewProductCategoryUseCase(productCategoryRepository)
	productCategoryHandler := http.NewProductCategoryHandler(productCategoryUseCase)
	return productCategoryHandler
}
