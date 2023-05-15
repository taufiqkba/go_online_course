package usecase

import (
	"fmt"
	"go_online_course/internal/order_detail/entity"
	"go_online_course/internal/order_detail/repository"
)

type OrderDetailUseCase interface {
	Create(entity entity.OrderDetail) (*entity.OrderDetail, error)
}

type OrderDetailUseCaseImpl struct {
	repository repository.OrderDetailRepository
}

func (useCase *OrderDetailUseCaseImpl) Create(entity entity.OrderDetail) (*entity.OrderDetail, error) {
	data, err := useCase.repository.Create(entity)
	if err != nil {
		fmt.Println(err)
	}
	return data, nil
}

func NewOrderDetailUseCase(repository repository.OrderDetailRepository) OrderDetailUseCase {
	return &OrderDetailUseCaseImpl{repository: repository}
}
