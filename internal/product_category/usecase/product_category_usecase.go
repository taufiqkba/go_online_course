package usecase

import (
	"go_online_course/internal/product_category/dto"
	"go_online_course/internal/product_category/entity"
	"go_online_course/internal/product_category/repository"
)

type ProductCategoryUseCase interface {
	FindAll(offset int, limit int) []entity.ProductCategory
	FindByID(id int) (*entity.ProductCategory, error)
	Create(dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, error)
	Update(id int, dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, error)
	Delete(id int) error
}

type ProductCategoryUseCaseImpl struct {
	repository repository.ProductCategoryRepository
}

// Create implements ProductCategoryUseCase
func (usecase *ProductCategoryUseCaseImpl) Create(dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, error) {
	productCategoryEntity := entity.ProductCategory{
		Name:        dto.Name,
		CreatedByID: dto.CreatedBy,
	}

	productCategory, err := usecase.repository.Create(productCategoryEntity)
	if err != nil {
		return nil, err
	}
	return productCategory, nil
}

// Delete implements ProductCategoryUseCase
func (usecase *ProductCategoryUseCaseImpl) Delete(id int) error {
	productCategory, err := usecase.repository.FindByID(id)

	if err != nil {
		return err
	}

	if err := usecase.repository.Delete(*productCategory); err != nil {
		return err
	}
	return nil
}

// FindAll implements ProductCategoryUseCase
func (usecase *ProductCategoryUseCaseImpl) FindAll(offset int, limit int) []entity.ProductCategory {
	return usecase.repository.FindAll(offset, limit)
}

// FindByID implements ProductCategoryUseCase
func (usecase *ProductCategoryUseCaseImpl) FindByID(id int) (*entity.ProductCategory, error) {
	return usecase.repository.FindByID(id)
}

// Update implements ProductCategoryUseCase
func (usecase *ProductCategoryUseCaseImpl) Update(id int, dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, error) {
	//Find product_category by id
	productCategory, err := usecase.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	productCategory.Name = dto.Name
	productCategory.UpdatedByID = &dto.UpdatedBy

	//	call repository to update data product category
	updateProductCategory, err := usecase.repository.Update(*productCategory)
	if err != nil {
		return nil, err
	}
	return updateProductCategory, nil
}

func NewProductCategoryUseCase(repository repository.ProductCategoryRepository) ProductCategoryUseCase {
	return &ProductCategoryUseCaseImpl{repository}
}
