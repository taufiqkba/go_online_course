package usecase

import (
	"go_online_course/internal/product/dto"
	entity3 "go_online_course/internal/product/entity"
	"go_online_course/internal/product/repository"
)

type ProductUseCase interface {
	FindAll(offset int, limit int) []entity3.Product
	FindById(id int) (*entity3.Product, error)
	Create(dto dto.ProductRequestBody) (*entity3.Product, error)
	Update(id int, dto dto.ProductRequestBody) (*entity3.Product, error)
	Delete(id int) error
}

type ProductUseCaseImpl struct {
	repository repository.ProductRepository
}

func (usecase *ProductUseCaseImpl) FindAll(offset int, limit int) []entity3.Product {
	return usecase.repository.FindAll(offset, limit)
}

func (usecase *ProductUseCaseImpl) FindById(id int) (*entity3.Product, error) {
	return usecase.repository.FindById(id)
}

func (usecase *ProductUseCaseImpl) Create(dto dto.ProductRequestBody) (*entity3.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (usecase *ProductUseCaseImpl) Update(id int, dto dto.ProductRequestBody) (*entity3.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (usecase *ProductUseCaseImpl) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewProductUseCase(repository repository.ProductRepository) ProductUseCase {
	return &ProductUseCaseImpl{repository: repository}
}
