package repository

import (
	"github.com/stretchr/testify/mock"
	"go_online_course/internal/user/entity"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UserRepositoryMock) FindByEmail(email string) (*entity.User, error) {
	panic("unimplemented")
}

func (repository *UserRepositoryMock) FindAll(offset int, limit int) []entity.User {
	panic("unimplemented")
}

func (repository *UserRepositoryMock) FindById(id int) (*entity.User, error) {
	//arguments := repository.Mock.Called(id)
	panic("unimplemented")
}

func (repository *UserRepositoryMock) Create(entity entity.User) (*entity.User, error) {
	panic("unimplemented")
}

func (repository *UserRepositoryMock) Update(entity entity.User) (*entity.User, error) {
	panic("unimplemented")
}

func (repository *UserRepositoryMock) Delete(entity entity.User) error {
	panic("unimplemented")
}
