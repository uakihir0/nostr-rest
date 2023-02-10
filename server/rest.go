package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uakihir0/nostr-rest/server/api"
	"github.com/uakihir0/nostr-rest/server/errors"
)

// Serve
func Serve() {

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.HTTPErrorHandler = errors.ErrorHandler

	handler := api.Handler{}
	handler.RegisterHandler(e)

	e.Logger.Fatal(e.Start(":8080"))
}
