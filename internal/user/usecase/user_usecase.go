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
	FindByEmail(email string) (*entity.User, error)
	Create(userDto dto.UserRequestBody) (*entity.User, error)
	Update(userDto dto.UserRequestBody) (*entity.User, error)
	Delete(id int) error
}

type UserUseCaseImpl struct {
	repository repository.UserRepository
}

func NewUserUseCase(repository repository.UserRepository) *UserUseCaseImpl {
	return &UserUseCaseImpl{repository}
}

func (uu UserUseCaseImpl) FindAll(offset int, limit int) []entity.User {
	//TODO implement me
	panic("implement me")
}

func (uu UserUseCaseImpl) FindById(id int) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (uu UserUseCaseImpl) Create(userDto dto.UserRequestBody) (*entity.User, error) {
	//	Find by email
	checkUser, err := uu.repository.FindByEmail(*userDto.Email)
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
	dataUser, err := uu.repository.Create(user)
	if err != nil {
		return nil, err
	}
	return dataUser, nil
}

func (uu UserUseCaseImpl) Update(userDto dto.UserRequestBody) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (uu UserUseCaseImpl) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
