package server

import (
	"fmt"

	oapi "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/uakihir0/nostr-rest/server/api"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/errors"
	"github.com/uakihir0/nostr-rest/server/mastodon/api"
	"github.com/uakihir0/nostr-rest/server/mastodon/openapi"
	"github.com/uakihir0/nostr-rest/server/openapi"
)

// Serve
func Serve() {

	e := echo.New()
	e.Use(domain.NewContext)
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.HTTPErrorHandler = errors.ErrorHandler

	dSwagger, err := openapi.GetSwagger()
	if err != nil {
		panic(fmt.Sprintf("Error loading default swagger spec\n: %s", err))
	}
	mSwagger, err := mopenapi.GetSwagger()
	if err != nil {
		panic(fmt.Sprintf("Error loading mastodon swagger spec\n: %s", err))
	}

	// Host validation is not performed
	mSwagger.Servers = nil
	dSwagger.Servers = nil

	// Request validation with OpenAPISpec
	e.Use(oapi.OapiRequestValidatorWithOptions(
		mSwagger, &oapi.Options{
			Skipper: api.Skipper,
			Options: openapi3filter.Options{
				ExcludeResponseBody: true,
				MultiError:          true,
				AuthenticationFunc:  api.AuthHandler,
			}}))
	e.Use(oapi.OapiRequestValidatorWithOptions(
		mSwagger, &oapi.Options{
			Skipper: mapi.Skipper,
			Options: openapi3filter.Options{
				ExcludeResponseBody: true,
				MultiError:          true,
				AuthenticationFunc:  mapi.AuthHandler,
			}}))

	handler := Handler{}
	RegisterHandlers(e, handler)

	e.Logger.Fatal(e.Start(":8080"))
}
