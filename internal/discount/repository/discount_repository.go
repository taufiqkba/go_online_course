package repository

import (
	"go_online_course/internal/discount/entity"
	"go_online_course/pkg/utils"

	"gorm.io/gorm"
)

type DiscountRepository interface {
	FindAll(offset int, limit int) []entity.Discount
	FindByID(id int) (*entity.Discount, error)
	FindByCode(code string) (*entity.Discount, error)
	Create(entity entity.Discount) (*entity.Discount, error)
	Update(entity entity.Discount) (*entity.Discount, error)
	Delete(entity entity.Discount) error
}

type DiscountRepositoryImpl struct {
	db *gorm.DB
}

func (repository *DiscountRepositoryImpl) FindAll(offset int, limit int) []entity.Discount {
	var discounts []entity.Discount

	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&discounts)
	return discounts
}

func (repository *DiscountRepositoryImpl) FindByID(id int) (*entity.Discount, error) {
	var discount entity.Discount

	err := repository.db.First(&discount, id).Error
	if err != nil {
		return nil, err
	}
	return &discount, nil
}

func (repository *DiscountRepositoryImpl) FindByCode(code string) (*entity.Discount, error) {
	var discount entity.Discount

	err := repository.db.Where("code = ?", code).First(&discount).Error
	if err != nil {
		return nil, err
	}
	return &discount, nil
}

func (repository *DiscountRepositoryImpl) Create(entity entity.Discount) (*entity.Discount, error) {
	err := repository.db.Create(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repository *DiscountRepositoryImpl) Update(entity entity.Discount) (*entity.Discount, error) {
	err := repository.db.Save(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repository *DiscountRepositoryImpl) Delete(entity entity.Discount) error {
	err := repository.db.Delete(&entity).Error
	if err != nil {
		return err
	}
	return nil
}

func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	return &DiscountRepositoryImpl{db: db}
}
