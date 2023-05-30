package usecase

import (
	"go_online_course/internal/user/entity"
	"go_online_course/internal/user/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var userUseCase = UserUseCaseImpl{repository: userRepository}

func TestUserUseCase_FindByIDSuccess(t *testing.T) {
	userData := entity.User{
		ID:   1,
		Name: "user",
	}
	userRepository.Mock.On("FindById", 1).Return(userData)

	user, err := userUseCase.FindById(1)
	assert.NotNil(t, user)
	assert.Nil(t, err)
}

func TestUserUseCaseImpl_FindByIdNotFound(t *testing.T) {
	userRepository.Mock.On("FindById", 2).Return(nil)
	user, err := userUseCase.FindById(2)
	assert.Nil(t, user)
	assert.Nil(t, err)
}
