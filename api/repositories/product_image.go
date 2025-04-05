package repositories

import (
	"context"

	"github.com/g-villarinho/xp-life-api/models"
	"github.com/g-villarinho/xp-life-api/persistence"
	"github.com/g-villarinho/xp-life-api/pkgs"
)

type ProductImageRepository interface {
	CreateProductImage(ctx context.Context, productImage *models.ProductImage) error
}

type productImageRepository struct {
	di   *pkgs.Di
	repo persistence.Repository
}

func NewProductImageRepository(di *pkgs.Di) (ProductImageRepository, error) {
	repo, err := pkgs.Invoke[persistence.Repository](di)
	if err != nil {
		return nil, err
	}

	return &productImageRepository{
		di:   di,
		repo: repo,
	}, nil
}

func (p *productImageRepository) CreateProductImage(ctx context.Context, productImage *models.ProductImage) error {
	if err := p.repo.Create(ctx, productImage); err != nil {
		return err
	}

	return nil
}
