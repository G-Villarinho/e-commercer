package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrBillboardNotFound      = errors.New("billboard not found")
	ErrBillboardNotPertenence = errors.New("billboard not pertenence to store")
)

type Billboard struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey"`
	Label     string         `gorm:"not null"`
	ImageURL  sql.NullString `gorm:"default:null"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt sql.NullTime   `gorm:"default:null"`

	StoreID uuid.UUID `gorm:"type:uuid;not null;index"`
	Store   Store     `gorm:"foreignKey:StoreID"`

	Catetegories []Category `gorm:"foreignKey:BillboardID"`
}

type BillboardPagination struct {
	*Pagination
	Label *string
}

type BillboardResponse struct {
	ID        uuid.UUID `json:"id"`
	Label     string    `json:"label"`
	ImageURL  string    `json:"imageUrl"`
	CreatedAt time.Time `json:"createdAt"`
}

type BillboardBasicResponse struct {
	ID    uuid.UUID `json:"id"`
	Label string    `json:"label"`
}

func (b *Billboard) ToBillboardResponse() BillboardResponse {
	return BillboardResponse{
		ID:        b.ID,
		Label:     b.Label,
		ImageURL:  b.ImageURL.String,
		CreatedAt: b.CreatedAt,
	}
}

func (b *Billboard) ToBillboardBasicResponse() BillboardBasicResponse {
	return BillboardBasicResponse{
		ID:    b.ID,
		Label: b.Label,
	}
}

func NewBillboard(label string, storeID string, imageURL string) (*Billboard, error) {
	storeIDuuid, err := uuid.Parse(storeID)
	if err != nil {
		return nil, err
	}

	return &Billboard{
		ID:        uuid.New(),
		Label:     label,
		CreatedAt: time.Now(),
		ImageURL:  sql.NullString{String: imageURL, Valid: imageURL != ""},
		StoreID:   storeIDuuid,
	}, nil
}

func NewBillboardPagination(page, limit string, label *string) *BillboardPagination {
	return &BillboardPagination{
		Pagination: NewPagination(page, limit),
		Label:      label,
	}
}
