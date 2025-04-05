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

type CategoryHandler interface {
	CreateCategory(ectx echo.Context) error
	GetCategoriesPagedList(ectx echo.Context) error
	DeleteCategory(ectx echo.Context) error
	GetCategoryByID(ectx echo.Context) error
}

type categoryHandler struct {
	di  *pkgs.Di
	rdp pkgs.RequestDataCtx
	cs  services.CategoryService
}

func NewCategoryHandler(di *pkgs.Di) (CategoryHandler, error) {
	ctxData, err := pkgs.Invoke[pkgs.RequestDataCtx](di)
	if err != nil {
		return nil, err
	}

	cs, err := pkgs.Invoke[services.CategoryService](di)
	if err != nil {
		return nil, err
	}

	return &categoryHandler{
		di:  di,
		cs:  cs,
		rdp: ctxData,
	}, nil
}

func (c *categoryHandler) CreateCategory(ectx echo.Context) error {
	logger := slog.With(
		"handler", "category",
		"method", "CreateCategory",
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

	var payload models.CreateCategoryPayload
	if err := jsoniter.NewDecoder(ectx.Request().Body).Decode(&payload); err != nil {
		logger.Error("error to bind payload", "error", err)
		return ectx.NoContent(http.StatusBadRequest)
	}

	userID, ok := c.rdp.GetUserID(ectx.Request().Context())
	if !ok {
		logger.Error("get user id from context")
		DelCookieSession(ectx)
		return ectx.NoContent(http.StatusUnauthorized)
	}

	category, err := payload.ToCategory(storeID)
	if err != nil {
		logger.Error("error to convert payload to category", "error", err)
		return ectx.NoContent(http.StatusBadRequest)
	}

	if err := c.cs.CreateCategory(ectx.Request().Context(), userID, *category); err != nil {
		if err == models.ErrStoreNotFound {
			logger.Warn("store not found", "error", err)
			return ectx.NoContent(http.StatusNotFound)
		}

		if err == models.ErrBillboardNotFound {
			logger.Warn("billboard not found", "error", err)
			return ectx.NoContent(http.StatusNotFound)
		}

		if err == models.ErrStoreNotPertenence {
			logger.Warn("store not pertenence to user", "error", err)
			return ectx.NoContent(http.StatusForbidden)
		}

		logger.Error("error to create category", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.NoContent(http.StatusCreated)
}

func (c *categoryHandler) GetCategoriesPagedList(ectx echo.Context) error {
	logger := slog.With(
		"handler", "category",
		"method", "GetCategoriesPagedList",
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

	pag := models.NewCategoryPagination(
		ectx.QueryParam("page"),
		ectx.QueryParam("limit"),
		utils.GetQueryStringPointer(ectx.QueryParam("name")),
		utils.GetQueryStringPointer(ectx.QueryParam("billboardId")),
	)

	resp, err := c.cs.GetCategoriesPagedList(ectx.Request().Context(), storeID, *pag)
	if err != nil {
		if err == models.ErrStoreNotFound {
			logger.Warn("store not found", "error", err)
			return ectx.NoContent(http.StatusNotFound)
		}

		if err == models.ErrStoreNotPertenence {
			logger.Warn("store not pertenence to user", "error", err)
			return ectx.NoContent(http.StatusForbidden)
		}

		logger.Error("error to get categories paged list", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.JSON(http.StatusOK, resp)
}

func (c *categoryHandler) GetCategoryByID(ectx echo.Context) error {
	logger := slog.With(
		"handler", "category",
		"method", "GetCategoryByID",
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

	categoryID := ectx.Param("categoryId")
	if categoryID == "" {
		logger.Warn("categoryID is empty")
		return ectx.NoContent(http.StatusBadRequest)
	}

	if _, err := uuid.Parse(categoryID); err != nil {
		logger.Warn("invalid categoryID format", "categoryID", categoryID)
		return ectx.NoContent(http.StatusBadRequest)
	}

	resp, err := c.cs.GetCategoryByID(ectx.Request().Context(), storeID, categoryID)
	if err != nil {
		if err == models.ErrCategoryNotFound {
			logger.Warn("category not found", "error", err)
			return ectx.NoContent(http.StatusNotFound)
		}

		logger.Error("error to get category by id", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.JSON(http.StatusOK, resp)
}

func (c *categoryHandler) DeleteCategory(ectx echo.Context) error {
	logger := slog.With(
		"handler", "category",
		"method", "DeleteCategory",
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

	categoryID := ectx.Param("categoryId")
	if categoryID == "" {
		logger.Warn("categoryID is empty")
		return ectx.NoContent(http.StatusBadRequest)
	}

	if _, err := uuid.Parse(categoryID); err != nil {
		logger.Warn("invalid categoryID format", "categoryID", categoryID)
		return ectx.NoContent(http.StatusBadRequest)
	}

	userID, ok := c.rdp.GetUserID(ectx.Request().Context())
	if !ok {
		logger.Error("get user id from context")
		DelCookieSession(ectx)
		return ectx.NoContent(http.StatusUnauthorized)
	}

	if err := c.cs.DeleteCategory(ectx.Request().Context(), userID, storeID, categoryID); err != nil {
		if err == models.ErrStoreNotFound {
			logger.Warn("store not found", "error", err)
			return ectx.NoContent(http.StatusNotFound)
		}

		if err == models.ErrStoreNotPertenence {
			logger.Warn("store not pertenence to user", "error", err)
			return ectx.NoContent(http.StatusForbidden)
		}

		if err == models.ErrCategoryNotFound {
			logger.Warn("category not found", "error", err)
			return ectx.NoContent(http.StatusNotFound)
		}

		logger.Error("error to delete category", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.NoContent(http.StatusOK)
}
