package services

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/g-villarinho/flash-buy-api/models"
	"github.com/g-villarinho/flash-buy-api/pkgs"
	"github.com/g-villarinho/flash-buy-api/repositories"
)

type ProductService interface {
	CreateProduct(ctx context.Context, userID string, product models.Product, images []*multipart.FileHeader) error
}

type productService struct {
	di  *pkgs.Di
	srs StoreService
	szs SizeService
	crs ColorService
	ces CategoryService
	pis ProductImageService
	pr  repositories.ProductRepository
}

func NewProductService(di *pkgs.Di) (ProductService, error) {
	srs, err := pkgs.Invoke[StoreService](di)
	if err != nil {
		return nil, err
	}

	szs, err := pkgs.Invoke[SizeService](di)
	if err != nil {
		return nil, err
	}

	crs, err := pkgs.Invoke[ColorService](di)
	if err != nil {
		return nil, err
	}

	pis, err := pkgs.Invoke[ProductImageService](di)
	if err != nil {
		return nil, err
	}

	ces, err := pkgs.Invoke[CategoryService](di)
	if err != nil {
		return nil, err
	}

	pr, err := pkgs.Invoke[repositories.ProductRepository](di)
	if err != nil {
		return nil, err
	}

	return &productService{
		di:  di,
		srs: srs,
		szs: szs,
		crs: crs,
		ces: ces,
		pis: pis,
		pr:  pr,
	}, nil
}

func (p *productService) CreateProduct(ctx context.Context, userID string, product models.Product, images []*multipart.FileHeader) error {
	_, err := p.srs.GetStoreByStoreID(ctx, product.StoreID.String(), userID)
	if err != nil {
		return err
	}

	_, err = p.szs.GetSizeByID(ctx, product.SizeID.String(), product.StoreID.String())
	if err != nil {
		return err
	}

	_, err = p.crs.GetColorByID(ctx, product.StoreID.String(), product.ColorID.String())
	if err != nil {
		return err
	}

	_, err = p.ces.GetCategoryByID(ctx, product.StoreID.String(), product.CategoryID.String())
	if err != nil {
		return err
	}

	if err := p.pr.CreateProduct(ctx, &product); err != nil {
		return fmt.Errorf("create product: %w", err)
	}

	go p.pis.CreateProductImage(ctx, product.ID.String(), images)

	return nil
}
