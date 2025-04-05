package main

import (
	"context"
	"log"

	"github.com/g-villarinho/flash-buy-api/clients"
	"github.com/g-villarinho/flash-buy-api/databases"
	"github.com/g-villarinho/flash-buy-api/handlers"
	"github.com/g-villarinho/flash-buy-api/middlewares"
	"github.com/g-villarinho/flash-buy-api/notifications"
	"github.com/g-villarinho/flash-buy-api/persistence"
	"github.com/g-villarinho/flash-buy-api/pkgs"
	"github.com/g-villarinho/flash-buy-api/repositories"
	"github.com/g-villarinho/flash-buy-api/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initDeps(di *pkgs.Di) {
	// DB
	DB, err := databases.NewPostgresDatabase(context.Background())
	if err != nil {
		log.Fatalf("connect to database: %v", err)
	}

	pkgs.Provide(di, func(di *pkgs.Di) (*gorm.DB, error) {
		return DB, nil
	})

	// Config
	pkgs.Provide(di, pkgs.NewEcdsaKeyPair)
	pkgs.Provide(di, pkgs.NewRequestInfoCtx)

	// Clients
	pkgs.Provide(di, clients.NewSMTPClient)
	pkgs.Provide(di, clients.NewCloudflareClient)

	// Persistence
	pkgs.Provide(di, persistence.NewPostgresRepository)

	// Repositories
	pkgs.Provide(di, repositories.NewOTPRepository)
	pkgs.Provide(di, repositories.NewSessionRepository)
	pkgs.Provide(di, repositories.NewUserRepository)
	pkgs.Provide(di, repositories.NewStoreRepository)
	pkgs.Provide(di, repositories.NewBillboardRepository)
	pkgs.Provide(di, repositories.NewCategoryRepository)
	pkgs.Provide(di, repositories.NewSizeRepository)
	pkgs.Provide(di, repositories.NewColorRepository)
	pkgs.Provide(di, repositories.NewProductRepository)
	pkgs.Provide(di, repositories.NewProductImageRepository)

	// Services
	pkgs.Provide(di, services.NewAuthService)
	pkgs.Provide(di, services.NewOTPService)
	pkgs.Provide(di, services.NewSessionService)
	pkgs.Provide(di, services.NewUserService)
	pkgs.Provide(di, services.NewTokenService)
	pkgs.Provide(di, services.NewRegisterService)
	pkgs.Provide(di, services.NewStoreService)
	pkgs.Provide(di, services.NewBillboardService)
	pkgs.Provide(di, services.NewImageService)
	pkgs.Provide(di, services.NewCategoryService)
	pkgs.Provide(di, services.NewSizeService)
	pkgs.Provide(di, services.NewColorService)
	pkgs.Provide(di, services.NewProductService)
	pkgs.Provide(di, services.NewProductImageService)

	// Handlers
	pkgs.Provide(di, handlers.NewAuthHandler)
	pkgs.Provide(di, handlers.NewRegisterHandler)
	pkgs.Provide(di, handlers.NewConfigHandler)
	pkgs.Provide(di, handlers.NewStoreHandler)
	pkgs.Provide(di, handlers.NewBillboardHandler)
	pkgs.Provide(di, handlers.NewCategoryHandler)
	pkgs.Provide(di, handlers.NewSizeHandler)
	pkgs.Provide(di, handlers.NewColorHandler)
	pkgs.Provide(di, handlers.NewProductHandler)

	//Notifications
	pkgs.Provide(di, notifications.NewEmailNotification)

}

func setupRoutes(e *echo.Echo, di *pkgs.Di) {
	setupEnvironmentRoutes(e, di)
	setupRegisterRouters(e, di)
	setupAuthRoutes(e, di)
	setupStoreRoutes(e, di)
	setupBillboardRoutes(e, di)
	setupCategoryRoutes(e, di)
	setupSizeRoutes(e, di)
	setupColorRoutes(e, di)
	setupProductRoutes(e, di)
}

func setupEnvironmentRoutes(e *echo.Echo, di *pkgs.Di) {
	ch, err := pkgs.Invoke[handlers.EnvironmentHandler](di)
	if err != nil {
		e.Logger.Fatal(err)
	}

	group := e.Group("/v1")
	group.GET("/env", ch.GetEnvs)
}

func setupRegisterRouters(e *echo.Echo, di *pkgs.Di) {
	rh, err := pkgs.Invoke[handlers.RegisterHandler](di)
	if err != nil {
		e.Logger.Fatal(err)
	}

	group := e.Group("/v1")
	group.POST("/register", rh.Register)
}

func setupAuthRoutes(e *echo.Echo, di *pkgs.Di) {
	ah, err := pkgs.Invoke[handlers.AuthHandler](di)
	if err != nil {
		e.Logger.Fatal(err)
	}

	am, err := middlewares.NewAuthMiddleware(di)
	if err != nil {
		e.Logger.Fatal(err)
	}

	group := e.Group("/v1")
	group.POST("/login", ah.Login)
	group.POST("/verify-code", ah.VerifyCode, am.AuthenticateWithoutEmailVerification)
	group.POST("/resend-code", ah.ResendCode, am.AuthenticateWithoutEmailVerification)
	group.GET("/check-code", ah.CheckCode, am.AuthenticateWithoutEmailVerification)
}

