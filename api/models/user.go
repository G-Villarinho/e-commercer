package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type User struct {
	ID        uuid.UUID    `gorm:"type:uuid;primaryKey"`
	Name      string       `gorm:"not null"`
	Username  string       `gorm:"not null;unique"`
	Email     string       `gorm:"not null;unique"`
	CreatedAt time.Time    `gorm:"not null"`
	UpdatedAt sql.NullTime `gorm:"default:null"`

	Sessions []Session `gorm:"foreignKey:UserID"`
	OTPs     []OTP     `gorm:"foreignKey:UserID"`
	Stores   []Store   `gorm:"foreignKey:UserID"`
}

type CreateUserPayload struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func (p *CreateUserPayload) ToUser() *User {
	return &User{
		ID:        uuid.New(),
		Name:      p.Name,
		Username:  p.Email,
		Email:     p.Email,
		CreatedAt: time.Now(),
	}
}
