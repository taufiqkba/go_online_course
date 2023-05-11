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

// Create implements ProductRepository
func (repository *ProductRepositoryImpl) Create(entity entity3.Product) (*entity3.Product, error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements ProductRepository
func (repository *ProductRepositoryImpl) Delete(entity entity3.Product) error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements ProductRepository
func (repository *ProductRepositoryImpl) FindAll(offset int, limit int) []entity3.Product {
	var product []entity3.Product

	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&product)
	return product
}

// FindById implements ProductRepository
func (repository *ProductRepositoryImpl) FindById(id int) (*entity3.Product, error) {
	var product entity3.Product

	if err := repository.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// Update implements ProductRepository
func (repository *ProductRepositoryImpl) Update(entity entity3.Product) (*entity3.Product, error) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db}
}
