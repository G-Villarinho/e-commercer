package repositories

import (
	"context"
	"errors"

	"github.com/g-villarinho/flash-buy-api/models"
	"github.com/g-villarinho/flash-buy-api/persistence"
	"github.com/g-villarinho/flash-buy-api/pkgs"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, ID string) (*models.User, error)
}

type userRepository struct {
	di   *pkgs.Di
	repo persistence.Repository
}

func NewUserRepository(di *pkgs.Di) (UserRepository, error) {
	repo, err := pkgs.Invoke[persistence.Repository](di)
	if err != nil {
		return nil, err
	}
	return &userRepository{
		di:   di,
		repo: repo,
	}, nil
}

func (u *userRepository) CreateUser(ctx context.Context, user models.User) error {
	if err := u.repo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := u.repo.FindOne(ctx, &user, persistence.WithConditions("email = ?", email)); err != nil {
		if errors.Is(err, persistence.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (u *userRepository) GetUserByID(ctx context.Context, ID string) (*models.User, error) {
	var user models.User
	if err := u.repo.FindByID(ctx, ID, &user); err != nil {
		if errors.Is(err, persistence.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
