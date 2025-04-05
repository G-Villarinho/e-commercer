package services

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/g-villarinho/flash-buy-api/clients"
	"github.com/g-villarinho/flash-buy-api/pkgs"
	"github.com/g-villarinho/flash-buy-api/utils"
)

type ImageService interface {
	UploadImage(ctx context.Context, file *multipart.FileHeader, filename string) (string, error)
}

type imageService struct {
	di *pkgs.Di
	cc clients.CloudflareClient
}

func NewImageService(di *pkgs.Di) (ImageService, error) {
	cc, err := pkgs.Invoke[clients.CloudflareClient](di)
	if err != nil {
		return nil, err
	}
	return &imageService{
		di: di,
		cc: cc,
	}, nil
}

func (i *imageService) UploadImage(ctx context.Context, file *multipart.FileHeader, filename string) (string, error) {
	imageBytes, err := utils.ConvertImageToBytes(file)
	if err != nil {
		return "", fmt.Errorf("convert image to bytes: %w", err)
	}

	imageURL, err := i.cc.UploadImage(ctx, imageBytes, filename)
	if err != nil {
		return "", fmt.Errorf("upload image: %w", err)
	}

	return imageURL, nil
}
