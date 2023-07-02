package mapi

import (
	"context"
	"errors"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/nbd-wtf/go-nostr/nip19"
	"github.com/samber/lo"

	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/mastodon/openapi"
)

type MastodonHandler struct {
}

var _ mopenapi.ServerInterface = (*MastodonHandler)(nil)

// Skipper
func Skipper(c echo.Context) bool {
	return !IsMastodonPath(c.Path())
}

// IsMastodonPath
func IsMastodonPath(path string) bool {
	return strings.HasPrefix(path, "/api/")
}

// AuthHandler
func AuthHandler(ctx context.Context, input *openapi3filter.AuthenticationInput) error {

	if input.SecuritySchemeName == "NPubOrNSec" {
		authHeader := input.RequestValidationInput.Request.Header["Authorization"]

		if authHeader == nil {
			return errors.New("authorization not found")
		}
		if !strings.HasPrefix(authHeader[0], "Bearer ") {
			return errors.New("illegal authorization")
		}
		value := authHeader[0][7:]

		// set public key or secret key in custom context
		c := ctx.Value(middleware.EchoContextKey).(*domain.Context)
		if strings.HasPrefix(value, "npub") {
			_, decoded, err := nip19.Decode(value)
			if err != nil {
				return errors.New("illegal public key")
			}
			c.PubKey = lo.ToPtr(domain.UserPubKey(decoded.(string)))
		}
	}
	return nil
}
