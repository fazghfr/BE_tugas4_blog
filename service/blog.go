package service

import (
	"context"
	"tugas4/dto"
	"tugas4/entitiy"
	"tugas4/repository"

	"github.com/google/uuid"
)

type BlogService interface {
	CreateBlog(ctx context.Context, blog dto.BlogCreateDTO) (entitiy.Blog, error)
	GetAllBlogs(ctx context.Context) ([]entitiy.Blog, error)
	FindBlogByID(ctx context.Context, blogID uuid.UUID) (entitiy.Blog, error)
	LikeBlog(ctx context.Context, blogID uuid.UUID) (entitiy.Blog, error)
}

type blogService struct {
	blogRepo repository.BlogRepository
}

func NewBlogService(blogRepo repository.BlogRepository) BlogService {
	return &blogService{
		blogRepo: blogRepo,
	}
}

func (bs *blogService) LikeBlog(ctx context.Context, blogID uuid.UUID) (entitiy.Blog, error) {
	return bs.blogRepo.LikeBlog(ctx, blogID)
}

func (bs *blogService) CreateBlog(ctx context.Context, blog dto.BlogCreateDTO) (entitiy.Blog, error) {
	b := entitiy.Blog{
		Title:   blog.Title,
		Content: blog.Content,
		UserID:  blog.UserID,
	}
	return bs.blogRepo.CreateBlog(ctx, b)
}

func (bs *blogService) GetAllBlogs(ctx context.Context) ([]entitiy.Blog, error) {
	return bs.blogRepo.GetAllBlogs(ctx)
}

func (bs *blogService) FindBlogByID(ctx context.Context, blogID uuid.UUID) (entitiy.Blog, error) {
	return bs.blogRepo.FindBlogByID(ctx, blogID)
}
