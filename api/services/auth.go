package services

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/g-villarinho/xp-life-api/models"
	"github.com/g-villarinho/xp-life-api/notifications"
	"github.com/g-villarinho/xp-life-api/pkgs"
)

type AuthService interface {
	Login(ctx context.Context, email string, ssi models.SessionSecurityInfo) (*models.Session, error)
	VerifyCode(ctx context.Context, code, token string) (*models.Session, error)
	ResendCode(ctx context.Context, email, token string) error
}

type authService struct {
	di *pkgs.Di
	us UserService
	ss SessionService
	os OTPService
	en notifications.EmailNotification
}

func NewAuthService(di *pkgs.Di) (AuthService, error) {
	userService, err := pkgs.Invoke[UserService](di)
	if err != nil {
		return nil, err
	}

	otpService, err := pkgs.Invoke[OTPService](di)
	if err != nil {
		return nil, err
	}

	emailNotification, err := pkgs.Invoke[notifications.EmailNotification](di)
	if err != nil {
		return nil, err
	}

	sessionService, err := pkgs.Invoke[SessionService](di)
	if err != nil {
		return nil, err
	}

	return &authService{
		di: di,
		us: userService,
		ss: sessionService,
		os: otpService,
		en: emailNotification,
	}, nil
}

func (a *authService) Login(ctx context.Context, email string, ssi models.SessionSecurityInfo) (*models.Session, error) {
	user, err := a.us.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, models.ErrUserNotFound
	}

	session, err := a.ss.CreateSession(ctx, user, ssi)
	if err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}

	code, err := a.os.GeneratOTP(ctx, user.ID.String(), models.UserVerificationFLow, session.Token)
	if err != nil {
		return nil, fmt.Errorf("create otp: %w", err)
	}

	go func() {
		if err := a.en.SendVerificationEmail(ctx, user.Email, code); err != nil {
			slog.Error("send verification email", "error", err)
		}
	}()

	return session, nil
}

func (a *authService) VerifyCode(ctx context.Context, code, token string) (*models.Session, error) {
	if err := a.os.VerifyOTP(ctx, code, token); err != nil {
		return nil, err
	}

	session, err := a.ss.ValidSession(ctx, token)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (a *authService) ResendCode(ctx context.Context, email string, token string) error {
	code, err := a.os.UpdateCode(ctx, token)
	if err != nil {
		return err
	}

	go func() {
		if err := a.en.SendVerificationEmail(ctx, email, code); err != nil {
			slog.Error("send verification email", "error", err)
		}
	}()

	return nil
}
