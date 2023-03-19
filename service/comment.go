package service

import (
	"context"
	"tugas4/dto"
	"tugas4/entitiy"
	"tugas4/repository"

	"github.com/google/uuid"
)

type CommentService interface {
	AddComment(ctx context.Context, blogID uuid.UUID, commentDTO dto.CommentDTO) (entitiy.Comment, error)
}

type commentService struct {
	commentRepo repository.CommentRepository
}

func NewCommentService(commentRepo repository.CommentRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}

func (cs *commentService) AddComment(ctx context.Context, blogID uuid.UUID, commentDTO dto.CommentDTO) (entitiy.Comment, error) {
	c := entitiy.Comment{
		BlogID:  blogID,
		Content: commentDTO.Content,
	}
	return cs.commentRepo.AddComment(ctx, c)
}
