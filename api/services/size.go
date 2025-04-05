package services

import (
	"context"
	"fmt"

	"github.com/g-villarinho/flash-buy-api/models"
	"github.com/g-villarinho/flash-buy-api/pkgs"
	"github.com/g-villarinho/flash-buy-api/repositories"
)

type SizeService interface {
	CreateSize(ctx context.Context, userID string, size models.Size) error
	GetSizesPagedList(ctx context.Context, storeID string, pag models.SizePagination) (*models.PaginatedResponse, error)
	GetSizeByID(ctx context.Context, sizeID, storeID string) (*models.SizeResponse, error)
	DeleteSize(ctx context.Context, userID, storeID, sizeID string) error
}

type sizeService struct {
	di *pkgs.Di
	ss StoreService
	sr repositories.SizeRepository
}

func NewSizeService(di *pkgs.Di) (SizeService, error) {
	ss, err := pkgs.Invoke[StoreService](di)
	if err != nil {
		return nil, err
	}

	sr, err := pkgs.Invoke[repositories.SizeRepository](di)
	if err != nil {
		return nil, err
	}

	return &sizeService{
		di: di,
		ss: ss,
		sr: sr,
	}, nil
}

func (s *sizeService) CreateSize(ctx context.Context, userID string, size models.Size) error {
	_, err := s.ss.GetStoreByStoreID(ctx, size.StoreID.String(), userID)
	if err != nil {
		return err
	}

	if err := s.sr.CreateSize(ctx, &size); err != nil {
		return err
	}

	return nil
}

func (s *sizeService) GetSizesPagedList(ctx context.Context, storeID string, pag models.SizePagination) (*models.PaginatedResponse, error) {
	sizes, err := s.sr.GetSizesPagedList(ctx, storeID, pag)
	if err != nil {
		return nil, err
	}

	return &models.PaginatedResponse{
		Data:       toSizeResponseList(*sizes.Data.(*[]models.Size)),
		Total:      sizes.Total,
		TotalPages: sizes.TotalPages,
		Page:       sizes.Page,
		Limit:      sizes.Limit,
	}, nil
}

func (s *sizeService) GetSizeByID(ctx context.Context, sizeID, storeID string) (*models.SizeResponse, error) {
	size, err := s.sr.GetSizeByID(ctx, sizeID)
	if err != nil {
		return nil, err
	}

	if size.StoreID.String() != storeID {
		return nil, models.ErrSizeNotFound
	}

	return size.ToSizeResponse(), nil
}

func (s *sizeService) DeleteSize(ctx context.Context, userID string, storeID string, sizeID string) error {
	_, err := s.ss.GetStoreByStoreID(ctx, storeID, userID)
	if err != nil {
		return err
	}

	size, err := s.sr.GetSizeByID(ctx, sizeID)
	if err != nil {
		return fmt.Errorf("get size by id %s: %w", sizeID, err)
	}

	if size == nil {
		return models.ErrSizeNotFound
	}

	if size.StoreID.String() != storeID {
		return models.ErrSizeNotFound
	}

	if err := s.sr.DeleteSize(ctx, sizeID); err != nil {
		return fmt.Errorf("delete size %s: %w", sizeID, err)
	}

	return nil
}

func toSizeResponseList(sizes []models.Size) []models.SizeResponse {
	var response = make([]models.SizeResponse, len(sizes))

	for i, size := range sizes {
		response[i] = *size.ToSizeResponse()
	}

	return response
}
