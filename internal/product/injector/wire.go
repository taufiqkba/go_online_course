//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"go_online_course/internal/product/delivery/http"
	"go_online_course/internal/product/repository"
	"go_online_course/internal/product/usecase"
	"go_online_course/pkg/file_upload/cloudinary"
	"gorm.io/gorm"
)

func InitializedServices(db *gorm.DB) *http.ProductHandler {
	wire.Build(
		http.NewProductHandler,
		usecase.NewProductUseCase,
		repository.NewProductRepository,
		cloudinary.NewFileUpload,
	)
	return &http.ProductHandler{}
}
