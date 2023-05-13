package dto

type UserRequestBody struct {
	Name      *string `json:"name" binding:"required"`
	Email     *string `json:"email" binding:"required"`
	Password  *string `json:"password"`
	CreatedBy *int64  `json:"created_by"`
	UpdatedBy *int64  `json:"updated_by"`
}
