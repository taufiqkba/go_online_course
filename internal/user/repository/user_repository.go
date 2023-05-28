package repository

import (
	"go_online_course/internal/user/entity"
	"go_online_course/pkg/utils"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(offset int, limit int) []entity.User
	FindById(id int) (*entity.User, error)
	Count() int
	FindByEmail(email string) (*entity.User, error)
	Create(entity entity.User) (*entity.User, error)
	Update(entity entity.User) (*entity.User, error)
	Delete(entity entity.User) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (repository UserRepositoryImpl) Count() int {
	var user entity.User

	var totalUser int64
	repository.db.Model(&user).Count(&totalUser)
	return int(totalUser)
}

func (repository UserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := repository.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (repository UserRepositoryImpl) FindAll(offset int, limit int) []entity.User {
	var users []entity.User

	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&users)

	return users
}

func (repository UserRepositoryImpl) FindById(id int) (*entity.User, error) {
	var users entity.User
	if err := repository.db.First(&users, id).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (repository UserRepositoryImpl) Create(entity entity.User) (*entity.User, error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repository UserRepositoryImpl) Update(entity entity.User) (*entity.User, error) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repository UserRepositoryImpl) Delete(entity entity.User) error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return err
	}
	return nil
}
