package services

import (
	"context"
	"fmt"

	"github.com/g-villarinho/flash-buy-api/models"
	"github.com/g-villarinho/flash-buy-api/pkgs"
	"github.com/g-villarinho/flash-buy-api/repositories"
)

type UserService interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, ID string) (*models.User, error)
}

type userService struct {
	di *pkgs.Di
	ur repositories.UserRepository
}

func NewUserService(di *pkgs.Di) (UserService, error) {
	ur, err := pkgs.Invoke[repositories.UserRepository](di)
	if err != nil {
		return nil, err
	}
	return &userService{
		di: di,
		ur: ur,
	}, nil
}

func (u *userService) CreateUser(ctx context.Context, user models.User) error {
	userFromEmail, err := u.ur.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return fmt.Errorf("get user by email: %w", err)
	}

	if userFromEmail != nil {
		return models.ErrUserAlreadyExists
	}

	if err := u.ur.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}

func (u *userService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := u.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("get user by email: %w", err)
	}

	if user == nil {
		return nil, models.ErrUserNotFound
	}

	return user, nil
}

func (u *userService) GetUserByID(ctx context.Context, ID string) (*models.User, error) {
	user, err := u.ur.GetUserByID(ctx, ID)
	if err != nil {
		return nil, fmt.Errorf("get user by ID %s: %w", ID, err)
	}

	if user == nil {
		return nil, models.ErrUserNotFound
	}

	return user, nil
}
