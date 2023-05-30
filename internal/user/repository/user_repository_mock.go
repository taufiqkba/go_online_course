package repository

import (
	"github.com/stretchr/testify/mock"
	"go_online_course/internal/user/entity"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UserRepositoryMock) Count() int {
	//TODO implement me
	panic("implement me")
}

func (repository *UserRepositoryMock) FindByEmail(email string) (*entity.User, error) {
	panic("unimplemented")
}

func (repository *UserRepositoryMock) FindAll(offset int, limit int) []entity.User {
	panic("unimplemented")
}

func (repository *UserRepositoryMock) FindById(id int) (*entity.User, error) {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil, nil
	}
	user := arguments.Get(0).(entity.User)
	return &user, nil
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
