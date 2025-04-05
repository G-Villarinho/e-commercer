package repositories

import (
	"context"
	"fmt"

	"github.com/g-villarinho/xp-life-api/models"
	"github.com/g-villarinho/xp-life-api/persistence"
	"github.com/g-villarinho/xp-life-api/pkgs"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *models.Category) error
	GetCategoriesPagedList(ctx context.Context, storeID string, pag models.CategoryPagination) (*models.PaginatedResponse, error)
	GetCategoryByID(ctx context.Context, ID string) (*models.Category, error)
	DeleteCategory(ctx context.Context, ID string) error
}

type categoryRepository struct {
	di   *pkgs.Di
	repo persistence.Repository
}

func NewCategoryRepository(di *pkgs.Di) (CategoryRepository, error) {
	repo, err := pkgs.Invoke[persistence.Repository](di)
	if err != nil {
		return nil, err
	}

	return &categoryRepository{
		di:   di,
		repo: repo,
	}, nil
}

func (c *categoryRepository) CreateCategory(ctx context.Context, category *models.Category) error {
	if err := c.repo.Create(ctx, category); err != nil {
		return err
	}

	return nil
}

func (c *categoryRepository) GetCategoriesPagedList(ctx context.Context, storeID string, pag models.CategoryPagination) (*models.PaginatedResponse, error) {
	var categories []models.Category

	opts := []persistence.QueryOption{}

	opts = append(opts, persistence.WithConditions("store_id = ?", storeID))
	opts = append(opts, persistence.WithPreload("Billboard"))

	if pag.Name != nil {
		opts = append(opts, persistence.WithConditions("name LIKE ?", fmt.Sprintf("%%%s%%", *pag.Name)))
	}

	if pag.BillboardID != nil {
		opts = append(opts, persistence.WithConditions("billboard_id = ?", *pag.BillboardID))
	}

	result, err := c.repo.Paginate(ctx, &categories, *pag.Pagination, opts...)
	if err != nil {
		if err == persistence.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return result, nil
}

func (c *categoryRepository) GetCategoryByID(ctx context.Context, ID string) (*models.Category, error) {
	var category models.Category

	err := c.repo.FindOne(ctx, &category, persistence.WithConditions("id = ?", ID), persistence.WithPreload("Billboard"))
	if err != nil {
		if err == persistence.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &category, nil
}

func (c *categoryRepository) DeleteCategory(ctx context.Context, ID string) error {
	if err := c.repo.Delete(ctx, ID, &models.Category{}); err != nil {
		return err
	}

	return nil
}
