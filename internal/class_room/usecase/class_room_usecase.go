package usecase

import (
	"errors"
	"go_online_course/internal/class_room/dto"
	entity4 "go_online_course/internal/class_room/entity"
	"go_online_course/internal/class_room/repository"
	"gorm.io/gorm"
)

type ClassRoomUseCase interface {
	FindAllByUserID(offset int, limit int, userID int) dto.ClassRoomListResponse
	Create(dto dto.ClassRoom) (*entity4.ClassRoom, error)
}

type ClassRoomUseCaseImpl struct {
	repository repository.ClassRoomRepository
}

// Create implements ClassRoomUseCase
func (useCase *ClassRoomUseCaseImpl) Create(dto dto.ClassRoom) (*entity4.ClassRoom, error) {
	//	to validation user_id and product_id is available
	dataClassRoom, err := useCase.repository.FindByOneByUserIDAndProductID(int(dto.UserID), int(dto.ProductID))

	//if data product is empty, skip this condition
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if dataClassRoom != nil {
		return nil, errors.New("you are already purchased this class")
	}

	classRoom := entity4.ClassRoom{
		ID:        dto.UserID,
		ProductID: dto.ProductID,
	}
	data, err := useCase.repository.Create(classRoom)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// FindAllByUserID implements ClassRoomUseCase
func (useCase *ClassRoomUseCaseImpl) FindAllByUserID(offset int, limit int, userID int) dto.ClassRoomListResponse {
	classRooms := useCase.repository.FindAllByUserID(offset, limit, userID)
	classRoomsResp := dto.CreateClassRoomListResponse(classRooms)
	return classRoomsResp
}

func NewClassRoomUseCase(repository repository.ClassRoomRepository) ClassRoomUseCase {
	return &ClassRoomUseCaseImpl{repository: repository}
}
