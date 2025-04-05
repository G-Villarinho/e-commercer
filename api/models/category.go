package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
)

type Category struct {
	ID        uuid.UUID    `gorm:"type:uuid;primaryKey"`
	Name      string       `gorm:"not null"`
	CreatedAt time.Time    `gorm:"not null"`
	UpdatedAt sql.NullTime `gorm:"default:null"`

	StoreID uuid.UUID `gorm:"type:uuid;not null;index"`
	Store   Store     `gorm:"foreignKey:StoreID"`

	BillboardID uuid.UUID `gorm:"type:uuid;not null;index"`
	Billboard   Billboard `gorm:"foreignKey:BillboardID"`
}

type CategoryPagination struct {
	*Pagination
	Name        *string
	BillboardID *string
}

type CreateCategoryPayload struct {
	Name        string `json:"name" binding:"required"`
	BillboardID string `json:"billboardId" binding:"required"`
}

type CategoryResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`

	BillboardResponse BillboardBasicResponse `json:"billboard"`
}

func (p *CreateCategoryPayload) ToCategory(storeID string) (*Category, error) {
	storeUUID, err := uuid.Parse(storeID)
	if err != nil {
		return nil, fmt.Errorf("invalid storeID: %w", err)
	}

	billboardUUID, err := uuid.Parse(p.BillboardID)
	if err != nil {
		return nil, fmt.Errorf("invalid billboardID: %w", err)
	}

	return &Category{
		ID:          uuid.New(),
		Name:        p.Name,
		StoreID:     storeUUID,
		BillboardID: billboardUUID,
		CreatedAt:   time.Now(),
	}, nil
}

func (c *Category) ToCategoryResponse() *CategoryResponse {
	return &CategoryResponse{
		ID:                c.ID,
		Name:              c.Name,
		CreatedAt:         c.CreatedAt,
		BillboardResponse: c.Billboard.ToBillboardBasicResponse(),
	}
}

func NewCategoryPagination(page, limit string, name, billboardID *string) *CategoryPagination {
	return &CategoryPagination{
		Pagination:  NewPagination(page, limit),
		Name:        name,
		BillboardID: billboardID,
	}
}
