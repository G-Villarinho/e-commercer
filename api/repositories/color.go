package repositories

import (
	"context"
	"fmt"

	"github.com/g-villarinho/flash-buy-api/models"
	"github.com/g-villarinho/flash-buy-api/persistence"
	"github.com/g-villarinho/flash-buy-api/pkgs"
)

type ColorRepository interface {
	CreateColor(ctx context.Context, color *models.Color) error
	GetColorByHexAndStoreID(ctx context.Context, hex, storeID string) (*models.Color, error)
	GetColorsPagedList(ctx context.Context, storeID string, pag models.ColorPagination) (*models.PaginatedResponse, error)
	GetColorByID(ctx context.Context, ID string) (*models.Color, error)
	DeleteColor(ctx context.Context, ID string) error
}

type colorRepository struct {
	di   *pkgs.Di
	repo persistence.Repository
}

func NewColorRepository(di *pkgs.Di) (ColorRepository, error) {
	repo, err := pkgs.Invoke[persistence.Repository](di)
	if err != nil {
		return nil, err
	}

	return &colorRepository{
		di:   di,
		repo: repo,
	}, nil
}

func (c *colorRepository) CreateColor(ctx context.Context, color *models.Color) error {
	if err := c.repo.Create(ctx, color); err != nil {
		return err
	}

	return nil
}

func (c *colorRepository) GetColorByHexAndStoreID(ctx context.Context, hex, storeID string) (*models.Color, error) {
	var color models.Color

	if err := c.repo.FindOne(ctx, &color, persistence.WithConditions("hex = ? AND store_id = ?", hex, storeID)); err != nil {
		if err == persistence.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &color, nil
}

func (c *colorRepository) GetColorsPagedList(ctx context.Context, storeID string, pag models.ColorPagination) (*models.PaginatedResponse, error) {
	var colors []models.Color

	opts := []persistence.QueryOption{}

	opts = append(opts, persistence.WithConditions("store_id = ?", storeID))

	if pag.Name != nil {
		opts = append(opts, persistence.WithConditions("name LIKE ?", fmt.Sprintf("%%%s%%", *pag.Name)))
	}

	result, err := c.repo.Paginate(ctx, &colors, *pag.Pagination, opts...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *colorRepository) GetColorByID(ctx context.Context, ID string) (*models.Color, error) {
	var color models.Color

	if err := c.repo.FindByID(ctx, ID, &color); err != nil {
		if err == persistence.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &color, nil
}

func (c *colorRepository) DeleteColor(ctx context.Context, ID string) error {
	if err := c.repo.Delete(ctx, ID, &models.Color{}); err != nil {
		return err
	}

	return nil
}
