package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID           uuid.UUID    `gorm:"type:uuid;primaryKey"`
	Name         string       `gorm:"not null"`
	PriceInCents int64        `gorm:"not null"`
	IsFeatured   bool         `gorm:"not null;default:false"`
	IsArchived   bool         `gorm:"not null;default:false"`
	CreatedAt    time.Time    `gorm:"not null"`
	UpdatedAt    sql.NullTime `gorm:"default:null"`

	StoreID uuid.UUID `gorm:"type:uuid;not null;index"`
	Store   Store     `gorm:"foreignKey:StoreID"`

	CategoryID uuid.UUID `gorm:"type:uuid;not null;index"`
	Category   Category  `gorm:"foreignKey:CategoryID"`

	ColorID uuid.UUID `gorm:"type:uuid;not null;index"`
	Color   Color     `gorm:"foreignKey:ColorID"`

	SizeID uuid.UUID `gorm:"type:uuid;not null;index"`
	Size   Size      `gorm:"foreignKey:SizeID"`

	ProductImages []ProductImage `gorm:"foreignKey:ProductID"`
}

type CreateProductPayload struct {
	Name       string  `form:"name" binding:"required"`
	Price      float64 `form:"price" binding:"required"`
	IsFeatured bool    `form:"isFeatured"`
	IsArchived bool    `form:"isArchived"`
	CategoryID string  `form:"categoryId" binding:"required"`
	ColorID    string  `form:"colorId" binding:"required"`
	SizeID     string  `form:"sizeId" binding:"required"`
}

func (p *CreateProductPayload) ToProduct(storeId string) (*Product, error) {
	storeUUID, err := uuid.Parse(storeId)
	if err != nil {
		return nil, err
	}

	categoryUUID, err := uuid.Parse(p.CategoryID)
	if err != nil {
		return nil, err
	}

	colorUUID, err := uuid.Parse(p.ColorID)
	if err != nil {
		return nil, err
	}

	sizeUUID, err := uuid.Parse(p.SizeID)
	if err != nil {
		return nil, err
	}

	return &Product{
		ID:           uuid.New(),
		Name:         p.Name,
		PriceInCents: int64(p.Price * 100),
		IsFeatured:   p.IsFeatured,
		IsArchived:   p.IsArchived,
		CreatedAt:    time.Now(),
		StoreID:      storeUUID,
		CategoryID:   categoryUUID,
		ColorID:      colorUUID,
		SizeID:       sizeUUID,
	}, nil
}
