package handlers

import (
	"log/slog"
	"net/http"

	"github.com/g-villarinho/flash-buy-api/models"
	"github.com/g-villarinho/flash-buy-api/pkgs"
	"github.com/g-villarinho/flash-buy-api/services"
	"github.com/g-villarinho/flash-buy-api/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type BillboardHandler interface {
	CreateBillboard(ectx echo.Context) error
	GetBillboards(ectx echo.Context) error
	DeleteBillboard(ectx echo.Context) error
	GetBillboardByID(ectx echo.Context) error
}

type billboardHandler struct {
	di  *pkgs.Di
	rdp pkgs.RequestDataCtx
	bs  services.BillboardService
}

func NewBillboardHandler(di *pkgs.Di) (BillboardHandler, error) {
	ctxData, err := pkgs.Invoke[pkgs.RequestDataCtx](di)
	if err != nil {
		return nil, err
	}

	billboardService, err := pkgs.Invoke[services.BillboardService](di)
	if err != nil {
		return nil, err
	}

	return &billboardHandler{
		di:  di,
		rdp: ctxData,
		bs:  billboardService,
	}, nil
}

func (b *billboardHandler) CreateBillboard(ectx echo.Context) error {
	logger := slog.With(
		"handler", "billboard",
		"method", "CreateBillboard",
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

	userID, ok := b.rdp.GetUserID(ectx.Request().Context())
	if !ok {
		logger.Error("get user id from context")
		DelCookieSession(ectx)
		return ectx.NoContent(http.StatusUnauthorized)
	}

	file, err := ectx.FormFile("image")
	if err != nil {
		logger.Error("error to bind payload", "error", err)
		return ectx.NoContent(http.StatusBadRequest)
	}

	label := ectx.FormValue("label")
	if label == "" {
		logger.Warn("label is empty")
		return ectx.NoContent(http.StatusBadRequest)
	}

	if err := b.bs.CreateBillboard(ectx.Request().Context(), storeID, userID, file, label); err != nil {
		if err == models.ErrStoreNotFound {
			logger.Warn("store not found", "storeID", storeID)
			return ectx.NoContent(http.StatusNotFound)
		}

		if err == models.ErrStoreNotPertenence {
			logger.Warn("store not pertenence to user", "storeID", storeID)
			return ectx.NoContent(http.StatusForbidden)
		}

		logger.Error("create billboard", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.NoContent(http.StatusCreated)
}

func (b *billboardHandler) GetBillboards(ectx echo.Context) error {
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

	pag := models.NewBillboardPagination(
		ectx.QueryParam("page"),
		ectx.QueryParam("limit"),
		utils.GetQueryStringPointer(ectx.QueryParam("label")),
	)

	resp, err := b.bs.GetBillboardsPagedList(ectx.Request().Context(), storeID, *pag)
	if err != nil {
		logger.Error("failed to fetch billboards", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.JSON(http.StatusOK, resp)
}

func (b *billboardHandler) DeleteBillboard(ectx echo.Context) error {
	logger := slog.With(
		"handler", "billboard",
		"method", "DeleteBillboard",
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

	billboardID := ectx.Param("billboardId")
	if billboardID == "" {
		logger.Warn("billboardID is empty")
		return ectx.NoContent(http.StatusBadRequest)
	}

	if _, err := uuid.Parse(billboardID); err != nil {
		logger.Warn("invalid billboardID format", "billboardID", billboardID)
		return ectx.NoContent(http.StatusBadRequest)
	}

	userID, ok := b.rdp.GetUserID(ectx.Request().Context())
	if !ok {
		logger.Error("get user id from context")
		DelCookieSession(ectx)
		return ectx.NoContent(http.StatusUnauthorized)
	}

	if err := b.bs.DeleteBillboard(ectx.Request().Context(), storeID, userID, billboardID); err != nil {
		if err == models.ErrStoreNotFound {
			logger.Warn("store not found", "storeID", storeID)
			return ectx.NoContent(http.StatusNotFound)
		}

		if err == models.ErrStoreNotPertenence {
			logger.Warn("store not pertenence to user", "storeID", storeID)
			return ectx.NoContent(http.StatusForbidden)
		}

		if err == models.ErrBillboardNotFound {
			logger.Warn("billboard not found", "billboardID", billboardID)
			return ectx.NoContent(http.StatusNotFound)
		}

		if err == models.ErrBillboardNotPertenence {
			logger.Warn("billboard not pertenence to store", "billboardID", billboardID)
			return ectx.NoContent(http.StatusForbidden)
		}

		logger.Error("delete billboard", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.NoContent(http.StatusNoContent)
}

func (b *billboardHandler) GetBillboardByID(ectx echo.Context) error {
	logger := slog.With(
		"handler", "billboard",
		"method", "GetBillboardByID",
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

	billboardID := ectx.Param("billboardId")
	if billboardID == "" {
		logger.Warn("billboardID is empty")
		return ectx.NoContent(http.StatusBadRequest)
	}

	resp, err := b.bs.GetBillboardByID(ectx.Request().Context(), storeID, billboardID)
	if err != nil {
		if err == models.ErrBillboardNotFound {
			logger.Warn("billboard not found", "billboardID", billboardID)
			return ectx.NoContent(http.StatusNotFound)
		}

		logger.Error("get billboard by id", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.JSON(http.StatusOK, resp)
}
