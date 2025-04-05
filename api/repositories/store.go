package repositories

import (
	"context"

	"github.com/g-villarinho/flash-buy-api/models"
	"github.com/g-villarinho/flash-buy-api/persistence"
	"github.com/g-villarinho/flash-buy-api/pkgs"
)

type StoreRepository interface {
	CreateStore(ctx context.Context, store *models.Store) error
	GetStoreByID(ctx context.Context, ID string) (*models.Store, error)
	FindFirstStoreByUserID(ctx context.Context, userID string) (*models.Store, error)
	GetStoresByUserID(ctx context.Context, userID string) ([]models.Store, error)
	UpdateStore(ctx context.Context, store *models.Store) error
	DeleteStore(ctx context.Context, ID string) error
}

type storeRepository struct {
	di   *pkgs.Di
	repo persistence.Repository
}

func NewStoreRepository(di *pkgs.Di) (StoreRepository, error) {
	repo, err := pkgs.Invoke[persistence.Repository](di)
	if err != nil {
		return nil, err
	}
	return &storeRepository{
		di:   di,
		repo: repo,
	}, nil
}

func (s *storeRepository) CreateStore(ctx context.Context, store *models.Store) error {
	if err := s.repo.Create(ctx, store); err != nil {
		return err
	}

	return nil
}

func (s *storeRepository) GetStoreByID(ctx context.Context, ID string) (*models.Store, error) {
	var store models.Store
	if err := s.repo.FindByID(ctx, ID, &store); err != nil {
		if err == persistence.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &store, nil
}

func (s *storeRepository) FindFirstStoreByUserID(ctx context.Context, userID string) (*models.Store, error) {
	var store models.Store
	if err := s.repo.FindOne(ctx, &store, persistence.WithConditions("user_id = ?", userID)); err != nil {
		if err == persistence.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &store, nil
}

func (s *storeRepository) GetStoresByUserID(ctx context.Context, userID string) ([]models.Store, error) {
	var stores []models.Store
	if err := s.repo.FindAll(ctx, &stores, persistence.WithConditions("user_id = ?", userID)); err != nil {
		return nil, err
	}

	return stores, nil
}

func (s *storeRepository) UpdateStore(ctx context.Context, store *models.Store) error {
	if err := s.repo.Update(ctx, store); err != nil {
		return err
	}

	return nil
}

func (s *storeRepository) DeleteStore(ctx context.Context, ID string) error {
	if err := s.repo.Delete(ctx, ID, &models.Store{}); err != nil {
		return err
	}

	return nil
}
