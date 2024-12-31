package dto

import "github.com/hebertzin/scheduler/internal/domains"

type UserDTO struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	City  string `json:"city"`
}

func MapToUserDTO(user *domains.User) *UserDTO {
	return &UserDTO{
		Name:  user.Name,
		Email: user.Email,
	}
}
