package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrOTPNotFound = errors.New("otp not found")
	ErrOTPInvalid  = errors.New("otp invalid")
	ErrOTPExpired  = errors.New("otp expired")
)

type OTPFlow string

const (
	UserVerificationFLow OTPFlow = "user_verification"
)

type OTP struct {
	ID                uuid.UUID `gorm:"primaryKey"`
	Code              string    `gorm:"type:varchar(6);not null"`
	Flow              OTPFlow   `gorm:"type:varchar(50);not null"`
	VerificationToken string    `gorm:"not null;unique"`
	ExpiresAt         time.Time `gorm:"not null"`

	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User   User      `gorm:"foreignKey:UserID"`
}

type VerifyOTPPayload struct {
	Code string `json:"code"`
}

func (o *OTP) IsExpired() bool {
	return time.Now().After(o.ExpiresAt)
}

func NewOTP(userID, code string, flow OTPFlow, expiresAt time.Time, token string) (*OTP, error) {
	userIDuuid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	return &OTP{
		ID:                uuid.New(),
		Flow:              flow,
		UserID:            userIDuuid,
		VerificationToken: token,
		Code:              code,
		ExpiresAt:         expiresAt,
	}, nil
}
