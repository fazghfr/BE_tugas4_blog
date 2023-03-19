package service

import (
	"context"
	"fmt"
	"tugas4/dto"
	"tugas4/entitiy"
	"tugas4/repository"
	utils "tugas4/util"

	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type UserService interface {
	RegisterUser(ctx context.Context, userDTO dto.UserCreateDTO) (entitiy.User, error)
	GetAllUser(ctx context.Context) ([]entitiy.User, error)
	FindUserByEmail(ctx context.Context, email string) (entitiy.User, error)
	Verify(ctx context.Context, email string, password string) (bool, error)
	CheckUser(ctx context.Context, email string) (bool, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	UpdateUser(ctx context.Context, userDTO dto.UserUpdateDTO) error
	MeUser(ctx context.Context, userID uuid.UUID) (entitiy.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepository: ur,
	}
}

func (us *userService) RegisterUser(ctx context.Context, userDTO dto.UserCreateDTO) (entitiy.User, error) {
	user := entitiy.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(userDTO))
	if err != nil {
		return user, err
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return user, err
	}
	user.Password = hashedPassword

	return us.userRepository.RegisterUser(ctx, user)
}

func (us *userService) GetAllUser(ctx context.Context) ([]entitiy.User, error) {
	return us.userRepository.GetAllUser(ctx)
}

func (us *userService) FindUserByEmail(ctx context.Context, email string) (entitiy.User, error) {
	return us.userRepository.FindUserByEmail(ctx, email)
}

func (us *userService) Verify(ctx context.Context, email string, password string) (bool, error) {
	res, err := us.userRepository.FindUserByEmail(ctx, email)
	fmt.Println(res)
	if err != nil {
		return false, err
	}
	CheckPassword, err := utils.CheckPassword(res.Password, []byte(password))
	if err != nil {
		return false, err
	}
	if res.Email == email && CheckPassword {
		return true, nil
	}
	return false, nil
}

func (us *userService) CheckUser(ctx context.Context, email string) (bool, error) {
	result, err := us.userRepository.FindUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	if result.Email == "" {
		return false, nil
	}
	return true, nil
}

func (us *userService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return us.userRepository.DeleteUser(ctx, userID)
}

func (us *userService) UpdateUser(ctx context.Context, userDTO dto.UserUpdateDTO) error {
	user := entitiy.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(userDTO))
	if err != nil {
		return err
	}
	return us.userRepository.UpdateUser(ctx, user)
}

func (us *userService) MeUser(ctx context.Context, userID uuid.UUID) (entitiy.User, error) {
	return us.userRepository.FindUserByID(ctx, userID)
}
