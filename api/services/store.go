package services

import (
	"context"
	"fmt"

	"github.com/g-villarinho/flash-buy-api/models"
	"github.com/g-villarinho/flash-buy-api/pkgs"
	"github.com/g-villarinho/flash-buy-api/repositories"
)

type StoreService interface {
	CreateStore(ctx context.Context, store *models.Store) (*models.CreateStoreResponse, error)
	GetUserFirstStore(ctx context.Context, userID string) (*models.StoreResponse, error)
	GetStoreByStoreID(ctx context.Context, storeID, userID string) (*models.StoreResponse, error)
	GetStoresByUserID(ctx context.Context, userID string) ([]models.StoreResponse, error)
	UpdateStore(ctx context.Context, userID, storeID, name string) error
	DeleteStore(ctx context.Context, storeID, userID string) error
}

type storeService struct {
	di *pkgs.Di
	sr repositories.StoreRepository
}

func NewStoreService(di *pkgs.Di) (StoreService, error) {
	sr, err := pkgs.Invoke[repositories.StoreRepository](di)
	if err != nil {
		return nil, err
	}
	return &storeService{
		di: di,
		sr: sr,
	}, nil
}

func (s *storeService) CreateStore(ctx context.Context, store *models.Store) (*models.CreateStoreResponse, error) {
	if err := s.sr.CreateStore(ctx, store); err != nil {
		return nil, fmt.Errorf("create store: %w", err)
	}

	return &models.CreateStoreResponse{
		StoreId: store.ID.String(),
	}, nil
}

func (s *storeService) GetUserFirstStore(ctx context.Context, userID string) (*models.StoreResponse, error) {
	store, err := s.sr.FindFirstStoreByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("find first store by user ID %s: %w", userID, err)
	}

	if store == nil {
		return nil, models.ErrStoreNotFound
	}

	return store.ToStoreResponse(), nil
}

func (s *storeService) GetStoreByStoreID(ctx context.Context, storeId string, userID string) (*models.StoreResponse, error) {
	store, err := s.sr.GetStoreByID(ctx, storeId)
	if err != nil {
		return nil, fmt.Errorf("get store by id %s: %w", storeId, err)
	}

	if store == nil {
		return nil, models.ErrStoreNotFound
	}

	if store.UserID.String() != userID {
		return nil, models.ErrStoreNotPertenence
	}

	return store.ToStoreResponse(), nil
}

func (s *storeService) GetStoresByUserID(ctx context.Context, userID string) ([]models.StoreResponse, error) {
	stores, err := s.sr.GetStoresByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get stores by user ID %s: %w", userID, err)
	}

	var storesResponse []models.StoreResponse
	for _, store := range stores {
		storesResponse = append(storesResponse, *store.ToStoreResponse())
	}

	return storesResponse, nil
}

func (s *storeService) UpdateStore(ctx context.Context, userID, storeID, name string) error {
	store, err := s.sr.GetStoreByID(ctx, storeID)
	if err != nil {
		return fmt.Errorf("get store by id %s: %w", storeID, err)
	}

	if store == nil {
		return models.ErrStoreNotFound
	}

	if store.UserID.String() != userID {
		return models.ErrStoreNotPertenence
	}

	store.Name = name

	if err := s.sr.UpdateStore(ctx, store); err != nil {
		return fmt.Errorf("update store: %w", err)
	}

	return nil
}

func (s *storeService) DeleteStore(ctx context.Context, storeID string, userID string) error {
	store, err := s.sr.GetStoreByID(ctx, storeID)
	if err != nil {
		return fmt.Errorf("get store by id %s: %w", storeID, err)
	}

	if store == nil {
		return models.ErrStoreNotFound
	}

	if store.UserID.String() != userID {
		return models.ErrStoreNotPertenence
	}

	if err := s.sr.DeleteStore(ctx, storeID); err != nil {
		return fmt.Errorf("delete store: %w", err)
	}

	return nil
}
