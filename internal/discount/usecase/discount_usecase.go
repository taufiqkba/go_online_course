package usecase

import (
	"database/sql"
	"errors"
	"go_online_course/internal/discount/dto"
	"go_online_course/internal/discount/entity"
	"go_online_course/internal/discount/repository"
)

type DiscountUseCase interface {
	FindAll(offset int, limit int) []entity.Discount
	FindByID(id int) (*entity.Discount, error)
	FindByCode(code string) (*entity.Discount, error)
	Create(dto dto.DiscountRequestBody) (*entity.Discount, error)
	Update(id int, dto dto.DiscountRequestBody) (*entity.Discount, error)
	Delete(id int) error
	UpdateRemainingQuantity(id int, quantity int, operator string) (*entity.Discount, error)
}

type DiscountUseCaseImpl struct {
	repository repository.DiscountRepository
}

func (useCase *DiscountUseCaseImpl) FindAll(offset int, limit int) []entity.Discount {
	return useCase.repository.FindAll(offset, limit)
}

func (useCase *DiscountUseCaseImpl) FindByID(id int) (*entity.Discount, error) {
	return useCase.repository.FindByID(id)
}

func (useCase *DiscountUseCaseImpl) FindByCode(code string) (*entity.Discount, error) {
	return useCase.repository.FindByCode(code)
}

func (useCase *DiscountUseCaseImpl) Create(dto dto.DiscountRequestBody) (*entity.Discount, error) {
	//	create object discount
	discount := entity.Discount{
		ID:                0,
		Name:              dto.Name,
		Code:              dto.Code,
		Quantity:          dto.Quantity,
		RemainingQuantity: 0,
		Type:              dto.Type,
		Value:             dto.Value,
		StartDate: sql.NullTime{
			Time:  dto.StartDate,
			Valid: true,
		},
		EndDate: sql.NullTime{
			Time:  dto.EndDate,
			Valid: true,
		},
		CreatedByID: &dto.CreatedBy,
	}

	//	insert data
	data, err := useCase.repository.Create(discount)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (useCase *DiscountUseCaseImpl) Update(id int, dto dto.DiscountRequestBody) (*entity.Discount, error) {
	//find data discount based from id
	discount, err := useCase.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	//assign update data
	discount.Name = dto.Name
	discount.Code = dto.Code
	discount.Quantity = dto.Quantity
	discount.RemainingQuantity = dto.RemainingQuantity
	discount.Type = dto.Type
	discount.Value = dto.Value
	discount.UpdatedByID = &dto.UpdatedBy
	discount.StartDate.Time = dto.StartDate
	discount.EndDate.Time = dto.EndDate

	//update data discount to repository
	updateDiscount, err := useCase.repository.Update(*discount)
	if err != nil {
		return nil, err
	}
	return updateDiscount, nil
}

func (useCase *DiscountUseCaseImpl) Delete(id int) error {
	//	find data discount based from id
	discount, err := useCase.repository.FindByID(id)
	if err != nil {
		return err
	}
	//	call repository delete
	err = useCase.repository.Delete(*discount)
	if err != nil {
		return err
	}
	return nil
}

func (useCase *DiscountUseCaseImpl) UpdateRemainingQuantity(id int, quantity int, operator string) (*entity.Discount, error) {
	//find discount data based from id
	discount, err := useCase.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	//check operator update discount
	if operator == "+" {
		discount.RemainingQuantity = discount.RemainingQuantity + int64(quantity)
	} else if operator == "-" {
		discount.RemainingQuantity = discount.RemainingQuantity - int64(quantity)
	} else {
		return nil, errors.New("operator not handled")
	}

	//update data discount to repository
	updateDiscount, err := useCase.repository.Update(*discount)
	if err != nil {
		return nil, err
	}
	return updateDiscount, nil
}

func NewDiscountUseCase(repository repository.DiscountRepository) DiscountUseCase {
	return &DiscountUseCaseImpl{repository: repository}
}
