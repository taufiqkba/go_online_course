package usecase

import (
	"go_online_course/internal/user/dto"
	"go_online_course/internal/user/usecase"
)

type RegisterUseCase interface {
	Register(userDto dto.UserRequestBody) error
}

type RegisterUseCaseImpl struct {
	userUseCase usecase.UserUseCase
}

func NewRegisterUseCase(userUseCase usecase.UserUseCase) RegisterUseCase {
	return &RegisterUseCaseImpl{userUseCase}
}

func (ru *RegisterUseCaseImpl) Register(userDto dto.UserRequestBody) error {
	//	check into user module
	_, err := ru.userUseCase.Create(userDto)
	if err != nil {
		return err
	}
	//	send email

	return nil
}
