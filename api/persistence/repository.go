package persistence

import (
	"context"
	"errors"

	"github.com/g-villarinho/xp-life-api/models"
	"gorm.io/gorm"
)

var ErrRecordNotFound = errors.New("record not found")

type QueryOption func(*gorm.DB) *gorm.DB

type Repository interface {
	Create(ctx context.Context, entity any) error
	FindByID(ctx context.Context, id string, out any) error
	Update(ctx context.Context, entity any) error
	Delete(ctx context.Context, id string, model any) error
	FindAll(ctx context.Context, out any, opts ...QueryOption) error
	FindOne(ctx context.Context, out any, opts ...QueryOption) error
	Paginate(ctx context.Context, out any, pagination models.Pagination, opts ...QueryOption) (*models.PaginatedResponse, error)
}
