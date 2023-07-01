package api

import (
	"context"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/uakihir0/nostr-rest/server/openapi"
	"strings"
)

type SimpleHandler struct {
}

var _ openapi.ServerInterface = (*SimpleHandler)(nil)

// Skipper
func Skipper(c echo.Context) bool {
	return !IsDefaultPath(c.Path())
}

// IsDefaultPath
func IsDefaultPath(path string) bool {
	return strings.HasPrefix(path, "/api/")
}

// AuthHandler
func AuthHandler(_ context.Context, _ *openapi3filter.AuthenticationInput) error {
	// Authentication is not performed for the default path.
	return nil
}
