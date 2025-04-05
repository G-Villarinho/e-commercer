package services

import (
	"context"
	"log/slog"
	"mime/multipart"
	"time"

	"github.com/g-villarinho/xp-life-api/models"
	"github.com/g-villarinho/xp-life-api/pkgs"
	"github.com/g-villarinho/xp-life-api/repositories"
)

type ProductImageService interface {
	CreateProductImage(ctx context.Context, productID string, images []*multipart.FileHeader)
}

type productImageService struct {
	di *pkgs.Di
	is ImageService
	pr repositories.ProductImageRepository
}

func NewProductImageService(di *pkgs.Di) (ProductImageService, error) {
	is, err := pkgs.Invoke[ImageService](di)
	if err != nil {
		return nil, err
	}

	pr, err := pkgs.Invoke[repositories.ProductImageRepository](di)
	if err != nil {
		return nil, err
	}

	return &productImageService{
		di: di,
		is: is,
		pr: pr,
	}, nil
}

func (p *productImageService) CreateProductImage(ctx context.Context, productID string, images []*multipart.FileHeader) {
	logger := slog.With(
		"service", "product_image",
		"method", "CreateProductImage",
		"productID", productID,
	)

	for _, image := range images {
		go func(img *multipart.FileHeader) {
			fileLogger := logger.With(
				"filename", img.Filename,
				"size", img.Size,
			)

			const maxRetries = 3
			var lastErr error

			for attempt := 1; attempt <= maxRetries; attempt++ {
				uploadCtx := context.Background()

				fileLogger.Debug("attempting image upload", "attempt", attempt)

				imageURL, err := p.is.UploadImage(uploadCtx, img, img.Filename)
				if err != nil {
					lastErr = err
					fileLogger.Warn("upload failed", "error", err, "retry", attempt)
					time.Sleep(time.Duration(attempt) * time.Second)
					continue
				}

				productImage, err := models.NewProductImage(productID, imageURL)
				if err != nil {
					lastErr = err
					fileLogger.Error("failed to create product image", "error", err)
					break
				}

				if err := p.pr.CreateProductImage(uploadCtx, productImage); err != nil {
					lastErr = err
					fileLogger.Error("failed to save product image", "error", err)
					break
				}

				fileLogger.Info("image successfully processed")
				return
			}

			if lastErr != nil {
				fileLogger.Error("final upload failure after retries", "error", lastErr)
			}
		}(image)
	}
}
