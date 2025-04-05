package repositories

import (
	"context"
	"fmt"

	"github.com/g-villarinho/flash-buy-api/models"
	"github.com/g-villarinho/flash-buy-api/persistence"
	"github.com/g-villarinho/flash-buy-api/pkgs"
)

type SizeRepository interface {
	CreateSize(ctx context.Context, size *models.Size) error
	GetSizesPagedList(ctx context.Context, storeID string, pag models.SizePagination) (*models.PaginatedResponse, error)
	GetSizeByID(ctx context.Context, ID string) (*models.Size, error)
	DeleteSize(ctx context.Context, ID string) error
}

type sizeRepository struct {
	di   *pkgs.Di
	repo persistence.Repository
}

func NewSizeRepository(di *pkgs.Di) (SizeRepository, error) {
	repo, err := pkgs.Invoke[persistence.Repository](di)
	if err != nil {
		return nil, err
	}

	return &sizeRepository{
		di:   di,
		repo: repo,
	}, nil
}

func (s *sizeRepository) CreateSize(ctx context.Context, size *models.Size) error {
	if err := s.repo.Create(ctx, size); err != nil {
		return err
	}

	return nil
}

func (s *sizeRepository) GetSizesPagedList(ctx context.Context, storeID string, pag models.SizePagination) (*models.PaginatedResponse, error) {
	var sizes []models.Size

	opts := []persistence.QueryOption{}

	opts = append(opts, persistence.WithConditions("store_id = ?", storeID))

	if pag.Name != nil {
		opts = append(opts, persistence.WithConditions("name LIKE ?", fmt.Sprintf("%%%s%%", *pag.Name)))
	}

	result, err := s.repo.Paginate(ctx, &sizes, *pag.Pagination, opts...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *sizeRepository) GetSizeByID(ctx context.Context, ID string) (*models.Size, error) {
	var size models.Size
	if err := s.repo.FindByID(ctx, ID, &size); err != nil {
		if err == persistence.ErrRecordNotFound {
			return nil, models.ErrSizeNotFound
		}

		return nil, err
	}

	return &size, nil
}

func (s *sizeRepository) DeleteSize(ctx context.Context, ID string) error {
	if err := s.repo.Delete(ctx, ID, &models.Size{}); err != nil {
		return err
	}

	return nil
}
