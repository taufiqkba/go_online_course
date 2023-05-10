package usecase

import (
	"go_online_course/internal/admin/dto"
	"go_online_course/internal/admin/entity"
	"go_online_course/internal/admin/repository"
)

type AdminUseCase interface {
	FindAll(offset int, limit int) []entity.Admin
	FindByID(id int) (*entity.Admin, error)
	FindByEmail(email string) (*entity.Admin, error)
	Create(dto dto.AdminRequestBody) (*entity.Admin, error) //why depend to AdminRequestBody
	Update(dto dto.AdminRequestBody) (*entity.Admin, error)
	Delete(id int) error
}

type AdminUseCaseImpl struct {
	repository repository.AdminRepository
}

// Create implements AdminUseCase
func (usecase *AdminUseCaseImpl) Create(dto dto.AdminRequestBody) (*entity.Admin, error) {
	panic("unimplemented")
}

// Delete implements AdminUseCase
func (usecase *AdminUseCaseImpl) Delete(id int) error {
	panic("unimplemented")
}

// FindAll implements AdminUseCase
func (usecase *AdminUseCaseImpl) FindAll(offset int, limit int) []entity.Admin {
	panic("unimplemented")
}

// FindByEmail implements AdminUseCase
func (usecase *AdminUseCaseImpl) FindByEmail(email string) (*entity.Admin, error) {
	return usecase.repository.FindByEmail(email)
}

// FindByID implements AdminUseCase
func (usecase *AdminUseCaseImpl) FindByID(id int) (*entity.Admin, error) {
	panic("unimplemented")
}

// Update implements AdminUseCase
func (usecase *AdminUseCaseImpl) Update(dto dto.AdminRequestBody) (*entity.Admin, error) {
	panic("unimplemented")
}

func NewAdminUseCase(repository repository.AdminRepository) AdminUseCase {
	return &AdminUseCaseImpl{repository}
}

// type AdminUseCase interface {
// 	FindAll(offset int, limit int) []entity.Admin
// 	FindById(id int) (*entity.Admin, error)
// 	FindByEmail(email string) (*entity.Admin, error)
// 	Create(dto dto.AdminRequestBody) (*entity.Admin, error)
// 	Update(id int, dto dto.AdminRequestBody) (*entity.Admin, error)
// 	Delete(id int) error
// }

// type AdminUseCaseImpl struct {
// 	repository repository.AdminRepository
// }

// // Create implements AdminUseCase
// func (usecase *AdminUseCaseImpl) Create(dto dto.AdminRequestBody) (*entity.Admin, error) {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return nil, err
// 	}

// 	dataAdmin := entity.Admin{
// 		Name:        dto.Name,
// 		Email:       dto.Email,
// 		Password:    dto.Password,
// 		CreatedById: dto.CreatedBy,
// 	}

// 	admin, err := usecase.repository.Create(dataAdmin)

// }

// // Delete implements AdminUseCase
// func (usecase *AdminUseCaseImpl) Delete(id int) error {
// 	panic("unimplemented")
// }

// // FindAll implements AdminUseCase
// func (usecase *AdminUseCaseImpl) FindAll(offset int, limit int) []entity.Admin {
// 	panic("unimplemented")
// }

// // FindByEmail implements AdminUseCase
// func (usecase *AdminUseCaseImpl) FindByEmail(email string) (*entity.Admin, error) {
// 	return usecase.repository.FindByEmail(email)
// }

// // FindById implements AdminUseCase
// func (usecase *AdminUseCaseImpl) FindById(id int) (*entity.Admin, error) {
// 	return usecase.repository.FindById(id)
// }

// // Update implements AdminUseCase
// func (usecase *AdminUseCaseImpl) Update(id int, dto dto.AdminRequestBody) (*entity.Admin, error) {
// 	panic("unimplemented")
// }

// func NewAdminUseCase(repository repository.AdminRepository) AdminUseCase {
// 	return &AdminUseCaseImpl{repository}
// }
