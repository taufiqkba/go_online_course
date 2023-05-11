package repository

import (
	entity3 "go_online_course/internal/product/entity"
	"go_online_course/pkg/utils"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(offset int, limit int) []entity3.Product
	FindById(id int) (*entity3.Product, error)
	Create(entity entity3.Product) (*entity3.Product, error)
	Update(entity entity3.Product) (*entity3.Product, error)
	Delete(entity entity3.Product) error
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func (repository *ProductRepositoryImpl) FindAll(offset int, limit int) []entity3.Product {
	var products []entity3.Product
	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&products)
	return products
}

func (repository *ProductRepositoryImpl) FindById(id int) (*entity3.Product, error) {
	var product entity3.Product
	err := repository.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (repository *ProductRepositoryImpl) Create(entity entity3.Product) (*entity3.Product, error) {
	err := repository.db.Create(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repository *ProductRepositoryImpl) Update(entity entity3.Product) (*entity3.Product, error) {
	err := repository.db.Save(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repository *ProductRepositoryImpl) Delete(entity entity3.Product) error {
	err := repository.db.Delete(entity).Error
	if err != nil {
		return err
	}
	return nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}
