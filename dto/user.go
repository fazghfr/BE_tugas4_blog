package dto

import "github.com/google/uuid"

type UserLoginDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserCreateDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type UserUpdateDTO struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required"`
	Password string    `json:"password"`
}
