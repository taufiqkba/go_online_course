package repository

import (
	"go_online_course/internal/product_category/entity"
	"go_online_course/pkg/utils"

	"gorm.io/gorm"
)

type ProductCategoryRepository interface {
	FindAll(offset int, limit int) []entity.ProductCategory
	FindByID(id int) (*entity.ProductCategory, error)
	Create(entity entity.ProductCategory) (*entity.ProductCategory, error)
	Update(entity entity.ProductCategory) (*entity.ProductCategory, error)
	Delete(entity entity.ProductCategory) error
}

type ProductCategoryRepositoryImpl struct {
	db *gorm.DB
}

// Create implements ProductCategoryRepository
func (repository *ProductCategoryRepositoryImpl) Create(entity entity.ProductCategory) (*entity.ProductCategory, error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements ProductCategoryRepository
func (repository *ProductCategoryRepositoryImpl) Delete(entity entity.ProductCategory) error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements ProductCategoryRepository
func (repository *ProductCategoryRepositoryImpl) FindAll(offset int, limit int) []entity.ProductCategory {
	var productCategories []entity.ProductCategory

	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&productCategories)
	return productCategories
}

// FindByID implements ProductCategoryRepository
func (repository *ProductCategoryRepositoryImpl) FindByID(id int) (*entity.ProductCategory, error) {
	var productCategory entity.ProductCategory

	if err := repository.db.First(&productCategory, id).Error; err != nil {
		return nil, err
	}
	return &productCategory, nil
}

// Update implements ProductCategoryRepository
func (repository *ProductCategoryRepositoryImpl) Update(entity entity.ProductCategory) (*entity.ProductCategory, error) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewProductCategoryRepository(db *gorm.DB) ProductCategoryRepository {
	return &ProductCategoryRepositoryImpl{db}
}
