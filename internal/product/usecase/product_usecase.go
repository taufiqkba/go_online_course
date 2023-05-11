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

func (p ProductUseCaseImpl) FindAll(offset int, limit int) []entity3.Product {
	return p.repository.FindAll(offset, limit)
}

func (p ProductUseCaseImpl) FindById(id int) (*entity3.Product, error) {
	return p.repository.FindById(id)
}

func (p ProductUseCaseImpl) Create(dto dto.ProductRequestBody) (*entity3.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUseCaseImpl) Update(id int, dto dto.ProductRequestBody) (*entity3.Product, error) {
	//	find product data by id
	product, err := p.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	product.Title = dto.Title
	product.Description = dto.Description
	product.Price = dto.Price

	//	if file image is available
	if dto.Image != nil {
		//	TODO IMPLEMENT ME
	}

	if dto.Video != nil {
		//	TODO IMPLEMENT ME
	}
}

func (p ProductUseCaseImpl) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewProductUseCase(repository repository.ProductRepository) ProductUseCase {
	return &ProductUseCaseImpl{repository}
}
