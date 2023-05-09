// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"go_online_course/internal/profile/delivery/http"
	usecase2 "go_online_course/internal/profile/usecase"
	"go_online_course/internal/user/repository"
	"go_online_course/internal/user/usecase"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *http.ProfileHandler {
	userRepository := repository.NewUserRepositoryImpl(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	profileUseCase := usecase2.NewProfileUseCase(userUseCase)
	profileHandler := http.NewProfileHandler(profileUseCase)
	return profileHandler
}
