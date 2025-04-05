package services

import (
	"context"
	"fmt"

	"github.com/g-villarinho/xp-life-api/models"
	"github.com/g-villarinho/xp-life-api/pkgs"
	"github.com/g-villarinho/xp-life-api/repositories"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, userID string, category models.Category) error
	GetCategoriesPagedList(ctx context.Context, storeID string, pag models.CategoryPagination) (*models.PaginatedResponse, error)
	GetCategoryByID(ctx context.Context, storeID, categoryID string) (*models.CategoryResponse, error)
	DeleteCategory(ctx context.Context, userID, storeID, categoryID string) error
}

type categoryService struct {
	di *pkgs.Di
	ss StoreService
	bs BillboardService
	cr repositories.CategoryRepository
}

func NewCategoryService(di *pkgs.Di) (CategoryService, error) {
	ss, err := pkgs.Invoke[StoreService](di)
	if err != nil {
		return nil, err
	}

	bs, err := pkgs.Invoke[BillboardService](di)
	if err != nil {
		return nil, err
	}

	cr, err := pkgs.Invoke[repositories.CategoryRepository](di)
	if err != nil {
		return nil, err
	}

	return &categoryService{
		di: di,
		ss: ss,
		bs: bs,
		cr: cr,
	}, nil
}

func (c *categoryService) CreateCategory(ctx context.Context, userID string, category models.Category) error {
	_, err := c.ss.GetStoreByStoreID(ctx, category.StoreID.String(), userID)
	if err != nil {
		return err
	}

	_, err = c.bs.GetBillboardByID(ctx, category.StoreID.String(), category.BillboardID.String())
	if err != nil {
		return err
	}

	if err := c.cr.CreateCategory(ctx, &category); err != nil {
		return err
	}

	return nil
}

func (c *categoryService) GetCategoriesPagedList(ctx context.Context, storeID string, pag models.CategoryPagination) (*models.PaginatedResponse, error) {
	result, err := c.cr.GetCategoriesPagedList(ctx, storeID, pag)
	if err != nil {
		return nil, err
	}

	return &models.PaginatedResponse{
		Data:       toCaretegoryResponseList(*result.Data.(*[]models.Category)),
		Total:      result.Total,
		TotalPages: result.TotalPages,
		Page:       result.Page,
		Limit:      result.Limit,
	}, nil
}

func (c *categoryService) GetCategoryByID(ctx context.Context, storeID string, categoryID string) (*models.CategoryResponse, error) {
	category, err := c.cr.GetCategoryByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, models.ErrCategoryNotFound
	}

	if category.StoreID.String() != storeID {
		return nil, models.ErrCategoryNotFound
	}

	return category.ToCategoryResponse(), nil
}

func (c *categoryService) DeleteCategory(ctx context.Context, userID, storeID string, categoryID string) error {
	store, err := c.ss.GetStoreByStoreID(ctx, storeID, userID)
	if err != nil {
		return err
	}

	category, err := c.cr.GetCategoryByID(ctx, categoryID)
	if err != nil {
		return fmt.Errorf("get category by id %s: %w", categoryID, err)
	}

	if category == nil {
		return models.ErrCategoryNotFound
	}

	if store.ID.String() != storeID {
		return models.ErrCategoryNotFound
	}

	if err := c.cr.DeleteCategory(ctx, categoryID); err != nil {
		return fmt.Errorf("delete category: %w", err)
	}

	return nil
}

func toCaretegoryResponseList(categories []models.Category) []models.CategoryResponse {
	responses := make([]models.CategoryResponse, len(categories))
	for i, category := range categories {
		responses[i] = *category.ToCategoryResponse()
	}
	return responses
}