func setupStoreRoutes(e *echo.Echo, di *pkgs.Di) {
	sh, err := pkgs.Invoke[handlers.StoreHandler](di)
	if err != nil {
		e.Logger.Fatal(err)
	}

	am, err := middlewares.NewAuthMiddleware(di)
	if err != nil {
		e.Logger.Fatal(err)
	}

	group := e.Group("/v1")
	group.POST("/stores", sh.CreateStore, am.Authenticate)
	group.GET("/me/stores/first", sh.GetUserFirstStore, am.Authenticate)
	group.GET("/stores/:storeId", sh.GetStoreByID, am.Authenticate)
	group.GET("/me/stores", sh.GetStoresByUserID, am.Authenticate)
	group.PUT("/stores/:storeId", sh.UpdateStore, am.Authenticate)
	group.DELETE("/stores/:storeId", sh.DeleteStore, am.Authenticate)
}

func setupBillboardRoutes(e *echo.Echo, di *pkgs.Di) {
	bh, err := pkgs.Invoke[handlers.BillboardHandler](di)
	if err != nil {
		e.Logger.Fatal(err)
	}

	am, err := middlewares.NewAuthMiddleware(di)
	if err != nil {
		e.Logger.Fatal(err)
	}

	group := e.Group("/v1")
	group.POST("/stores/:storeId/billboards", bh.CreateBillboard, am.Authenticate)
	group.GET("/stores/:storeId/billboards", bh.GetBillboards, am.Authenticate)
	group.GET("/stores/:storeId/billboards/:billboardId", bh.GetBillboardByID, am.Authenticate)
	group.DELETE("/stores/:storeId/billboards/:billboardId", bh.DeleteBillboard, am.Authenticate)
}

func setupCategoryRoutes(e *echo.Echo, di *pkgs.Di) {
	ch, err := pkgs.Invoke[handlers.CategoryHandler](di)
	if err != nil {
		e.Logger.Fatal(err)
	}

	am, err := middlewares.NewAuthMiddleware(di)
	if err != nil {
		e.Logger.Fatal(err)
	}

	group := e.Group("/v1")
	group.POST("/stores/:storeId/categories", ch.CreateCategory, am.Authenticate)
	group.GET("/stores/:storeId/categories", ch.GetCategoriesPagedList, am.Authenticate)
	group.GET("/stores/:storeId/categories/:categoryId", ch.GetCategoryByID, am.Authenticate)
	group.DELETE("/stores/:storeId/categories/:categoryId", ch.DeleteCategory, am.Authenticate)
}

func setupSizeRoutes(e *echo.Echo, di *pkgs.Di) {
	sh, err := pkgs.Invoke[handlers.SizeHandler](di)
	if err != nil {
		e.Logger.Fatal(err)
	}

	am, err := middlewares.NewAuthMiddleware(di)
	if err != nil {
		e.Logger.Fatal(err)
	}

	group := e.Group("/v1")
	group.POST("/stores/:storeId/sizes", sh.CreateSize, am.Authenticate)
	group.GET("/stores/:storeId/sizes", sh.GetSizesPagedList, am.Authenticate)
	group.GET("/stores/:storeId/sizes/:sizeId", sh.GetSizeByID, am.Authenticate)
	group.DELETE("/stores/:storeId/sizes/:sizeId", sh.DeleteSize, am.Authenticate)
}

func setupColorRoutes(e *echo.Echo, di *pkgs.Di) {
	ch, err := pkgs.Invoke[handlers.ColorHandler](di)
	if err != nil {
		e.Logger.Fatal(err)
	}

	am, err := middlewares.NewAuthMiddleware(di)
	if err != nil {
		e.Logger.Fatal(err)
	}

	group := e.Group("/v1")
	group.POST("/stores/:storeId/colors", ch.CreateColor, am.Authenticate)
	group.GET("/stores/:storeId/colors", ch.GetColors, am.Authenticate)
	group.GET("/stores/:storeId/colors/:colorId", ch.GetColorByID, am.Authenticate)
	group.DELETE("/stores/:storeId/colors/:colorId", ch.DeleteColor, am.Authenticate)
}

func setupProductRoutes(e *echo.Echo, di *pkgs.Di) {
	ph, err := pkgs.Invoke[handlers.ProductHandler](di)
	if err != nil {
		e.Logger.Fatal(err)
	}

	am, err := middlewares.NewAuthMiddleware(di)
	if err != nil {
		e.Logger.Fatal(err)
	}

	group := e.Group("/v1")
	group.POST("/stores/:storeId/products", ph.CreateProduct, am.Authenticate)
}
