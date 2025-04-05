package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrSizeNotFound = errors.New("size not found")
)

type Size struct {
	ID        uuid.UUID    `gorm:"type:uuid;primaryKey"`
	Name      string       `gorm:"not null"`
	Value     string       `gorm:"not null"`
	CreatedAt time.Time    `gorm:"not null"`
	UpdatedAt sql.NullTime `gorm:"default:null"`

	StoreID uuid.UUID `gorm:"type:uuid;not null;index"`
	Store   Store     `gorm:"foreignKey:StoreID"`

	Products []Product `gorm:"foreignKey:SizeID"`
}

type SizePagination struct {
	*Pagination
	Name *string
}

type CreateSizePayload struct {
	Name  string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type SizeResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
}

func (p *CreateSizePayload) ToSize(storeID string) (*Size, error) {
	storeUUID, err := uuid.Parse(storeID)
	if err != nil {
		return nil, err
	}

	return &Size{
		ID:        uuid.New(),
		Name:      p.Name,
		Value:     p.Value,
		CreatedAt: time.Now(),
		StoreID:   storeUUID,
	}, nil
}

func (s *Size) ToSizeResponse() *SizeResponse {
	return &SizeResponse{
		ID:        s.ID,
		Name:      s.Name,
		Value:     s.Value,
		CreatedAt: s.CreatedAt,
	}
}

func NewSizePagination(page, limit string, name *string) *SizePagination {
	pagination := NewPagination(page, limit)
	return &SizePagination{
		Pagination: pagination,
		Name:       name,
	}
}
