package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/g-villarinho/flash-buy-api/models"
	"github.com/g-villarinho/flash-buy-api/pkgs"
	"github.com/g-villarinho/flash-buy-api/services"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

type RegisterHandler interface {
	Register(ectx echo.Context) error
}

type registerHandler struct {
	di *pkgs.Di
	rs services.RegisterService
}

func NewRegisterHandler(di *pkgs.Di) (RegisterHandler, error) {
	rs, err := pkgs.Invoke[services.RegisterService](di)
	if err != nil {
		return nil, err
	}
	return &registerHandler{
		di: di,
		rs: rs,
	}, nil
}

func (r *registerHandler) Register(ectx echo.Context) error {
	logger := slog.With(
		"handler", "register",
		"method", "Register",
	)

	var payload models.CreateUserPayload
	if err := jsoniter.NewDecoder(ectx.Request().Body).Decode(&payload); err != nil {
		logger.Error("bind payload", "error", err)
		return ectx.NoContent(http.StatusBadRequest)
	}

	user := payload.ToUser()
	ssinfo := models.SessionSecurityInfo{
		IP:        ectx.RealIP(),
		UserAgent: ectx.Request().UserAgent(),
	}

	session, err := r.rs.Register(ectx.Request().Context(), user, ssinfo)
	if err != nil {
		if errors.Is(err, models.ErrUserAlreadyExists) {
			logger.Warn("user already exists", "error", err)
			return ectx.NoContent(http.StatusConflict)
		}

		logger.Error("create user", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	SetCookieSession(ectx, *session)
	return ectx.NoContent(http.StatusCreated)
}
