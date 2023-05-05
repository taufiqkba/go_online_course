package repository

import (
	"go_online_course/internal/user/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(offset int, limit int) []entity.User
	FindById(id int) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Create(entity entity.User) (*entity.User, error)
	Update(entity entity.User) (*entity.User, error)
	Delete(entity entity.User) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (ur UserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (ur UserRepositoryImpl) FindAll(offset int, limit int) []entity.User {
	var users []entity.User
	ur.db.Find(&users)
	return users
}

func (ur UserRepositoryImpl) FindById(id int) (*entity.User, error) {
	var users entity.User
	if err := ur.db.First(&users, id).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (ur UserRepositoryImpl) Create(entity entity.User) (*entity.User, error) {
	if err := ur.db.Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (ur UserRepositoryImpl) Update(entity entity.User) (*entity.User, error) {
	if err := ur.db.Save(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (ur UserRepositoryImpl) Delete(entity entity.User) error {
	if err := ur.db.Delete(&entity).Error; err != nil {
		return err
	}
	return nil
}
