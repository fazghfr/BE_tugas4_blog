package dto

import "github.com/google/uuid"

type BlogCreateDTO struct {
	Title   string    `json:"title" binding:"required"`
	Content string    `json:"content" binding:"required"`
	UserID  uuid.UUID `json:"userid"`
}

type CommentDTO struct {
	Content string `json:"content" binding:"required"`
}
