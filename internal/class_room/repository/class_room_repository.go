package repository

import (
	entity4 "go_online_course/internal/class_room/entity"
	"go_online_course/pkg/utils"
	"gorm.io/gorm"
)

type ClassRoomRepository interface {
	FindAllByUserID(offset int, limit int, userID int) []entity4.ClassRoom
	FindOneByUserIDAndProductID(userID int, productID int) (*entity4.ClassRoom, error)
	Create(entity entity4.ClassRoom) (*entity4.ClassRoom, error)
}

type ClassRoomRepositoryImpl struct {
	db *gorm.DB
}

func (repository *ClassRoomRepositoryImpl) FindAllByUserID(offset int, limit int, userID int) []entity4.ClassRoom {
	var classRoom []entity4.ClassRoom

	repository.db.Scopes(utils.Paginate(offset, limit)).Preload("Product.ProductCategory").Where("user_id = ?", userID).Find(&classRoom)

	return classRoom
}

func (repository *ClassRoomRepositoryImpl) FindOneByUserIDAndProductID(userID int, productID int) (*entity4.ClassRoom, error) {
	var classRoom entity4.ClassRoom

	if err := repository.db.Where("user_id = ?", userID).Where("product_id", productID).First(&classRoom).Error; err != nil {
		return nil, err
	}
	return &classRoom, nil
}

func (repository *ClassRoomRepositoryImpl) Create(entity entity4.ClassRoom) (*entity4.ClassRoom, error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func NewClassRoomRepository(db *gorm.DB) ClassRoomRepository {
	return &ClassRoomRepositoryImpl{db: db}
}
