package usecase

import (
	register "go_online_course/internal/register/dto"
	"go_online_course/internal/user/dto"
	"go_online_course/internal/user/usecase"
	"go_online_course/pkg/mail/sendgrid"
)

type RegisterUseCase interface {
	Register(userDto dto.UserRequestBody) error
}

type RegisterUseCaseImpl struct {
	userUseCase usecase.UserUseCase
	mail        sendgrid.Mail
}

func NewRegisterUseCase(userUseCase usecase.UserUseCase, mail sendgrid.Mail) RegisterUseCase {
	return &RegisterUseCaseImpl{userUseCase, mail}
}

func (ru *RegisterUseCaseImpl) Register(userDto dto.UserRequestBody) error {
	//	check into user module
	user, err := ru.userUseCase.Create(userDto)
	if err != nil {
		return err
	}
	//	send email
	email := register.CrateEmailVerification{
		Subject:          "Verification Code",
		Email:            user.Email,
		VerificationCode: user.CodeVerified,
	}
	ru.mail.SendVerificationEmail(user.Email, email)

	return nil
}
