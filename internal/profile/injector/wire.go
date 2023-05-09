//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"go_online_course/internal/profile/delivery/http"
	"go_online_course/internal/profile/usecase"
	"go_online_course/internal/user/repository"
	usecase2 "go_online_course/internal/user/usecase"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *http.ProfileHandler {
	wire.Build(
		http.NewProfileHandler,
		usecase.NewProfileUseCase,
		repository.NewUserRepositoryImpl,
		usecase2.NewUserUseCase,
	)

	return &http.ProfileHandler{}
}
