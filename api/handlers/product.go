package handlers

import (
	"log/slog"
	"net/http"

	"github.com/g-villarinho/xp-life-api/models"
	"github.com/g-villarinho/xp-life-api/pkgs"
	"github.com/g-villarinho/xp-life-api/services"
	"github.com/go-playground/form/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProductHandler interface {
	CreateProduct(ectx echo.Context) error
}

type productHandler struct {
	di      *pkgs.Di
	ps      services.ProductService
	rdp     pkgs.RequestDataCtx
	decoder *form.Decoder
}

func NewProductHandler(di *pkgs.Di) (ProductHandler, error) {
	productService, err := pkgs.Invoke[services.ProductService](di)
	if err != nil {
		return nil, err
	}

	ctxData, err := pkgs.Invoke[pkgs.RequestDataCtx](di)
	if err != nil {
		return nil, err
	}

	return &productHandler{
		di:      di,
		ps:      productService,
		rdp:     ctxData,
		decoder: form.NewDecoder(),
	}, nil
}

func (p *productHandler) CreateProduct(ectx echo.Context) error {
	logger := slog.With(
		"handler", "product",
		"method", "CreateProduct",
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

	userID, ok := p.rdp.GetUserID(ectx.Request().Context())
	if !ok {
		logger.Warn("user id not found in context")
		return ectx.NoContent(http.StatusUnauthorized)
	}

	form, err := ectx.MultipartForm()
	if err != nil {
		logger.Error("failed to parse multipart form", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	var payload models.CreateProductPayload
	if err = p.decoder.Decode(&payload, form.Value); err != nil {
		logger.Error("failed to decode form data", "error", err)
		return ectx.NoContent(http.StatusBadRequest)
	}

	images := form.File["images"]
	if len(images) == 0 {
		logger.Warn("no images uploaded")
		return ectx.NoContent(http.StatusBadRequest)
	}

	product, err := payload.ToProduct(storeID)
	if err != nil {
		logger.Error("failed to parse payload to product", "error", err)
		return ectx.NoContent(http.StatusBadRequest)
	}

	if err := p.ps.CreateProduct(ectx.Request().Context(), userID, *product, images); err != nil {
		if err == models.ErrStoreNotFound {
			logger.Warn("store not found", "storeID", storeID)
			return ectx.NoContent(http.StatusNotFound)
		}

		if err == models.ErrStoreNotPertenence {
			logger.Warn("store not pertenence to user", "storeID", storeID)
			return ectx.NoContent(http.StatusForbidden)
		}

		if err == models.ErrCategoryNotFound {
			logger.Warn("category not found", "categoryID", product.CategoryID)
			return ectx.NoContent(http.StatusNotFound)
		}

		if err == models.ErrSizeNotFound {
			logger.Warn("size not found", "sizeID", product.SizeID)
			return ectx.NoContent(http.StatusNotFound)
		}

		if err == models.ErrColorNotFound {
			logger.Warn("color not found", "colorID", product.ColorID)
			return ectx.NoContent(http.StatusNotFound)
		}

		logger.Error("failed to create product", "error", err)
		return ectx.NoContent(http.StatusInternalServerError)
	}

	return ectx.NoContent(http.StatusOK)
}
