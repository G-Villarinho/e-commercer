package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrSessionNotFoundOrExpired = errors.New("session not found or expired")
)

type SessionSecurityInfo struct {
	IP        string
	UserAgent string
}

type Session struct {
	ID         uuid.UUID    `gorm:"type:uuid;primaryKey"`
	Email      string       `gorm:"not null"`
	Token      string       `gorm:"not null"`
	ExpiresAt  time.Time    `gorm:"not null"`
	VerifiedAt sql.NullTime `gorm:"default:null"`
	IP         string       `gorm:"not null"`
	UserAgent  string       `gorm:"not null"`
	CreatedAt  time.Time    `gorm:"not null"`
	UpdatedAt  sql.NullTime `gorm:"default:null"`

	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User   User      `gorm:"foreignKey:UserID"`
}

func (s *Session) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}
