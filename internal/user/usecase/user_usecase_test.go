package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go_online_course/internal/user/entity"
	"go_online_course/internal/user/repository"
	"testing"
)

var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var userUseCase = UserUseCaseImpl{repository: userRepository}

func TestUserUseCase_FindByIDSuccess(t *testing.T) {
	userData := entity.User{
		ID:   1,
		Name: "taufiqkba",
	}
	userRepository.Mock.On("FindByID", 1).Return(userData)

	user, err := userUseCase.FindById(1)
	assert.Nil(t, user)
	assert.NotNil(t, err)
}
