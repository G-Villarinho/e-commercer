package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type ProductImage struct {
	ID        uuid.UUID    `gorm:"type:uuid;primaryKey"`
	ImageURL  string       `gorm:"not null"`
	CreatedAt time.Time    `gorm:"not null"`
	UpdatedAt sql.NullTime `gorm:"default:null"`

	ProductID uuid.UUID `gorm:"type:uuid;not null;index"`
	Product   Product   `gorm:"foreignKey:ProductID"`
}

func NewProductImage(productID, imageURL string) (*ProductImage, error) {
	productUUID, err := uuid.Parse(productID)
	if err != nil {
		return nil, err
	}

	return &ProductImage{
		ID:        uuid.New(),
		ProductID: productUUID,
		ImageURL:  imageURL,
		CreatedAt: time.Now(),
	}, nil
}
