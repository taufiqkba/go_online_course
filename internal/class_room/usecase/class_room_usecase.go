package usecase

import (
	"go_online_course/internal/class_room/dto"
	entity4 "go_online_course/internal/class_room/entity"
	"go_online_course/internal/class_room/repository"
)

type ClassRoomUseCase interface {
	FindAllByUserID(offset int, limit int, userID int) dto.ClassRoomListResponse
	Create(dto dto.ClassRoom) (*entity4.ClassRoom, error)
}

type ClassRoomUseCaseImpl struct {
	repository repository.ClassRoomRepository
}

func (usecase *ClassRoomUseCaseImpl) FindAllByUserID(offset int, limit int, userID int) dto.ClassRoomListResponse {
	//if err := usecase.repository.FindAllByUserID(offset, limit, userID) != nil {
	//
	//}
	//TODO implement me
	panic("implement me")
}

func (usecase *ClassRoomUseCaseImpl) Create(dto dto.ClassRoom) (*entity4.ClassRoom, error) {
	//TODO implement me
	panic("implement me")
}

func NewClassRoomUseCase(repository repository.ClassRoomRepository) ClassRoomUseCase {
	return &ClassRoomUseCaseImpl{repository: repository}
}
