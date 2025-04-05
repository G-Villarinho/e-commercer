package repositories

import (
	"context"
	"fmt"

	"github.com/g-villarinho/xp-life-api/models"
	"github.com/g-villarinho/xp-life-api/persistence"
	"github.com/g-villarinho/xp-life-api/pkgs"
)

type BillboardRepository interface {
	CreateBillboard(ctx context.Context, billboard *models.Billboard) error
	GetBillboardsPagedList(ctx context.Context, storeID string, pag models.BillboardPagination) (*models.PaginatedResponse, error)
	DeleteBillboard(ctx context.Context, ID string) error
	GetBillboardByID(ctx context.Context, ID string) (*models.Billboard, error)
	GetAllByStoreID(ctx context.Context, storeID string) ([]models.Billboard, error)
}

type billboardRepository struct {
	di   *pkgs.Di
	repo persistence.Repository
}

func NewBillboardRepository(di *pkgs.Di) (BillboardRepository, error) {
	repo, err := pkgs.Invoke[persistence.Repository](di)
	if err != nil {
		return nil, err
	}

	return &billboardRepository{
		di:   di,
		repo: repo,
	}, nil
}

func (b *billboardRepository) CreateBillboard(ctx context.Context, billboard *models.Billboard) error {
	if err := b.repo.Create(ctx, billboard); err != nil {
		return err
	}

	return nil
}

func (b *billboardRepository) GetBillboardsPagedList(ctx context.Context, storeID string, pag models.BillboardPagination) (*models.PaginatedResponse, error) {
	var billboards []models.Billboard

	opts := []persistence.QueryOption{}

	opts = append(opts, persistence.WithConditions("store_id = ?", storeID))

	if pag.Label != nil {
		opts = append(opts, persistence.WithConditions("label LIKE ?", fmt.Sprintf("%%%s%%", *pag.Label)))
	}

	result, err := b.repo.Paginate(ctx, &billboards, *pag.Pagination, opts...)
	if err != nil {
		if err == persistence.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return result, nil
}

func (b *billboardRepository) DeleteBillboard(ctx context.Context, ID string) error {
	if err := b.repo.Delete(ctx, ID, &models.Billboard{}); err != nil {
		return err
	}

	return nil
}

func (b *billboardRepository) GetBillboardByID(ctx context.Context, ID string) (*models.Billboard, error) {
	var billboard models.Billboard
	if err := b.repo.FindByID(ctx, ID, &billboard); err != nil {
		if err == persistence.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &billboard, nil
}

func (b *billboardRepository) GetAllByStoreID(ctx context.Context, storeID string) ([]models.Billboard, error) {
	var billboards []models.Billboard

	if err := b.repo.FindAll(ctx, &billboards, persistence.WithConditions("store_id = ?", storeID)); err != nil {
		return nil, err
	}

	return billboards, nil
}
