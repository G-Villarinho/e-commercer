package handlers

import (
	"log/slog"
	"net/http"

	"github.com/g-villarinho/flash-buy-api/models"
	"github.com/g-villarinho/flash-buy-api/pkgs"
	"github.com/g-villarinho/flash-buy-api/services"
	"github.com/g-villarinho/flash-buy-api/utils"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

type ColorHandler interface {
	CreateColor(ectx echo.Context) error
	GetColors(ectx echo.Context) error
	DeleteColor(ectx echo.Context) error
	GetColorByID(ectx echo.Context) error
}

type colorHandler struct {
	di  *pkgs.Di
	rdp pkgs.RequestDataCtx
	cs  services.ColorService
}

func NewColorHandler(di *pkgs.Di) (ColorHandler, error) {
	ctxData, err := pkgs.Invoke[pkgs.RequestDataCtx](di)
	if err != nil {
		return nil, err
	}

	cs, err := pkgs.Invoke[services.ColorService](di)
	if err != nil {
		return nil, err
	}

	return &colorHandler{
		di:  di,
		cs:  cs,
		rdp: ctxData,
	}, nil
}

func (c *colorHandler) CreateColor(ectx echo.Context) error {
	logger := slog.With(
		"handler", "color",
		"method", "CreateColor",
	)

	storeID := ectx.Param("storeId")
	if storeID == "" {
		logger.Warn("storeID is empty")
		return ectx.NoContent(http.StatusBadRequest)
	}

	if _, err := uuid.Parse(storeID); err != nil {
		logger.Warn("invalid storeID format", "storeID", storeID)
		return ectx.NoContent(http.StatusBadRequest)
	}

	var payload models.CreateColorPayload
	if err := jsoniter.NewDecoder(ectx.Request().Body).Decode(&payload); err != nil {
		logger.Error("error to bind payload", "error", err)
		return ectx.NoContent(http.StatusBadRequest)
	}

	color, err := payload.ToColor(storeID)
	if err != nil {
		logger.Error("error to convert payload to color", "error", err)
		return ectx.NoContent(http.StatusBadRequest)
	}

	userID, ok := c.rdp.GetUserID(ectx.Request().Context())
	if !ok {
		logger.Warn("user id not found in context")
		return ectx.NoContent(http.StatusUnauthorized)
	}

	if err := c.cs.CreateColor(ectx.Request().Context(), userID, *color); err != nil {
		if err == models.ErrStoreNotFound {
			logger.Warn("store not found", "error", err)
			return ectx.NoContent(http.StatusNotFound)
		}

		if err == models.ErrStoreNotPertenence {
			logger.Warn("store not pertenence to user", "error", err)
			return ectx.NoContent(http.StatusForbidden)
		}

		if err == models.ErrHexAlreadyExists {
			logger.Warn("hex already exists", "error", err)
			return ectx.NoContent(http.StatusConflict)
		}

		logger.Error("error to create color", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.JSON(http.StatusOK, color)
}

func (c *colorHandler) GetColors(ectx echo.Context) error {
	logger := slog.With(
		"handler", "billboard",
		"method", "GetBillboards",
	)

	storeID := ectx.Param("storeId")
	if storeID == "" {
		logger.Warn("storeID is empty")
		return ectx.NoContent(http.StatusBadRequest)
	}

	if _, err := uuid.Parse(storeID); err != nil {
		logger.Warn("invalid storeID format", "storeID", storeID)
		return ectx.NoContent(http.StatusBadRequest)
	}

	pag := models.NewColorPagination(
		ectx.QueryParam("page"),
		ectx.QueryParam("limit"),
		utils.GetQueryStringPointer(ectx.QueryParam("name")),
	)

	resp, err := c.cs.GetColorsPagedList(ectx.Request().Context(), storeID, *pag)
	if err != nil {
		logger.Error("failed to fetch billboards", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.JSON(http.StatusOK, resp)
}

func (c *colorHandler) GetColorByID(ectx echo.Context) error {
	logger := slog.With(
		"handler", "billboard",
		"method", "GetBillboards",
	)

	storeID := ectx.Param("storeId")
	if storeID == "" {
		logger.Warn("storeID is empty")
		return ectx.NoContent(http.StatusBadRequest)
	}

	if _, err := uuid.Parse(storeID); err != nil {
		logger.Warn("invalid storeID format", "storeID", storeID)
		return ectx.NoContent(http.StatusBadRequest)
	}

	colorID := ectx.Param("colorId")
	if colorID == "" {
		logger.Warn("colorID is empty")
		return ectx.NoContent(http.StatusBadRequest)
	}

	if _, err := uuid.Parse(colorID); err != nil {
		logger.Warn("invalid colorID format", "colorID", colorID)
		return ectx.NoContent(http.StatusBadRequest)
	}

	resp, err := c.cs.GetColorByID(ectx.Request().Context(), storeID, colorID)
	if err != nil {
		if err == models.ErrColorNotFound {
			logger.Warn("color not found", "error", err)
			return ectx.NoContent(http.StatusNotFound)
		}

		logger.Error("failed to fetch color", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.JSON(http.StatusOK, resp)
}

func (c *colorHandler) DeleteColor(ectx echo.Context) error {
	logger := slog.With(
		"handler", "billboard",
		"method", "GetBillboards",
	)

	storeID := ectx.Param("storeId")
	if storeID == "" {
		logger.Warn("storeID is empty")
		return ectx.NoContent(http.StatusBadRequest)
	}

	if _, err := uuid.Parse(storeID); err != nil {
		logger.Warn("invalid storeID format", "storeID", storeID)
		return ectx.NoContent(http.StatusBadRequest)
	}

	colorID := ectx.Param("colorId")
	if colorID == "" {
		logger.Warn("colorID is empty")
		return ectx.NoContent(http.StatusBadRequest)
	}

	if _, err := uuid.Parse(colorID); err != nil {
		logger.Warn("invalid colorID format", "colorID", colorID)
		return ectx.NoContent(http.StatusBadRequest)
	}

	userID, ok := c.rdp.GetUserID(ectx.Request().Context())
	if !ok {
		logger.Warn("user id not found in context")
		return ectx.NoContent(http.StatusUnauthorized)
	}

	if err := c.cs.DeleteColor(ectx.Request().Context(), storeID, userID, colorID); err != nil {
		if err == models.ErrStoreNotFound {
			logger.Warn("store not found", "error", err)
			return ectx.NoContent(http.StatusNotFound)
		}

		if err == models.ErrStoreNotPertenence {
			logger.Warn("store not pertenence to user", "error", err)
			return ectx.NoContent(http.StatusForbidden)
		}

		if err == models.ErrColorNotFound {
			logger.Warn("color not found", "error", err)
			return ectx.NoContent(http.StatusNotFound)
		}

		logger.Error("failed to delete color", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.NoContent(http.StatusOK)
}
