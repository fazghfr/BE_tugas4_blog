package repository

import (
	"context"
	"errors"
	"tugas4/entitiy"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository interface {
	AddComment(ctx context.Context, comment entitiy.Comment) (entitiy.Comment, error)
	GetCommentsByBlogID(ctx context.Context, blogID uuid.UUID) ([]entitiy.Comment, error)
}

type commentConnection struct {
	connection *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentConnection{
		connection: db,
	}
}

func (c *commentConnection) AddComment(ctx context.Context, comment entitiy.Comment) (entitiy.Comment, error) {
	comment.ID = uuid.New()
	err := c.connection.Create(&comment).Error
	if err != nil {
		return entitiy.Comment{}, err
	}

	// Retrieve the blog to update
	blog := entitiy.Blog{}
	if err := c.connection.Where("id = ?", comment.BlogID).First(&blog).Error; err != nil {
		return entitiy.Comment{}, err
	}

	// Update the blog's comments array
	blog.Comments = append(blog.Comments, comment)

	// Save the updated blog to the database
	if err := c.connection.Save(&blog).Error; err != nil {
		return entitiy.Comment{}, err
	}

	return comment, nil
}

func (c *commentConnection) GetCommentsByBlogID(ctx context.Context, blogID uuid.UUID) ([]entitiy.Comment, error) {
	var comments []entitiy.Comment
	err := c.connection.Where("blog_id = ?", blogID).Find(&comments).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return comments, nil
}
