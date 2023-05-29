//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"go_online_course/internal/user/delivery/http"
	"go_online_course/internal/user/repository"
	"go_online_course/internal/user/usecase"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.UserHandler {
	wire.Build(
		http.NewUserHandler,
		repository.NewUserRepositoryImpl,
		usecase.NewUserUseCase,
	)
	return &http.UserHandler{}
}
