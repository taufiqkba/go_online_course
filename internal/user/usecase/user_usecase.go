package usecase

import (
	"errors"
	"go_online_course/internal/user/dto"
	"go_online_course/internal/user/entity"
	"go_online_course/internal/user/repository"
	"go_online_course/pkg/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCase interface {
	FindAll(offset int, limit int) []entity.User
	FindById(id int) (*entity.User, error)
	Count() int
	FindByEmail(email string) (*entity.User, error)
	Create(userDto dto.UserRequestBody) (*entity.User, error)
	Update(id int, dto dto.UserRequestBody) (*entity.User, error)
	Delete(id int) error
}

type UserUseCaseImpl struct {
	repository repository.UserRepository
}

func (usecase *UserUseCaseImpl) Count() int {
	return usecase.repository.Count()
}

// Create implements UserUseCase
func (usecase *UserUseCaseImpl) Create(userDto dto.UserRequestBody) (*entity.User, error) {
	//	Find by email
	checkUser, err := usecase.repository.FindByEmail(*userDto.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if checkUser != nil {
		return nil, errors.New("email has been registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := entity.User{
		ID:           0,
		Name:         *userDto.Name,
		Email:        *userDto.Email,
		Password:     string(hashedPassword),
		CodeVerified: utils.RandomString(12),
	}
	if userDto.CreatedBy != nil {
		user.CreatedByID = userDto.CreatedBy
	}

	dataUser, err := usecase.repository.Create(user)
	if err != nil {
		return nil, err
	}
	return dataUser, nil
}

// Delete implements UserUseCase
func (usecase *UserUseCaseImpl) Delete(id int) error {
	//TODO Implement me
	panic("unimplemented")
}

// FindAll implements UserUseCase
func (usecase *UserUseCaseImpl) FindAll(offset int, limit int) []entity.User {
	return usecase.repository.FindAll(offset, limit)
}

// FindByEmail implements UserUseCase
func (usecase *UserUseCaseImpl) FindByEmail(email string) (*entity.User, error) {
	return usecase.repository.FindByEmail(email)
}

// FindById implements UserUseCase
func (usecase *UserUseCaseImpl) FindById(id int) (*entity.User, error) {
	return usecase.repository.FindById(id)
}

// Update implements UserUseCase
func (usecase *UserUseCaseImpl) Update(id int, dto dto.UserRequestBody) (*entity.User, error) {
	user, err := usecase.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	if dto.Name != nil {
		user.Name = *dto.Name
	}
	if dto.Email != nil {
		if user.Email != *dto.Email {
			user.Email = *dto.Email
		}
	}

	if dto.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*dto.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	if dto.UpdatedBy != nil {
		user.UpdatedByID = dto.UpdatedBy
	}
	updateUser, err := usecase.repository.Update(*user)
	if err != nil {
		return nil, err
	}
	return updateUser, nil
}

func NewUserUseCase(repository repository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{repository}
}
