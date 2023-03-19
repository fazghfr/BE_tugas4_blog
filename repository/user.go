package repository

import (
	"context"
	"tugas4/entitiy"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, user entitiy.User) (entitiy.User, error)
	GetAllUser(ctx context.Context) ([]entitiy.User, error)
	FindUserByEmail(ctx context.Context, email string) (entitiy.User, error)
	FindUserByID(ctx context.Context, userID uuid.UUID) (entitiy.User, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	UpdateUser(ctx context.Context, user entitiy.User) error
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) RegisterUser(ctx context.Context, user entitiy.User) (entitiy.User, error) {
	user.ID = uuid.New()
	uc := db.connection.Create(&user)
	if uc.Error != nil {
		return entitiy.User{}, uc.Error
	}
	return user, nil
}

func (db *userConnection) GetAllUser(ctx context.Context) ([]entitiy.User, error) {
	var listUser []entitiy.User
	tx := db.connection.Find(&listUser)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listUser, nil
}

func (db *userConnection) FindUserByEmail(ctx context.Context, email string) (entitiy.User, error) {
	var user entitiy.User
	ux := db.connection.Where("email = ?", email).Take(&user)
	if ux.Error != nil {
		return user, ux.Error
	}
	return user, nil
}

func (db *userConnection) FindUserByID(ctx context.Context, userID uuid.UUID) (entitiy.User, error) {
	var user entitiy.User
	ux := db.connection.Where("id = ?", userID).Take(&user)
	if ux.Error != nil {
		return user, ux.Error
	}
	return user, nil
}

func (db *userConnection) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	uc := db.connection.Delete(&entitiy.User{}, &userID)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func (db *userConnection) UpdateUser(ctx context.Context, user entitiy.User) error {
	uc := db.connection.Updates(&user)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}
