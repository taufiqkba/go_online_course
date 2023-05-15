package repository

import (
	"go_online_course/internal/order_detail/entity"
	"gorm.io/gorm"
)

type OrderDetailRepository interface {
	Create(entity entity.OrderDetail) (*entity.OrderDetail, error)
}

type OrderDetailRepositoryImpl struct {
	db *gorm.DB
}

func (repository *OrderDetailRepositoryImpl) Create(entity entity.OrderDetail) (*entity.OrderDetail, error) {
	err := repository.db.Create(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewOrderDetailRepository(db *gorm.DB) OrderDetailRepository {
	return &OrderDetailRepositoryImpl{db: db}
}
