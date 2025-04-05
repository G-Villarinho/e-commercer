package services

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/g-villarinho/xp-life-api/models"
	"github.com/g-villarinho/xp-life-api/pkgs"
	"github.com/g-villarinho/xp-life-api/repositories"
)

type BillboardService interface {
	CreateBillboard(ctx context.Context, storeID, userID string, file *multipart.FileHeader, label string) error
	GetBillboardsPagedList(ctx context.Context, storeID string, pag models.BillboardPagination) (*models.PaginatedResponse, error)
	DeleteBillboard(ctx context.Context, storeID, userID, billboardID string) error
	GetBillboardByID(ctx context.Context, storeId, billboardID string) (*models.BillboardResponse, error)
	GetAllByStoreID(ctx context.Context, storeID string) ([]models.BillboardResponse, error)
}

type billboardService struct {
	di *pkgs.Di
	ss StoreService
	is ImageService
	br repositories.BillboardRepository
}

func NewBillboardService(di *pkgs.Di) (BillboardService, error) {
	ss, err := pkgs.Invoke[StoreService](di)
	if err != nil {
		return nil, err
	}

	is, err := pkgs.Invoke[ImageService](di)
	if err != nil {
		return nil, err
	}

	br, err := pkgs.Invoke[repositories.BillboardRepository](di)
	if err != nil {
		return nil, err
	}

	return &billboardService{
		di: di,
		ss: ss,
		is: is,
		br: br,
	}, nil
}

func (b *billboardService) CreateBillboard(ctx context.Context, storeID string, userID string, file *multipart.FileHeader, label string) error {
	store, err := b.ss.GetStoreByStoreID(ctx, storeID, userID)
	if err != nil {
		return err
	}

	imageURL, err := b.is.UploadImage(ctx, file, file.Filename)
	if err != nil {
		return err
	}

	billboard, err := models.NewBillboard(label, store.ID.String(), imageURL)
	if err != nil {
		return fmt.Errorf("new billboard: %w", err)
	}

	if err := b.br.CreateBillboard(ctx, billboard); err != nil {
		return fmt.Errorf("create billboard: %w", err)
	}

	return nil
}

func (b *billboardService) GetBillboardsPagedList(ctx context.Context, storeID string, pag models.BillboardPagination) (*models.PaginatedResponse, error) {
	result, err := b.br.GetBillboardsPagedList(ctx, storeID, pag)
	if err != nil {
		return nil, err
	}

	return &models.PaginatedResponse{
		Data:       toBillboardResponseList(*result.Data.(*[]models.Billboard)),
		Total:      result.Total,
		TotalPages: result.TotalPages,
		Page:       result.Page,
		Limit:      result.Limit,
	}, nil
}

func (b *billboardService) DeleteBillboard(ctx context.Context, storeID string, userID string, billboardID string) error {
	store, err := b.ss.GetStoreByStoreID(ctx, storeID, userID)
	if err != nil {
		return err
	}

	billboard, err := b.br.GetBillboardByID(ctx, billboardID)
	if err != nil {
		return err
	}

	if billboard.StoreID.String() != store.ID.String() {
		return models.ErrBillboardNotPertenence
	}

	if err := b.br.DeleteBillboard(ctx, billboardID); err != nil {
		return err
	}

	return nil
}

func (b *billboardService) GetBillboardByID(ctx context.Context, storeId, billboardID string) (*models.BillboardResponse, error) {
	billboard, err := b.br.GetBillboardByID(ctx, billboardID)
	if err != nil {
		return nil, err
	}

	if billboard == nil {
		return nil, models.ErrBillboardNotFound
	}

	if billboard.StoreID.String() != storeId {
		return nil, models.ErrBillboardNotFound
	}

	billboardResponse := billboard.ToBillboardResponse()

	return &billboardResponse, nil
}

func (b *billboardService) GetAllByStoreID(ctx context.Context, storeID string) ([]models.BillboardResponse, error) {
	billboards, err := b.br.GetAllByStoreID(ctx, storeID)
	if err != nil {
		return nil, err
	}

	return toBillboardResponseList(billboards), nil
}

func toBillboardResponseList(billboards []models.Billboard) []models.BillboardResponse {
	responses := make([]models.BillboardResponse, len(billboards))
	for i, billboard := range billboards {
		responses[i] = billboard.ToBillboardResponse()
	}
	return responses
}
