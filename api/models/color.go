package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrColorNotFound    = errors.New("color not found")
	ErrHexAlreadyExists = errors.New("hex already exists")
)

type Color struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey"`
	Name      string         `gorm:"not null"`
	Hex       string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt sql.NullString `gorm:"default:null"`

	StoreID uuid.UUID `gorm:"type:uuid;not null;index"`
	Store   Store     `gorm:"foreignKey:StoreID"`

	Products []Product `gorm:"foreignKey:ColorID"`
}

type ColorPagination struct {
	*Pagination
	Name *string
}

type CreateColorPayload struct {
	Name string `json:"name" binding:"required"`
	Hex  string `json:"hex" binding:"required"`
}

type ColorResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Hex       string    `json:"hex"`
	CreatedAt time.Time `json:"createdAt"`
}

func (p *CreateColorPayload) ToColor(storeID string) (*Color, error) {
	storeUUID, err := uuid.Parse(storeID)
	if err != nil {
		return nil, err
	}

	return &Color{
		ID:        uuid.New(),
		Name:      p.Name,
		Hex:       p.Hex,
		CreatedAt: time.Now(),
		StoreID:   storeUUID,
	}, nil
}

func (c *Color) ToColorResponse() *ColorResponse {
	return &ColorResponse{
		ID:        c.ID,
		Name:      c.Name,
		Hex:       c.Hex,
		CreatedAt: c.CreatedAt,
	}
}

func NewColorPagination(page, limit string, name *string) *ColorPagination {
	return &ColorPagination{
		Pagination: NewPagination(page, limit),
		Name:       name,
	}
}
