// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"go_online_course/internal/class_room/delivery/http"
	"go_online_course/internal/class_room/repository"
	"go_online_course/internal/class_room/usecase"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *http.ClassRoomHandler {
	classRoomRepository := repository.NewClassRoomRepository(db)
	classRoomUseCase := usecase.NewClassRoomUseCase(classRoomRepository)
	classRoomHandler := http.NewClassRoomHandler(classRoomUseCase)
	return classRoomHandler
}