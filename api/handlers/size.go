package handlers

import (
	"log/slog"
	"net/http"

	"github.com/g-villarinho/xp-life-api/models"
	"github.com/g-villarinho/xp-life-api/pkgs"
	"github.com/g-villarinho/xp-life-api/services"
	"github.com/g-villarinho/xp-life-api/utils"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

type SizeHandler interface {
	CreateSize(ectx echo.Context) error
	GetSizesPagedList(ectx echo.Context) error
	DeleteSize(ectx echo.Context) error
	GetSizeByID(ectx echo.Context) error
}

type sizeHandler struct {
	di  *pkgs.Di
	rdp pkgs.RequestDataCtx
	ss  services.SizeService
}

func NewSizeHandler(di *pkgs.Di) (SizeHandler, error) {
	ctxData, err := pkgs.Invoke[pkgs.RequestDataCtx](di)
	if err != nil {
		return nil, err
	}

	sizeService, err := pkgs.Invoke[services.SizeService](di)
	if err != nil {
		return nil, err
	}

	return &sizeHandler{
		di:  di,
		rdp: ctxData,
		ss:  sizeService,
	}, nil
}

func (s *sizeHandler) CreateSize(ectx echo.Context) error {
	logger := slog.With(
		"handler", "size",
		"method", "CreateSize",
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

	var payload models.CreateSizePayload
	if err := jsoniter.NewDecoder(ectx.Request().Body).Decode(&payload); err != nil {
		logger.Error("error to bind payload", "error", err)
		return ectx.NoContent(http.StatusBadRequest)
	}

	userID, ok := s.rdp.GetUserID(ectx.Request().Context())
	if !ok {
		logger.Warn("user id not found in context")
		return ectx.NoContent(http.StatusUnauthorized)
	}

	size, err := payload.ToSize(storeID)
	if err != nil {
		logger.Error("error to convert payload to size", "error", err)
		return ectx.NoContent(http.StatusBadRequest)
	}

	if err := s.ss.CreateSize(ectx.Request().Context(), userID, *size); err != nil {
		if err == models.ErrStoreNotFound {
			logger.Warn("store not found", "storeID", storeID)
			return ectx.NoContent(http.StatusNotFound)
		}

		if err == models.ErrStoreNotPertenence {
			logger.Warn("store not pertenence to user", "storeID", storeID)
			return ectx.NoContent(http.StatusForbidden)
		}

		logger.Error("error to create size", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.NoContent(http.StatusCreated)
}

func (s *sizeHandler) GetSizesPagedList(ectx echo.Context) error {
	logger := slog.With(
		"handler", "size",
		"method", "GetSizesPagedList",
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

	pag := models.NewSizePagination(
		ectx.QueryParam("page"),
		ectx.QueryParam("limit"),
		utils.GetQueryStringPointer(ectx.QueryParam("name")),
	)

	resp, err := s.ss.GetSizesPagedList(ectx.Request().Context(), storeID, *pag)
	if err != nil {
		logger.Error("get sizes paged list", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.JSON(http.StatusOK, resp)
}

func (s *sizeHandler) GetSizeByID(ectx echo.Context) error {
	logger := slog.With(
		"handler", "size",
		"method", "GetSizeByID",
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

	sizeID := ectx.Param("sizeId")
	if sizeID == "" {
		logger.Warn("sizeID is empty")
		return ectx.NoContent(http.StatusBadRequest)
	}

	if _, err := uuid.Parse(sizeID); err != nil {
		logger.Warn("invalid sizeID format", "sizeID", sizeID)
		return ectx.NoContent(http.StatusBadRequest)
	}

	resp, err := s.ss.GetSizeByID(ectx.Request().Context(), sizeID, storeID)
	if err != nil {
		if err == models.ErrSizeNotFound {
			logger.Warn("size not found", "sizeID", sizeID)
			return ectx.NoContent(http.StatusNotFound)
		}

		logger.Error("get size by id", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.JSON(http.StatusOK, resp)
}

func (s *sizeHandler) DeleteSize(ectx echo.Context) error {
	logger := slog.With(
		"handler", "size",
		"method", "DeleteSize",
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

	sizeID := ectx.Param("sizeId")
	if sizeID == "" {
		logger.Warn("sizeID is empty")
		return ectx.NoContent(http.StatusBadRequest)
	}

	if _, err := uuid.Parse(sizeID); err != nil {
		logger.Warn("invalid sizeID format", "sizeID", sizeID)
		return ectx.NoContent(http.StatusBadRequest)
	}

	userID, ok := s.rdp.GetUserID(ectx.Request().Context())
	if !ok {
		logger.Warn("user id not found in context")
		return ectx.NoContent(http.StatusUnauthorized)
	}

	if err := s.ss.DeleteSize(ectx.Request().Context(), userID, storeID, sizeID); err != nil {
		if err == models.ErrSizeNotFound {
			logger.Warn("size not found", "sizeID", sizeID)
			return ectx.NoContent(http.StatusNotFound)
		}

		if err == models.ErrStoreNotFound {
			logger.Warn("store not found", "storeID", storeID)
			return ectx.NoContent(http.StatusNotFound)
		}

		if err == models.ErrStoreNotPertenence {
			logger.Warn("store not pertenence to user", "storeID", storeID)
			return ectx.NoContent(http.StatusForbidden)
		}

		logger.Error("error to delete size", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.NoContent(http.StatusOK)
}
