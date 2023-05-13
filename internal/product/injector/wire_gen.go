// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"go_online_course/internal/product/delivery/http"
	"go_online_course/internal/product/repository"
	"go_online_course/internal/product/usecase"
	"go_online_course/pkg/file_upload/cloudinary"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializedServices(db *gorm.DB) *http.ProductHandler {
	productRepository := repository.NewProductRepository(db)
	fileUpload := cloudinary.NewFileUpload()
	productUseCase := usecase.NewProductUseCase(productRepository, fileUpload)
	productHandler := http.NewProductHandler(productUseCase)
	return productHandler
}