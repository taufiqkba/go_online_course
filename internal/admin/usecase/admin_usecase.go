package usecase

import (
	"go_online_course/internal/admin/dto"
	"go_online_course/internal/admin/entity"
	"go_online_course/internal/admin/repository"

	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase interface {
	FindAll(offset int, limit int) []entity.Admin
	FindByID(id int) (*entity.Admin, error)
	FindByEmail(email string) (*entity.Admin, error)
	Create(dto dto.AdminRequestBody) (*entity.Admin, error) //why depend to AdminRequestBody
	Update(id int, dto dto.AdminRequestBody) (*entity.Admin, error)
	Delete(id int) error
}

type AdminUseCaseImpl struct {
	repository repository.AdminRepository
}

// Create implements AdminUseCase
func (usecase *AdminUseCaseImpl) Create(dto dto.AdminRequestBody) (*entity.Admin, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	dataAdmin := entity.Admin{
		Email:       dto.Email,
		Name:        dto.Name,
		Password:    string(hashedPassword),
		CreatedByID: dto.CreatedBy,
	}

	admin, err := usecase.repository.Create(dataAdmin)

	if err != nil {
		return nil, err
	}
	return admin, nil
}

// Delete implements AdminUseCase
func (usecase *AdminUseCaseImpl) Delete(id int) error {
	// Search from table admin based on id
	admin, err := usecase.repository.FindById(id)
	if err != nil {
		return err
	}

	if err := usecase.repository.Delete(*admin); err != nil {
		return err
	}

	return nil
}

// FindAll implements AdminUseCase
func (usecase *AdminUseCaseImpl) FindAll(offset int, limit int) []entity.Admin {
	return usecase.repository.FindAll(offset, limit)
}

// FindByEmail implements AdminUseCase
func (usecase *AdminUseCaseImpl) FindByEmail(email string) (*entity.Admin, error) {
	return usecase.repository.FindByEmail(email)
}

// FindByID implements AdminUseCase
func (usecase *AdminUseCaseImpl) FindByID(id int) (*entity.Admin, error) {
	return usecase.repository.FindById(id)
}

// Update implements AdminUseCase
func (usecase *AdminUseCaseImpl) Update(id int, dto dto.AdminRequestBody) (*entity.Admin, error) {
	// Search from table admin by id
	admin, err := usecase.repository.FindById(id)
	admin.Name = dto.Name

	// Validation email admin if not equal while updated
	if admin.Email != dto.Email {
		admin.Email = dto.Email
	}

	if dto.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*dto.Password), bcrypt.DefaultCost)

		if err != nil {
			return nil, err
		}
		admin.Password = string(hashedPassword)
	}

	if err != nil {
		return nil, err
	}

	admin.UpdatedByID = &dto.UpdatedBy

	// update data with call the repository admin
	updateAdmin, err := usecase.repository.Update(*admin)
	if err != nil {
		return nil, err
	}

	return updateAdmin, nil
}

func NewAdminUseCase(repository repository.AdminRepository) AdminUseCase {
	return &AdminUseCaseImpl{repository}
}
