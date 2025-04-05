package repositories

import (
	"context"

	"github.com/g-villarinho/flash-buy-api/models"
	"github.com/g-villarinho/flash-buy-api/persistence"
	"github.com/g-villarinho/flash-buy-api/pkgs"
)

type OTPRepository interface {
	CreateOTP(ctx context.Context, otp models.OTP) error
	GetOTPByUserIDAndFlow(ctx context.Context, userID string, flow models.OTPFlow) (*models.OTP, error)
	DeleteOTP(ctx context.Context, ID string) error
	GetOTPByVerificationToken(ctx context.Context, token string) (*models.OTP, error)
	UpdateOTP(ctx context.Context, otp models.OTP) error
}

type otpRepository struct {
	di   *pkgs.Di
	repo persistence.Repository
}

func NewOTPRepository(di *pkgs.Di) (OTPRepository, error) {
	repo, err := pkgs.Invoke[persistence.Repository](di)
	if err != nil {
		return nil, err
	}
	return &otpRepository{
		di:   di,
		repo: repo,
	}, nil
}

func (o *otpRepository) CreateOTP(ctx context.Context, otp models.OTP) error {
	if err := o.repo.Create(ctx, otp); err != nil {
		return err
	}

	return nil
}

func (o *otpRepository) GetOTPByUserIDAndFlow(ctx context.Context, userID string, flow models.OTPFlow) (*models.OTP, error) {
	var otp models.OTP
	if err := o.repo.FindOne(ctx, &otp, persistence.WithConditions("user_id = ? AND flow = ?", userID, flow)); err != nil {
		if err == persistence.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &otp, nil
}

func (o *otpRepository) GetOTPByVerificationToken(ctx context.Context, token string) (*models.OTP, error) {
	var otp models.OTP
	if err := o.repo.FindOne(ctx, &otp, persistence.WithConditions("verification_token = ?", token)); err != nil {
		if err == persistence.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &otp, nil
}

func (o *otpRepository) DeleteOTP(ctx context.Context, ID string) error {
	if err := o.repo.Delete(ctx, ID, models.OTP{}); err != nil {
		return err
	}

	return nil
}

func (o *otpRepository) UpdateOTP(ctx context.Context, otp models.OTP) error {
	if err := o.repo.Update(ctx, otp); err != nil {
		return err
	}

	return nil
}
