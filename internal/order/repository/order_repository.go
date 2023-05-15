package repository

import (
	"go_online_course/internal/order/entity"
	"go_online_course/pkg/utils"
	"gorm.io/gorm"
)

type OrderRepository interface {
	FindAll(offset int, limit int) []entity.Order
	FindByID(id int) (*entity.Order, error)
	Create(entity entity.Order) (*entity.Order, error)
}

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func (repository *OrderRepositoryImpl) FindAll(offset int, limit int) []entity.Order {
	var orders []entity.Order

	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&orders)
	return orders
}

func (repository *OrderRepositoryImpl) FindByID(id int) (*entity.Order, error) {
	var order entity.Order
	err := repository.db.First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (repository *OrderRepositoryImpl) Create(entity entity.Order) (*entity.Order, error) {
	err := repository.db.Create(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{db: db}
}
