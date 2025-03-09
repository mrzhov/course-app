package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
}

func NewUserResponse(t User) UserResponse {
	return UserResponse{
		Id:    t.ID,
		Email: t.Email,
	}
}

type CreateBody struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required, min=8"`
}

type PatchBody struct {
	Email *string `json:"email" validate:"required, email"`
}
