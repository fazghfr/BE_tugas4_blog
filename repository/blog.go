package repository

import (
	"context"
	"tugas4/entitiy"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogRepository interface {
	CreateBlog(ctx context.Context, blog entitiy.Blog) (entitiy.Blog, error)
	GetAllBlogs(ctx context.Context) ([]entitiy.Blog, error)
	FindBlogByID(ctx context.Context, blogID uuid.UUID) (entitiy.Blog, error)
	LikeBlog(ctx context.Context, blogID uuid.UUID) (entitiy.Blog, error)
}

type blogConnection struct {
	connection *gorm.DB
}

func NewBlogRepository(db *gorm.DB) BlogRepository {
	return &blogConnection{
		connection: db,
	}
}

func (db *blogConnection) CreateBlog(ctx context.Context, blog entitiy.Blog) (entitiy.Blog, error) {
	blog.ID = uuid.New()
	uc := db.connection.Create(&blog)
	if uc.Error != nil {
		return entitiy.Blog{}, uc.Error
	}
	return blog, nil
}

func (db *blogConnection) LikeBlog(ctx context.Context, blogID uuid.UUID) (entitiy.Blog, error) {
	var blog entitiy.Blog
	tx := db.connection.Where("id = ?", blogID).Take(&blog)
	if tx.Error != nil {
		return blog, tx.Error
	}
	blog.Likes++
	tx = db.connection.Save(&blog)
	if tx.Error != nil {
		return blog, tx.Error
	}
	return blog, nil
}

func (db *blogConnection) GetAllBlogs(ctx context.Context) ([]entitiy.Blog, error) {
	var listBlog []entitiy.Blog
	tx := db.connection.Preload("Comments").Find(&listBlog)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listBlog, nil
}

func (db *blogConnection) FindBlogByID(ctx context.Context, blogID uuid.UUID) (entitiy.Blog, error) {
	var blog entitiy.Blog
	ux := db.connection.Preload("Comments").Where("id = ?", blogID).Take(&blog)
	if ux.Error != nil {
		return blog, ux.Error
	}
	return blog, nil
}
