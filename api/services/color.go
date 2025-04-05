package services

import (
	"context"
	"fmt"

	"github.com/g-villarinho/flash-buy-api/models"
	"github.com/g-villarinho/flash-buy-api/pkgs"
	"github.com/g-villarinho/flash-buy-api/repositories"
)

type ColorService interface {
	CreateColor(ctx context.Context, userID string, color models.Color) error
	GetColorsPagedList(ctx context.Context, storeID string, pag models.ColorPagination) (*models.PaginatedResponse, error)
	GetColorByID(ctx context.Context, storeID, colorID string) (*models.ColorResponse, error)
	DeleteColor(ctx context.Context, storeID, userID, colorID string) error
}

type colorService struct {
	di *pkgs.Di
	ss StoreService
	cr repositories.ColorRepository
}

func NewColorService(di *pkgs.Di) (ColorService, error) {
	ss, err := pkgs.Invoke[StoreService](di)
	if err != nil {
		return nil, err
	}

	cr, err := pkgs.Invoke[repositories.ColorRepository](di)
	if err != nil {
		return nil, err
	}

	return &colorService{
		di: di,
		ss: ss,
		cr: cr,
	}, nil
}

func (c *colorService) CreateColor(ctx context.Context, userID string, color models.Color) error {
	_, err := c.ss.GetStoreByStoreID(ctx, color.StoreID.String(), userID)
	if err != nil {
		return err
	}

	colorFromHex, err := c.cr.GetColorByHexAndStoreID(ctx, color.Hex, color.StoreID.String())
	if err != nil {
		return fmt.Errorf("get color by hex: %w", err)
	}

	if colorFromHex != nil {
		return models.ErrHexAlreadyExists
	}

	if err := c.cr.CreateColor(ctx, &color); err != nil {
		return err
	}

	return nil
}

func (c *colorService) GetColorsPagedList(ctx context.Context, storeID string, pag models.ColorPagination) (*models.PaginatedResponse, error) {
	result, err := c.cr.GetColorsPagedList(ctx, storeID, pag)
	if err != nil {
		return nil, err
	}

	return &models.PaginatedResponse{
		Data:       toColorResponseList(*result.Data.(*[]models.Color)),
		Total:      result.Total,
		TotalPages: result.TotalPages,
		Page:       result.Page,
		Limit:      result.Limit,
	}, nil
}

func (c *colorService) GetColorByID(ctx context.Context, storeID string, colorID string) (*models.ColorResponse, error) {
	color, err := c.cr.GetColorByID(ctx, colorID)
	if err != nil {
		return nil, err
	}

	if color == nil {
		return nil, models.ErrColorNotFound
	}

	if color.StoreID.String() != storeID {
		return nil, models.ErrColorNotFound
	}

	return color.ToColorResponse(), nil
}

func (c *colorService) DeleteColor(ctx context.Context, storeID, userID, colorID string) error {
	_, err := c.ss.GetStoreByStoreID(ctx, storeID, userID)
	if err != nil {
		return err
	}

	color, err := c.cr.GetColorByID(ctx, colorID)
	if err != nil {
		return fmt.Errorf("get color by id %s: %w", colorID, err)
	}

	if color == nil {
		return models.ErrColorNotFound
	}

	if color.StoreID.String() != storeID {
		return models.ErrColorNotFound
	}

	if err := c.cr.DeleteColor(ctx, colorID); err != nil {
		return fmt.Errorf("delete color: %w", err)
	}

	return nil
}

func toColorResponseList(colors []models.Color) []models.ColorResponse {
	var colorResponses = make([]models.ColorResponse, len(colors))

	for i, color := range colors {
		colorResponses[i] = *color.ToColorResponse()
	}

	return colorResponses
}
