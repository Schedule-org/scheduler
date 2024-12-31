package dto

import "github.com/hebertzin/scheduler/internal/domains"

type UserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func MapToUserDTO(user *domains.User) *UserDTO {
	return &UserDTO{
		Name:  user.Name,
		Email: user.Email,
	}
}
