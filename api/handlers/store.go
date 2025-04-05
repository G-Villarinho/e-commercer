package handlers

import (
	"log/slog"
	"net/http"

	"github.com/g-villarinho/flash-buy-api/models"
	"github.com/g-villarinho/flash-buy-api/pkgs"
	"github.com/g-villarinho/flash-buy-api/services"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

type StoreHandler interface {
	CreateStore(ectx echo.Context) error
	GetUserFirstStore(ectx echo.Context) error
	GetStoreByID(ectx echo.Context) error
	GetStoresByUserID(ectx echo.Context) error
	UpdateStore(ectx echo.Context) error
	DeleteStore(ectx echo.Context) error
}

type storeHandler struct {
	di  *pkgs.Di
	ss  services.StoreService
	rdp pkgs.RequestDataCtx
}

func NewStoreHandler(di *pkgs.Di) (StoreHandler, error) {
	ss, err := pkgs.Invoke[services.StoreService](di)
	if err != nil {
		return nil, err
	}

	ctxData, err := pkgs.Invoke[pkgs.RequestDataCtx](di)
	if err != nil {
		return nil, err
	}

	return &storeHandler{
		di:  di,
		ss:  ss,
		rdp: ctxData,
	}, nil
}

func (s *storeHandler) CreateStore(ectx echo.Context) error {
	logger := slog.With(
		"handler", "store",
		"method", "CreateStore",
	)

	var payload models.CreateStorePayload
	if err := jsoniter.NewDecoder(ectx.Request().Body).Decode(&payload); err != nil {
		logger.Error("error to bind payload", "error", err)
		return ectx.NoContent(http.StatusBadRequest)
	}

	userID, ok := s.rdp.GetUserID(ectx.Request().Context())
	if !ok {
		logger.Error("get user id from context")
		DelCookieSession(ectx)
		return ectx.NoContent(http.StatusUnauthorized)
	}

	store, err := payload.ToStore(userID)
	if err != nil {
		logger.Error("convert payload to store", "error", err)
		return ectx.NoContent(http.StatusBadRequest)
	}

	response, err := s.ss.CreateStore(ectx.Request().Context(), store)
	if err != nil {
		if err == models.ErrStoreAlreadyExists {
			logger.Warn("store already exists", "error", err)
			return ectx.NoContent(http.StatusConflict)
		}

		logger.Error("create store", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.JSON(http.StatusCreated, response)
}

func (s *storeHandler) GetUserFirstStore(ectx echo.Context) error {
	logger := slog.With(
		"handler", "store",
		"method", "GetUserFirstStore",
	)

	userID, ok := s.rdp.GetUserID(ectx.Request().Context())
	if !ok {
		logger.Error("get user id from context")
		DelCookieSession(ectx)
		return ectx.NoContent(http.StatusUnauthorized)
	}

	response, err := s.ss.GetUserFirstStore(ectx.Request().Context(), userID)
	if err != nil {
		if err == models.ErrStoreNotFound {
			logger.Warn("no store found for user", "userID", userID)

			return ectx.NoContent(http.StatusNotFound)
		}

		logger.Error("error getting first store", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.JSON(http.StatusOK, response)
}

func (s *storeHandler) GetStoreByID(ectx echo.Context) error {
	logger := slog.With(
		"handler", "store",
		"method", "GetStoreByID",
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

	userID, ok := s.rdp.GetUserID(ectx.Request().Context())
	if !ok {
		logger.Error("get user id from context")
		DelCookieSession(ectx)
		return ectx.NoContent(http.StatusUnauthorized)
	}

	response, err := s.ss.GetStoreByStoreID(ectx.Request().Context(), storeID, userID)
	if err != nil {
		if err == models.ErrStoreNotFound {
			logger.Warn("store not found", "storeID", storeID)
			return ectx.NoContent(http.StatusNotFound)
		}

		if err == models.ErrStoreNotPertenence {
			logger.Warn("store not pertenence to user", "storeID", storeID)
			return ectx.NoContent(http.StatusForbidden)
		}

		logger.Error("get store by id", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.JSON(http.StatusOK, response)
}

func (s *storeHandler) GetStoresByUserID(ectx echo.Context) error {
	logger := slog.With(
		"handler", "store",
		"method", "GetStoresByUserID",
	)

	userID, ok := s.rdp.GetUserID(ectx.Request().Context())
	if !ok {
		logger.Error("get user id from context")
		DelCookieSession(ectx)
		return ectx.NoContent(http.StatusUnauthorized)
	}

	response, err := s.ss.GetStoresByUserID(ectx.Request().Context(), userID)
	if err != nil {
		logger.Error("get stores by user id", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.JSON(http.StatusOK, response)
}

func (s *storeHandler) UpdateStore(ectx echo.Context) error {
	logger := slog.With(
		"handler", "store",
		"method", "UpdateStore",
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

	var payload models.UpdateStorePayload
	if err := jsoniter.NewDecoder(ectx.Request().Body).Decode(&payload); err != nil {
		logger.Error("error to bind payload", "error", err)
		return ectx.NoContent(http.StatusBadRequest)
	}

	userID, ok := s.rdp.GetUserID(ectx.Request().Context())
	if !ok {
		logger.Error("get user id from context")
		DelCookieSession(ectx)
		return ectx.NoContent(http.StatusUnauthorized)
	}

	if err := s.ss.UpdateStore(ectx.Request().Context(), userID, storeID, payload.Name); err != nil {
		if err == models.ErrStoreNotFound {
			logger.Warn("store not found", "slug", storeID)
			return ectx.NoContent(http.StatusNotFound)
		}

		if err == models.ErrStoreNotPertenence {
			logger.Warn("store not pertenence to user", "slug", storeID)
			return ectx.NoContent(http.StatusForbidden)
		}

		logger.Error("update store", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.NoContent(http.StatusOK)
}

func (s *storeHandler) DeleteStore(ectx echo.Context) error {
	logger := slog.With(
		"handler", "store",
		"method", "DeleteStore",
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

	userID, ok := s.rdp.GetUserID(ectx.Request().Context())
	if !ok {
		logger.Error("get user id from context")
		DelCookieSession(ectx)
		return ectx.NoContent(http.StatusUnauthorized)
	}

	if err := s.ss.DeleteStore(ectx.Request().Context(), storeID, userID); err != nil {
		if err == models.ErrStoreNotFound {
			logger.Warn("store not found", "slug", storeID)
			return ectx.NoContent(http.StatusNotFound)
		}

		if err == models.ErrStoreNotPertenence {
			logger.Warn("store not pertenence to user", "slug", storeID)
			return ectx.NoContent(http.StatusForbidden)
		}

		logger.Error("delete store", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.NoContent(http.StatusNoContent)
}
