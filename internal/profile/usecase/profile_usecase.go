package usecase

import (
	"go_online_course/internal/profile/dto"
	"go_online_course/internal/user/usecase"
)

type ProfileUseCase interface {
	GetProfile(id int) (*dto.ProfileResponseBody, error)
}

type ProfileUseCaseImpl struct {
	userUseCase usecase.UserUseCase
}

// GetProfile implements ProfileUseCase
func (usecase *ProfileUseCaseImpl) GetProfile(id int) (*dto.ProfileResponseBody, error) {
	// Get Profile data
	user, err := usecase.userUseCase.FindById(id)

	if err != nil {
		return nil, err
	}

	userResponse := dto.CreateProfileResponse(*user)
	return &userResponse, nil
}

func NewProfileUseCase(userUseCase usecase.UserUseCase) ProfileUseCase {
	return &ProfileUseCaseImpl{userUseCase}
}
