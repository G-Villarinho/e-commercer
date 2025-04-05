package services

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/g-villarinho/xp-life-api/models"
	"github.com/g-villarinho/xp-life-api/notifications"
	"github.com/g-villarinho/xp-life-api/pkgs"
)

type RegisterService interface {
	Register(ctx context.Context, user *models.User, ssi models.SessionSecurityInfo) (*models.Session, error)
}

type registerService struct {
	di *pkgs.Di
	us UserService
	ss SessionService
	os OTPService
	en notifications.EmailNotification
}

func NewRegisterService(di *pkgs.Di) (RegisterService, error) {
	userService, err := pkgs.Invoke[UserService](di)
	if err != nil {
		return nil, fmt.Errorf("invoke user service: %w", err)
	}

	otpService, err := pkgs.Invoke[OTPService](di)
	if err != nil {
		return nil, fmt.Errorf("invoke otp service: %w", err)
	}

	emailNotification, err := pkgs.Invoke[notifications.EmailNotification](di)
	if err != nil {
		return nil, fmt.Errorf("invoke email notification: %w", err)
	}

	sessionService, err := pkgs.Invoke[SessionService](di)
	if err != nil {
		return nil, fmt.Errorf("invoke session service: %w", err)
	}

	return &registerService{
		di: di,
		us: userService,
		ss: sessionService,
		os: otpService,
		en: emailNotification,
	}, nil
}

func (r *registerService) Register(ctx context.Context, user *models.User, ssi models.SessionSecurityInfo) (*models.Session, error) {
	if err := r.us.CreateUser(ctx, *user); err != nil {
		return nil, err
	}

	session, err := r.ss.CreateSession(ctx, user, ssi)
	if err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}

	code, err := r.os.GeneratOTP(ctx, user.ID.String(), models.UserVerificationFLow, session.Token)
	if err != nil {
		return nil, err
	}

	go func() {
		if err := r.en.SendVerificationEmail(ctx, user.Email, code); err != nil {
			slog.Error("send verification email", "error", err)
		}
	}()

	return session, nil
}
