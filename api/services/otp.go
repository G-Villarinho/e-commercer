package services

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/g-villarinho/flash-buy-api/models"
	"github.com/g-villarinho/flash-buy-api/pkgs"
	"github.com/g-villarinho/flash-buy-api/repositories"
)

const (
	alphanumericCharset = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	alphanumericLength  = 6
	otpExpiration       = 5 * 60 // 5 minutes
)

type OTPService interface {
	GeneratOTP(ctx context.Context, userID string, flow models.OTPFlow, token string) (string, error)
	VerifyOTP(ctx context.Context, code, verificationToken string) error
	UpdateCode(ctx context.Context, verificationToken string) (string, error)
}

type otpService struct {
	di *pkgs.Di
	or repositories.OTPRepository
}

func NewOTPService(di *pkgs.Di) (OTPService, error) {
	or, err := pkgs.Invoke[repositories.OTPRepository](di)
	if err != nil {
		return nil, fmt.Errorf("invoke otp repository: %w", err)
	}
	return &otpService{
		di: di,
		or: or,
	}, nil
}

func (o *otpService) GeneratOTP(ctx context.Context, userID string, flow models.OTPFlow, token string) (string, error) {
	code, err := generateCode()
	if err != nil {
		return "", fmt.Errorf("generate otp code: %w", err)
	}
	expirateAt := time.Now().Add(time.Duration(otpExpiration) * time.Second)

	otp, err := models.NewOTP(userID, code, flow, expirateAt, token)
	if err != nil {
		return "", fmt.Errorf("new otp: %w", err)
	}

	if err := o.or.CreateOTP(ctx, *otp); err != nil {
		return "", fmt.Errorf("create otp code: %w", err)
	}

	return code, nil
}

func (o *otpService) VerifyOTP(ctx context.Context, code, token string) error {
	codeLower := strings.ToUpper(code)

	otp, err := o.or.GetOTPByVerificationToken(ctx, token)
	if err != nil {
		return fmt.Errorf("get otp by verification token: %w", err)
	}

	if otp == nil {
		return models.ErrOTPNotFound
	}

	if otp.IsExpired() {
		return models.ErrOTPExpired
	}

	if otp.Code != codeLower {
		return models.ErrOTPInvalid
	}

	if err := o.or.DeleteOTP(ctx, otp.ID.String()); err != nil {
		return fmt.Errorf("delete otp: %w", err)
	}

	return nil
}

func (o *otpService) UpdateCode(ctx context.Context, verificationToken string) (string, error) {
	code, err := generateCode()
	if err != nil {
		return "", fmt.Errorf("generate otp code: %w", err)
	}

	expirateAt := time.Now().Add(time.Duration(otpExpiration) * time.Second)

	otp, err := o.or.GetOTPByVerificationToken(ctx, verificationToken)
	if err != nil {
		return "", fmt.Errorf("get otp by verification token: %w", err)
	}

	if otp == nil {
		return "", models.ErrOTPNotFound
	}

	otp.Code = code
	otp.ExpiresAt = expirateAt

	if err := o.or.UpdateOTP(ctx, *otp); err != nil {
		return "", fmt.Errorf("update otp code: %w", err)
	}

	return code, nil
}

func generateCode() (string, error) {
	otp := make([]byte, alphanumericLength)
	for i := range otp {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphanumericCharset))))
		if err != nil {
			return "", err
		}
		otp[i] = alphanumericCharset[n.Int64()]
	}

	return string(otp), nil
}
