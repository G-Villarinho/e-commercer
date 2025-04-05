package handlers

import (
	"log/slog"
	"net/http"

	"github.com/g-villarinho/xp-life-api/models"
	"github.com/g-villarinho/xp-life-api/pkgs"
	"github.com/g-villarinho/xp-life-api/services"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	Login(ectx echo.Context) error
	VerifyCode(ectx echo.Context) error
	ResendCode(ectx echo.Context) error
	CheckCode(ectx echo.Context) error
}

type authHandler struct {
	di  *pkgs.Di
	as  services.AuthService
	rdp pkgs.RequestDataCtx
}

func NewAuthHandler(di *pkgs.Di) (AuthHandler, error) {
	as, err := pkgs.Invoke[services.AuthService](di)
	if err != nil {
		return nil, err
	}

	ctxData, err := pkgs.Invoke[pkgs.RequestDataCtx](di)
	if err != nil {
		return nil, err
	}

	return &authHandler{
		di:  di,
		as:  as,
		rdp: ctxData,
	}, nil
}

func (a *authHandler) Login(ectx echo.Context) error {
	logger := slog.With(
		"handler", "auth",
		"method", "Login",
	)

	var payload models.LoginPayload
	if err := jsoniter.NewDecoder(ectx.Request().Body).Decode(&payload); err != nil {
		logger.Error("error to bind payload", "error", err)
		return ectx.NoContent(http.StatusBadRequest)
	}

	ssi := models.SessionSecurityInfo{
		IP:        ectx.RealIP(),
		UserAgent: ectx.Request().UserAgent(),
	}

	session, err := a.as.Login(ectx.Request().Context(), payload.Email, ssi)
	if err != nil {
		if err == models.ErrUserNotFound {
			logger.Warn("user not found", "error", err)
			return ectx.NoContent(http.StatusNotFound)
		}

		logger.Error("error to login", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	SetCookieSession(ectx, *session)
	return ectx.NoContent(http.StatusOK)
}

func (a *authHandler) VerifyCode(ectx echo.Context) error {
	logger := slog.With(
		"handler", "auth",
		"method", "VerifyCode",
	)

	var payload models.VerifyOTPPayload
	if err := jsoniter.NewDecoder(ectx.Request().Body).Decode(&payload); err != nil {
		logger.Error("error to bind payload", "error", err)
		return ectx.NoContent(http.StatusBadRequest)
	}

	userToken, ok := a.rdp.GetToken(ectx.Request().Context())
	if !ok {
		logger.Error("failed to get token from context")
		return ectx.NoContent(http.StatusBadRequest)
	}

	session, err := a.as.VerifyCode(ectx.Request().Context(), payload.Code, userToken)
	if err != nil {
		if err == models.ErrOTPNotFound {
			logger.Warn("otp not found", "error", err)
			DelCookieSession(ectx)
			return ectx.NoContent(http.StatusForbidden)
		}

		if err == models.ErrOTPInvalid {
			logger.Warn("otp invalid", "error", err)
			return ectx.NoContent(http.StatusBadRequest)
		}

		if err == models.ErrSessionNotFoundOrExpired {
			logger.Warn("session not found or expired", "error", err)
			DelCookieSession(ectx)
			return ectx.NoContent(http.StatusForbidden)
		}

		if err == models.ErrOTPExpired {
			logger.Warn("otp expired", "error", err)
			DelCookieSession(ectx)
			return ectx.NoContent(http.StatusForbidden)
		}

		logger.Error("failed to verify code", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	SetCookieSession(ectx, *session)
	return ectx.NoContent(http.StatusOK)
}

func (a *authHandler) ResendCode(ectx echo.Context) error {
	logger := slog.With(
		"handler", "auth",
		"method", "ResendCode",
	)

	userToken, ok := a.rdp.GetToken(ectx.Request().Context())
	if !ok {
		logger.Error("failed to get token from context")
		return ectx.NoContent(http.StatusBadRequest)
	}

	email, ok := a.rdp.GetEmail(ectx.Request().Context())
	if !ok {
		logger.Error("failed to get email from context")
		return ectx.NoContent(http.StatusBadRequest)
	}

	if err := a.as.ResendCode(ectx.Request().Context(), email, userToken); err != nil {
		if err == models.ErrOTPNotFound {
			logger.Warn("otp not found", "error", err)
			DelCookieSession(ectx)
			return ectx.NoContent(http.StatusForbidden)
		}

		logger.Error("failed to resend code", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.NoContent(http.StatusOK)
}

func (a *authHandler) CheckCode(ectx echo.Context) error {
	return ectx.NoContent(http.StatusOK)
}
