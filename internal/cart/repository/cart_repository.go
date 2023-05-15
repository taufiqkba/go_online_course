package repository

import (
	"go_online_course/internal/cart/entity"
	"go_online_course/pkg/utils"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindByUserID(userID int, offset int, limit int) []entity.Cart
	FindByID(id int) (*entity.Cart, error)
	Create(entity entity.Cart) (*entity.Cart, error)
	Delete(entity entity.Cart) error
	DeleteByUserID(userID int) error
}

type CartRepositoryImpl struct {
	db *gorm.DB
}

func (repository *CartRepositoryImpl) DeleteByUserID(userID int) error {
	var cart entity.Cart
	err := repository.db.Where("user_id = ?", userID).Delete(&cart).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *CartRepositoryImpl) FindByUserID(userID int, offset int, limit int) []entity.Cart {
	var carts []entity.Cart

	repository.db.Scopes(utils.Paginate(offset, limit)).Preload("User").Preload("Product").Where("user_id = ?", userID).Find(&carts)
	return carts
}

func (repository *CartRepositoryImpl) FindByID(id int) (*entity.Cart, error) {
	var cart entity.Cart
	if err := repository.db.Preload("User").Preload("Product").Find(&cart, id).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (repository *CartRepositoryImpl) Create(entity entity.Cart) (*entity.Cart, error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repository *CartRepositoryImpl) Delete(entity entity.Cart) error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return err
	}
	return nil
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &CartRepositoryImpl{db: db}
}
