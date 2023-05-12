package usecase

import (
	"errors"
	"go_online_course/internal/cart/dto"
	"go_online_course/internal/cart/entity"
	"go_online_course/internal/cart/repository"
)

type CartUseCase interface {
	FindByUserID(userID int, offset int, limit int) []entity.Cart
	FindByID(id int) (*entity.Cart, error)
	Create(dto dto.CartRequestBody) (*entity.Cart, error)
	Delete(id int, userID int) error
}

type CartUseCaseImpl struct {
	repository repository.CartRepository
}

func (usecase *CartUseCaseImpl) FindByUserID(userID int, offset int, limit int) []entity.Cart {
	return usecase.repository.FindByUserID(userID, offset, limit)
}

func (usecase *CartUseCaseImpl) FindByID(id int) (*entity.Cart, error) {
	return usecase.repository.FindByID(id)
}

func (usecase *CartUseCaseImpl) Create(dto dto.CartRequestBody) (*entity.Cart, error) {
	cart := entity.Cart{
		UserID:    dto.UserID,
		ProductID: dto.ProductID,
	}

	//	TODO Validation to check whether the user has ever entered data with the same product ID or not

	//	Insert Data
	data, err := usecase.repository.Create(cart)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (usecase *CartUseCaseImpl) Delete(id int, userID int) error {
	//	Search based form CartID
	cart, err := usecase.repository.FindByID(id)
	if err != nil {
		return err
	}

	if cart.User.ID != int64(userID) {
		return errors.New("this cart not yours")
	}
	err = usecase.repository.Delete(*cart)
	if err != nil {
		return err
	}
	return nil

}

func NewCartUseCase(repository repository.CartRepository) CartUseCase {
	return &CartUseCaseImpl{repository: repository}
}
