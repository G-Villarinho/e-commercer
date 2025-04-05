package repositories

import (
	"context"

	"github.com/g-villarinho/flash-buy-api/models"
	"github.com/g-villarinho/flash-buy-api/persistence"
	"github.com/g-villarinho/flash-buy-api/pkgs"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *models.Product) error
}

type productRepository struct {
	di   *pkgs.Di
	repo persistence.Repository
}

func NewProductRepository(di *pkgs.Di) (ProductRepository, error) {
	repo, err := pkgs.Invoke[persistence.Repository](di)
	if err != nil {
		return nil, err
	}

	return &productRepository{
		di:   di,
		repo: repo,
	}, nil
}

func (p *productRepository) CreateProduct(ctx context.Context, product *models.Product) error {
	if err := p.repo.Create(ctx, product); err != nil {
		return err
	}

	return nil
}
