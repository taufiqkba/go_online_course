package dto

import "go_online_course/internal/user/entity"

type ProfileResponseBody struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	IsVerified bool   `json:"is_verified"`
}

func CreateProfileResponse(user entity.User) ProfileResponseBody {
	IsVerified := false

	if user.EmailVerifiedAt.Valid {
		IsVerified = true
	}

	return ProfileResponseBody{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		IsVerified: IsVerified,
	}
}
