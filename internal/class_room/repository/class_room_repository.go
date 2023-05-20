package repository

import (
	entity4 "go_online_course/internal/class_room/entity"
	"go_online_course/pkg/utils"

	"gorm.io/gorm"
)

type ClassRoomRepository interface {
	FindAllByUserID(offset int, limit int, userID int) []entity4.ClassRoom
	FindByOneByUserIDAndProductID(userID int, productID int) (*entity4.ClassRoom, error)
	Create(entity entity4.ClassRoom) (*entity4.ClassRoom, error)
}

type ClassRoomRepositoryImpl struct {
	db *gorm.DB
}

// Create implements ClassRoomRepository
func (repository *ClassRoomRepositoryImpl) Create(entity entity4.ClassRoom) (*entity4.ClassRoom, error) {
	err := repository.db.Create(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

// FindAllByUserID implements ClassRoomRepository
func (repository *ClassRoomRepositoryImpl) FindAllByUserID(offset int, limit int, userID int) []entity4.ClassRoom {
	var classRoom []entity4.ClassRoom

	repository.db.Scopes(utils.Paginate(offset, limit)).Preload("Product.ProductCategory").Where("user_id = ?", userID).Find(&classRoom)
	return classRoom
}

// FindByOneByUserIDAndProductID implements ClassRoomRepository
func (repository *ClassRoomRepositoryImpl) FindByOneByUserIDAndProductID(userID int, productID int) (*entity4.ClassRoom, error) {
	var classRoom entity4.ClassRoom

	err := repository.db.Where("user_id = ?", userID).Where("product_id", productID).First(&classRoom).Error
	if err != nil {
		return nil, err
	}
	return &classRoom, nil
}

func NewClassRoomRepository(db *gorm.DB) ClassRoomRepository {
	return &ClassRoomRepositoryImpl{db: db}
}
