package persistence

import (
	"context"
	"errors"
	"fmt"

	"github.com/g-villarinho/flash-buy-api/models"
	"github.com/g-villarinho/flash-buy-api/pkgs"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	di *pkgs.Di
	db *gorm.DB
}

func NewPostgresRepository(di *pkgs.Di) (Repository, error) {
	db, err := pkgs.Invoke[*gorm.DB](di)
	if err != nil {
		return nil, fmt.Errorf("invoke gorm.DB: %w", err)
	}

	return &PostgresRepository{
		di: di,
		db: db,
	}, nil
}

func WithConditions(query any, args ...any) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args...)
	}
}

func WithPagination(page, pageSize int) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func WithOrder(order string) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

func WithPreload(association string) QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(association)
	}
}

func (r *PostgresRepository) Create(ctx context.Context, entity any) error {
	return r.db.WithContext(ctx).Create(entity).Error
}

func (r *PostgresRepository) FindByID(ctx context.Context, id string, out any) error {
	err := r.db.WithContext(ctx).Where("id = ?", id).First(out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrRecordNotFound
		}
		return err
	}
	return nil
}

func (r *PostgresRepository) Update(ctx context.Context, entity any) error {
	return r.db.WithContext(ctx).Save(entity).Error
}

func (r *PostgresRepository) Delete(ctx context.Context, id string, model any) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(model)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

func (r *PostgresRepository) FindAll(ctx context.Context, out any, opts ...QueryOption) error {
	db := r.db.WithContext(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	return db.Find(out).Error
}

func (r *PostgresRepository) FindOne(ctx context.Context, out any, opts ...QueryOption) error {
	db := r.db.WithContext(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.First(out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrRecordNotFound
		}
		return err
	}
	return nil
}

func (r *PostgresRepository) Paginate(ctx context.Context, out any, pagination models.Pagination, opts ...QueryOption) (*models.PaginatedResponse, error) {
	db := r.db.WithContext(ctx)

	for _, opt := range opts {
		db = opt(db)
	}

	var total int64
	if err := db.Model(out).Count(&total).Error; err != nil {
		return nil, err
	}

	offset := (pagination.Page - 1) * pagination.Limit

	if err := db.Offset(offset).Limit(pagination.Limit).Find(out).Error; err != nil {
		return nil, err
	}

	totalPages := int((total + int64(pagination.Limit) - 1) / int64(pagination.Limit))

	return &models.PaginatedResponse{
		Data:       out,
		Total:      total,
		TotalPages: totalPages,
		Page:       pagination.Page,
		Limit:      pagination.Limit,
	}, nil
}
