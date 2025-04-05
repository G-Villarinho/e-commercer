package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrStoreAlreadyExists = errors.New("store already exists")
	ErrStoreNotFound      = errors.New("store not found")
	ErrStoreNotPertenence = errors.New("store not pertenence to user")
)

type Store struct {
	ID        uuid.UUID    `gorm:"type:uuid;primaryKey"`
	Name      string       `gorm:"not null"`
	CreatedAt time.Time    `gorm:"not null"`
	UpdatedAt sql.NullTime `gorm:"default:null"`

	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User   User      `gorm:"foreignKey:UserID"`

	Billboards []Billboard `gorm:"foreignKey:StoreID"`
	Categories []Category  `gorm:"foreignKey:StoreID"`
	Sizes      []Size      `gorm:"foreignKey:StoreID"`
	Colors     []Color     `gorm:"foreignKey:StoreID"`
	Products   []Product   `gorm:"foreignKey:StoreID"`
}

type CreateStorePayload struct {
	Name string `json:"name" binding:"required"`
}

type UpdateStorePayload struct {
	Name string `json:"name"`
}

type CreateStoreResponse struct {
	StoreId string `json:"storeId"`
}

type StoreResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func (p *CreateStorePayload) ToStore(userID string) (*Store, error) {
	userIDuuid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	return &Store{
		ID:        uuid.New(),
		Name:      p.Name,
		CreatedAt: time.Now(),
		UserID:    userIDuuid,
	}, nil
}

func (s *Store) ToStoreResponse() *StoreResponse {
	return &StoreResponse{
		ID:        s.ID,
		Name:      s.Name,
		CreatedAt: s.CreatedAt,
	}
}
