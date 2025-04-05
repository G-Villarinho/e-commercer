package handlers

import (
	"net/http"
	"strings"

	"github.com/g-villarinho/xp-life-api/config"
	"github.com/g-villarinho/xp-life-api/pkgs"
	"github.com/labstack/echo/v4"
)

type EnvironmentHandler interface {
	GetEnvs(c echo.Context) error
}

type environmentHandler struct {
	di *pkgs.Di
}

func NewConfigHandler(di *pkgs.Di) (EnvironmentHandler, error) {
	return &environmentHandler{
		di: di,
	}, nil
}

func (*environmentHandler) GetEnvs(c echo.Context) error {
	if strings.ToLower(config.Env.Env) == "dev" {
		return c.JSON(http.StatusOK, config.Env)
	}

	return c.NoContent(http.StatusNotFound)
}
