package main

import (
	"fmt"

	"github.com/g-villarinho/flash-buy-api/config"
	"github.com/g-villarinho/flash-buy-api/middlewares"
	"github.com/g-villarinho/flash-buy-api/pkgs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	di := pkgs.NewDi()

	err := config.LoadEnv()
	if err != nil {
		e.Logger.Fatal(fmt.Sprintf("load env: %v", err))
	}

	config.NewLogger()

	e.Use(middlewares.CORS())
	e.Use(middleware.Recover())

	initDeps(di)
	setupRoutes(e, di)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Env.API.Port)))
}
